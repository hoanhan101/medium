package handlers

import (
	"net/http"
)

// FriendsHandler returns friends page.
func FriendsHandler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["PageTitle"] = "Friends"
	RenderTemplate(w, "./templates/friends.html", m)
}
