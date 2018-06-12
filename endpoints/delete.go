package endpoints

import (
	"net/http"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Post Endpoint"))
}
