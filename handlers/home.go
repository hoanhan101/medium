package handlers

import (
	"net/http"
)

// HomeHandler returns a homepage.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is our homepage."))
}
