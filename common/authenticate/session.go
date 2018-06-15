package authenticate

import (
	"log"
	"net/http"
	// "os"

	"github.com/hoanhan101/medium/models"

	"github.com/gorilla/sessions"
)

// SessionStore persists user's session data in the local filesystem on the
// webserver. If we need anything, it's easy to obtain it from the session
// rather than making a database call.
var SessionStore *sessions.FilesystemStore

func init() {
	// SessionStore = sessions.NewFilesystemStore("/tmp/medium-sessions", []byte(os.Getenv("MEDIUM_HASH_KEY")))

	// Leave the first parameter an empty string to use os.TempDir()
	SessionStore = sessions.NewFilesystemStore("", []byte("CRKVBJs0kfyeQ9Y1"))
}

// CreateUserSession creates a user session.
func CreateUserSession(u *models.User, sessionID string, w http.ResponseWriter, r *http.Request) error {
	// Fetch the session. In the first run, it will be created.
	session, err := SessionStore.Get(r, "medium-session")
	if err != nil {
		log.Println(err)
	}

	// Set the session vales.
	session.Values["sessionID"] = sessionID
	session.Values["username"] = u.Username
	session.Values["firstName"] = u.FirstName
	session.Values["lastName"] = u.LastName
	session.Values["email"] = u.Email

	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// ExpireUserSession is called by the LogoutHandler.
func ExpireUserSession(w http.ResponseWriter, r *http.Request) {
	session, err := SessionStore.Get(r, "medium-session")
	if err != nil {
		log.Println(err)
	}

	session.Options.MaxAge = -1
	session.Save(r, w)
}
