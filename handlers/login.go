package handlers

import (
	"net/http"
)

// LoginHandler TODO
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is our login page."))
}
