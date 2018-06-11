package handlers

import (
	"net/http"
)

// LogoutHandler TODO
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is our logout page."))
}
