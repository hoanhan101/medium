package authenticate

import (
	"log"
	"net/http"
	// "os"
	"time"

	"github.com/hoanhan101/medium/models"

	"github.com/gorilla/securecookie"
)

var (
	// hashKey is used to authenticated cookie value using hash-based message
	// authentication code (hmac)
	hashKey []byte

	// blockKey is used to encrypt the cookie value.
	blockKey []byte

	// s holds the hashKey and blockKey.
	s *securecookie.SecureCookie
)

func init() {
	// hashKey = []byte(os.Getenv("MEDIUM_HASH_KEY"))
	// blockKey = []byte(os.Getenv("MEDIUM_BLOCK_KEY"))
	hashKey = []byte("CRKVBJs0kfyeQ9Y1")
	blockKey = []byte("9LtmRLzVH27CwxrO")

	s = securecookie.New(hashKey, blockKey)
}

// CreateSecureCookie creates a secure cookie function.
func CreateSecureCookie(u *models.User, sessionID string, w http.ResponseWriter, r *http.Request) error {
	// Our cookie will only store 2 values: username and sessionID.
	value := map[string]string{
		"username": u.Username,
		"sid":      sessionID,
	}

	// If the cookie value is encoded properly, create a new cookie name
	// session with that value.
	if encoded, err := s.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:     "session",
			Value:    encoded,
			Path:     "/",
			Secure:   true,
			HttpOnly: true,
		}

		// SetCookie sens a set cookie header to the ResponseWriter's header.
		http.SetCookie(w, cookie)
	} else {
		log.Println(err)
		return err
	}

	return nil
}

// ReadSecureCookieValues returns the cookie values.
func ReadSecureCookieValues(w http.ResponseWriter, r *http.Request) (map[string]string, error) {
	if cookie, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("session", cookie.Value, &value); err == nil {
			return value, nil
		} else {
			return nil, err
		}
	} else {
		return nil, nil
	}
}

// ExpireSecureCookie is used to implement Logout functionality.
func ExpireSecureCookie(w http.ResponseWriter, r *http.Request) {
	// Set MaxAge to -1 to make that cookie expired.
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	// Set some headers here to prevent the web browser from caching previous
	// logout responses.
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")

	http.SetCookie(w, cookie)

	// Redirect user to the login page.
	http.Redirect(w, r, "/login", 301)
}
