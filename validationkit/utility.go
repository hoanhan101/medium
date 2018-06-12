package validationkit

import (
	"math/rand"
	"time"
)

// GenerateRandomUsername is a helper function used for testing.
func GenerateRandomUsername() string {
	rand.Seed(time.Now().UnixNano())

	usernameLength := rand.Intn(15) + 1

	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_")
	b := make([]rune, usernameLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	randomUsername := string(b)

	zeroOrOne := rand.Intn(2)
	if zeroOrOne == 1 {
		randomUsername = "@" + randomUsername
	}
	return randomUsername
}
