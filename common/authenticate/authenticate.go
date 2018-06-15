package authenticate

import (
	"log"
	"strings"

	"github.com/hoanhan101/medium/common"
	"github.com/hoanhan101/medium/common/utility"
)

// VerifyCredentials verifies username and password.
func VerifyCredentials(e *common.Env, username string, password string) bool {
	// Get the user's record in the datastore.
	u, err := e.DB.GetUser(username)
	if u == nil {
		return false
	}

	if err != nil {
		log.Println(err)
	}

	if strings.ToLower(username) == strings.ToLower(u.Username) && utility.SHA256OfString(password) == u.PasswordHash {
		log.Println("Successful login attempt from user:", u.Username)
		return true
	} else {
		log.Println("Unsuccessful login attempt from user:", u.Username)
		return false
	}
}
