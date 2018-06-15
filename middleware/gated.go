package middleware

import (
	"log"
	"net/http"

	"github.com/hoanhan101/medium/common/authenticate"
)

// GatedContentHandler is the authentication middleware.
func GatedContentHandler(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shouldRedirectToLogin := false

		secureCookieMap, err := authenticate.ReadSecureCookieValues(w, r)
		if err != nil {
			log.Println(err)
		}

		// Check if the sessionID (sid) has been set.
		if _, ok := secureCookieMap["sid"]; ok == true {
			// Get the user session.
			session, err := authenticate.SessionStore.Get(r, "medium-session")
			log.Printf("medium session values: %+v\n", session.Values)
			if err != nil {
				log.Println(err)
				return
			}

			// Check if the session id stored in the secure cookie matches the
			// id and username on the server-side session.
			if session.Values["sessionID"] == secureCookieMap["sid"] && session.Values["username"] == secureCookieMap["username"] {
				next(w, r)
			} else {
				shouldRedirectToLogin = true
			}
		} else {
			shouldRedirectToLogin = true
		}

		if shouldRedirectToLogin == true {
			http.Redirect(w, r, "/login", 302)
		}
	})
}
