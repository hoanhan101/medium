// Package socialmedia implements common functionality needed for social media
// web applications.
package socialmedia

import (
	"time"
)

// AuditableContent keeps the common logs for each post.
type AuditableContent struct {
	TimeCreated  time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
	CreatedBy    string    `json:"createdBy"`
	ModifiedBy   string    `json:"modifiedBy"`
}

// Post represents a social media post.
type Post struct {
	AuditableContent
	Caption      string   `json:"caption"`
	MessageBody  string   `json:"messageBody"`
	URL          string   `json:"url"`
	ImageURI     string   `json:"imageURI"`
	ThumbnailURI string   `json:"thumbnailURI"`
	Keywords     []string `json:"keywords"`
	Likers       []string `json:"likers"`
}

// NewPost is responsible for creating a new social media post.
func NewPost(
	username string,
	caption string,
	messageBody string,
	url string,
	imageURI string,
	thumbnailURI string,
	keywords []string,
) *Post {
	auditableContent := AuditableContent{
		CreatedBy:   username,
		TimeCreated: time.Now(),
	}
	return &Post{
		Caption:          caption,
		MessageBody:      messageBody,
		URL:              url,
		ImageURI:         imageURI,
		ThumbnailURI:     thumbnailURI,
		Keywords:         keywords,
		AuditableContent: auditableContent,
	}
}
