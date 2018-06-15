package handlers

import (
	"net/http"
)

// FeedHandler returns a feed page.
func FeedHandler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["PageTitle"] = "Feed"
	RenderTemplate(w, "./templates/feed.html", m)
}
