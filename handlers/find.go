package handlers

import (
	"net/http"
)

// FindHandler TODO
func FindHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is our find page."))
}
