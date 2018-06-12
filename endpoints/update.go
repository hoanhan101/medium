package endpoints

import (
	"net/http"
)

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Post Endpoint"))
}
