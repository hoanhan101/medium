package handlers

import (
	"net/http"
)

// MyProfileHandler returns our user's profile.
func MyProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my profile"))
}
