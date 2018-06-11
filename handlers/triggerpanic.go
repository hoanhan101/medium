package handlers

import (
	"net/http"
)

// TriggerPanicHandler simulates a panic.
func TriggerPanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("Triggering a panic!")
}
