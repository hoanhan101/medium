package handlers

import (
	"net/http"

	"github.com/hoanhan101/medium/models/socialmedia"
)

// PostForm represents a social media post form.
type PostForm struct {
	// FieldNames is a collection of all the fields that we want to prefill
	// in case user makes a mistake.
	FieldNames []string

	// Fields map field names to their corresponding values.
	Fields map[string]string

	// Errors map field names to their corresponding errors.
	Errors map[string]string
}

// PostPreviewHandler handles http request for postpreview route.
func PostPreviewHandler(w http.ResponseWriter, r *http.Request) {
	p := PostForm{}
	p.FieldNames = []string{
		"caption",
		"messageBody",
	}
	p.Fields = make(map[string]string)
	p.Errors = make(map[string]string)

	switch r.Method {
	case "GET":
		DisplayPostForm(w, r, &p)
	case "POST":
		ValidatePostForm(w, r, &p)
	default:
		DisplayPostForm(w, r, &p)
	}
}

// DisplayPostForm renders template with PostForm values.
func DisplayPostForm(w http.ResponseWriter, r *http.Request, p *PostForm) {
	RenderTemplate(w, "./templates/postform.html", p)
}

// ValidatePostForm validates user's input PostForm values.
func ValidatePostForm(w http.ResponseWriter, r *http.Request, p *PostForm) {
	// Prefill the values that already entered by user.
	PopulatePostFormField(r, p)

	// Make sure that every field is non-empty.
	if r.FormValue("caption") == "" {
		p.Errors["captionError"] = "Caption is required."
	}

	if r.FormValue("messageBody") == "" {
		p.Errors["messageBodyError"] = "Message body is required."
	}

	// Only process if there exists no error. Otherwise, display the PostForm
	// again with entered values.
	if len(p.Errors) > 0 {
		DisplayPostForm(w, r, p)
	} else {
		DisplayPostPreview(w, r, p)
	}
}

// PopulatePostFormField is responsible for prefilling the form.
func PopulatePostFormField(r *http.Request, p *PostForm) {
	for _, fieldName := range p.FieldNames {
		p.Fields[fieldName] = r.FormValue(fieldName)
	}
}

// DisplayPostPreview displays a preview of the post.
func DisplayPostPreview(w http.ResponseWriter, r *http.Request, p *PostForm) {
	// Sample preview.
	preview := socialmedia.NewPost(
		"hoanhan",
		p.Fields["caption"],
		p.Fields["messageBody"],
		"Sample URL",
		"Sample Image URI",
		"Sample Thumnail URI",
		[]string{"Key 1", "Key 2", "Key 3"},
	)

	RenderTemplate(w, "./templates/postpreview.html", preview)
}
