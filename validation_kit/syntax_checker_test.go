package validation_kit

import (
	"testing"
)

// These constant gives us checkboxes for visualization.
const (
	succeed = "\u2713"
	failed  = "\u2717"
)

// Test basic checking
func TestCheckUsernameBasic(t *testing.T) {
	tests := [][]string{
		[]string{"Minimum characters", ""},
		[]string{"Maximum characters", "hoanhanhoanhanhoanhhoanhanhoanhan"},
		[]string{"Special symbols", "g-@p!h#r"},
		[]string{"Underscore symbol", "hoanh_an"},
	}

	expectedResults := []bool{
		false,
		false,
		false,
		true,
	}

	for i := range tests {
		result := CheckUsernameSyntax(tests[i][1])
		if result == expectedResults[i] {
			t.Logf("\t%s\tCase: %v\tTarget: %v\t Result: %v", succeed, tests[i][0], tests[i][1], result)
		} else {
			t.Errorf("\t%s\tCase: %v\tTarget: %v\t Result: %v", failed, tests[i][0], tests[i][1], result)
		}
	}
}

// Test random valid usernames.
func TestCheckUsernameRandom(t *testing.T) {
	for i := 0; i < 10008; i++ {
		username := GenerateRandomUsername()
		result := CheckUsernameSyntax(username)
		if result == true {
			t.Logf("\t%s\tCase: %v\tTarget: %v\t Result: %v", succeed, "Auto generated valid username", username, result)
		} else {
			t.Logf("\t%s\tCase: %v\tTarget: %v\t Result: %v", failed, "Auto generated valid username", username, result)
		}
	}
}

// TODO: Test email validation
