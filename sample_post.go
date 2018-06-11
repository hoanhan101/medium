package main

import (
	"github.com/hoanhan101/isomorphic-web/socialmedia"
	"log"
)

func main() {
	post := socialmedia.NewPost(
		"Hoanh An",
		socialmedia.Moods["happy"],
		"Sample caption",
		"Sanple message body",
		"Sample URL",
		"Sample image URI",
		"Sample thumbnail URI",
		[]string{"Sample keyword 1", "Sample keyword 2"},
	)
	log.Printf("%+v", post)
}
