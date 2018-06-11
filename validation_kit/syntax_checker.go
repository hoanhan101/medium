package validation_kit

import (
	"log"
	"regexp"
)

const (
	// UsernameRegexp uses regular expression for user's naming.
	// Remember to use backtick instead of double quotes to use literal values of a
	// string. ^ start the pattern and $ end the pattern. Only accepts 0 or 1 @
	// character in the beginning of the username. (\) matches any case-insensitive
	// alpha-numeric character, also includes underscore character. {1,15} accepts
	// at least 1 minimum character and maximum 15 characters.
	UsernameRegexp string = `^@?(\w){1,15}$`

	// EmailRegedxp uses regular experssion for email naming.
	EmailRegexp string = `(?i)^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,3})+$`
)

// CheckUsernameSyntax checks username if matches with our regular expression
// pattern. It then returns a boolean.
func CheckUsernameSyntax(username string) bool {
	result := false
	r, err := regexp.Compile(UsernameRegexp)
	if err != nil {
		log.Fatal(err)
	}
	result = r.MatchString(username)
	return result
}

// CheckEmailSyntax checks email if matches with our regular expression
// pattern. It then returns a boolean.
func CheckEmailSyntax(email string) bool {
	result := false
	r, err := regexp.Compile(EmailRegexp)
	if err != nil {
		log.Fatal(err)
	}
	result = r.MatchString(email)
	return result
}
