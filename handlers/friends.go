package handlers

import (
	"net/http"
)

// FriendsHandler TODO
func FriendsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is our friends page."))
}
