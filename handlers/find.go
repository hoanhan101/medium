package handlers

import (
	"net/http"
)

// FindHandler returns a find page.
func FindHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is our find page."))
}
