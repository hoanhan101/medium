package handlers

import (
	"net/http"
)

// FooHandler gets the foo context value.
func FooHandler(w http.ResponseWriter, r *http.Request) {
	foo := r.Context().Value("foo").(string)
	w.Write([]byte(foo))
}
