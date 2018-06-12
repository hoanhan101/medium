package endpoints

import (
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Post Endpoint"))
}
