package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hoanhan101/medium/common/asyncq"
	"github.com/hoanhan101/medium/common/utility"
	"github.com/hoanhan101/medium/tasks"
)

// UploadImageForm is responsible for uploading image.
type UploadImageForm struct {
	PageTitle string

	// FieldNames is a collection of all the fields that we want to prefill
	// in case user makes a mistake.
	FieldNames []string

	// Fields map field names to their corresponding values.
	Fields map[string]string

	// Errors map field names to their corresponding errors.
	Errors map[string]string
}

// UploadImageHandler handles http request for upload/image route.
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	u := UploadImageForm{}
	u.PageTitle = "Upload Image"
	u.Fields = make(map[string]string)
	u.Errors = make(map[string]string)

	switch r.Method {
	case "GET":
		DisplayUploadImageForm(w, r, &u)
	case "POST":
		ValidateUploadImageForm(w, r, &u)
	default:
		DisplayUploadImageForm(w, r, &u)
	}
}

// DisplayUploadImageForm renders template with UploadImageForm values.
func DisplayUploadImageForm(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {
	RenderTemplate(w, "./templates/uploadimageform.html", u)
}

// ValidateUploadImageForm validates user's input UploadImageForm values.
func ValidateUploadImageForm(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {
	ProcessUploadImage(w, r, u)
}

// ProcessUploadImage stores images and creates thumbnail.
func ProcessUploadImage(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {
	// Asynchronous flag.
	shouldProcessThumbnailAsynchronously := true

	// Get the image file that upload by request.
	file, fileheader, err := r.FormFile("imageFile")
	defer file.Close()

	if err != nil {
		log.Println("Cannot read read upload file: ", err)
		return
	}

	// It's a good practice not to trust the filename that user sent.
	// Generate a uuid for it instead.
	randomFileName := utility.GenerateUUID()

	if fileheader != nil {
		// Get the filename extension.
		extension := filepath.Ext(fileheader.Filename)

		// The input parameter is the total maximum bytes that are stored in
		// memory for the file. In this case, It's about 32 MB.
		r.ParseMultipartForm(32 << 20)

		imagePathWithoutExtension := "./static/uploads/images/" + randomFileName
		f, err := os.OpenFile(imagePathWithoutExtension+extension, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()

		if err != nil {
			log.Println("Cannot open file:", err)
			return
		}

		// Copy the file content to our local file.
		io.Copy(f, file)

		// Initialize a new image resizing task.
		thumbnailResizeTask := tasks.NewResizeImageTask(imagePathWithoutExtension, extension)

		// Send the task over TaskQueue channel. Otherwise, process in a
		// synchronous manner.
		if shouldProcessThumbnailAsynchronously == true {
			asyncq.TaskQueue <- thumbnailResizeTask
		} else {
			thumbnailResizeTask.Perform()
		}

		// This is a data object that we will pass to the template
		m := make(map[string]string)
		m["PageTitle"] = "Upload Image Preview"
		m["thumbnailPath"] = strings.TrimPrefix(imagePathWithoutExtension, ".") + "_thumb.png"
		m["imagePath"] = strings.TrimPrefix(imagePathWithoutExtension, ".") + ".png"

		// Wait 1 second for the TaskQueue then render template.
		time.Sleep(1 * time.Second)
		RenderTemplate(w, "./templates/uploadimagepreview.html", m)
	} else {
		w.Write([]byte("Failed to process uploaded file."))
	}
}
