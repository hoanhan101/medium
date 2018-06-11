package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

// ProfileHandler returns a profile for a given username.
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	w.Write([]byte(strings.Join([]string{"This is a profile of", username}, " ")))
}
