package main

import (
	"github.com/hoanhan101/isomorphic-web/social_media"
	"log"
)

func main() {
	post := social_media.NewPost(
		"Hoanh An",
		social_media.Moods["happy"],
		"Sample caption",
		"Sanple message body",
		"Sample URL",
		"Sample image URI",
		"Sample thumbnail URI",
		[]string{"Sample keyword 1", "Sample keyword 2"},
	)
	log.Printf("%+v", post)
}
