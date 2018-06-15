package handlers

import (
	"net/http"
)

// MyProfileHandler returns user's profile.
func MyProfileHandler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["PageTitle"] = "My profile"
	RenderTemplate(w, "./templates/myprofile.html", m)
}
