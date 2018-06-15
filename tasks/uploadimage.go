package tasks

import (
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// ImageResizeTask implements the Task interface.
type ImageResizeTask struct {
	BaseImageName      string
	ImageFileExtension string
}

// NewResizeImageTasks creates a new image resizing task.
func NewResizeImageTask(baseImageName string, imageFileExtension string) *ImageResizeTask {
	return &ImageResizeTask{
		BaseImageName:      baseImageName,
		ImageFileExtension: imageFileExtension,
	}
}

// Perform performs the actual image resizing task.
func (t *ImageResizeTask) Perform() {
	// Create an image thumbnail path.
	thumbnailPath := t.BaseImageName + "_thumb.png"
	originalImage, err := os.Open(t.BaseImageName + t.ImageFileExtension)
	defer originalImage.Close()

	if err != nil {
		log.Println("Error", err)
		return
	}

	// Decode the original image.
	img, err := png.Decode(originalImage)
	if err != nil {
		log.Println("Encountered error while decoding the image:", err)
		return
	}

	// Resize to make a thumbnail.
	thumbnail := resize.Resize(270, 0, img, resize.Lanczos3)
	thumbnailFile, err := os.Create(thumbnailPath)
	defer thumbnailFile.Close()

	if err != nil {
		log.Println("Encountered error while resizing the image:", err)
		return
	}

	// Write back the thumbnail image to its file.
	png.Encode(thumbnailFile, thumbnail)
}
