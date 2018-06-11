package handlers

import (
	"net/http"
)

// FeedHandler TODO
func FeedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is our feed."))
}
