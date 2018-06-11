// Package social_media implements common functionality needed for social media
// web applications.
package social_media

import (
	"time"
)

//go:generate stringer -type=MoodState
type MoodState int

// Moods holds the various mood states with keys to serve as
// aliases to their respective mood state.
var Moods map[string]MoodState

// Possible mood states using an iota enumerator.
const (
	MoodStateNeutral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
	MoodStateHopeful
	MoodStateThrilled
	MoodStateBored
	MoodStateShy
	MoodStateComical
	MoodStateOnCloudNine
)

// AuditableContent keeps the common logs for each post.
type AuditableContent struct {
	TimeCreated  time.Time
	TimeModified time.Time
	CreatedBy    string
	ModifiedBy   string
}

// Post represents a social media post.
type Post struct {
	AuditableContent // Embedded type
	Caption          string
	MessageBody      string
	URL              string
	ImageURI         string
	ThumbnailURI     string
	Keywords         []string
	Likers           []string
	AuthorMood       MoodState
}

// init is responsible for initializing the mood state.
func init() {
	Moods = map[string]MoodState{
		"neutral":   MoodStateNeutral,
		"happy":     MoodStateHappy,
		"sad":       MoodStateSad,
		"angry":     MoodStateAngry,
		"hopeful":   MoodStateHopeful,
		"thrilled":  MoodStateThrilled,
		"bored":     MoodStateBored,
		"shy":       MoodStateShy,
		"comical":   MoodStateComical,
		"cloudnine": MoodStateOnCloudNine,
	}
}

// NewPost is responsible for creating a new social media post.
func NewPost(
	username string,
	mood MoodState,
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
		AuthorMood:       mood,
		Keywords:         keywords,
		AuditableContent: auditableContent,
	}
}
