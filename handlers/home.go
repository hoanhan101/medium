package handlers

import (
	"net/http"
)

// HomeHandler returns a homepage.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["PageTitle"] = "Home"
	RenderTemplate(w, "./templates/home.html", m)
}
