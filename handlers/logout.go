package handlers

import (
	"net/http"

	"github.com/hoanhan101/medium/common/authenticate"
)

// LogoutHandler handles logout request.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)
}
