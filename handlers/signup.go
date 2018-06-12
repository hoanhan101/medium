package handlers

import (
	"net/http"

	"github.com/hoanhan101/isocial/validationkit"
)

// SignUpForm is responsible for registering fields.
type SignUpForm struct {
	// FieldNames is a collection of all the fields that we want to prefill
	// in case user makes a mistake.
	FieldNames []string

	// Fields map field names to their corresponding values.
	Fields map[string]string

	// Errors map field names to their corresponding errors.
	Errors map[string]string
}

// SignUpHandler handles http request for signup route.
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	s := SignUpForm{}
	s.FieldNames = []string{
		"username",
		"firstName",
		"lastName",
		"email",
	}
	s.Fields = make(map[string]string)
	s.Errors = make(map[string]string)

	switch r.Method {
	case "GET":
		DisplaySignUpForm(w, r, &s)
	case "POST":
		ValidateSignUpForm(w, r, &s)
	default:
		DisplaySignUpForm(w, r, &s)
	}
}

// DisplaySignUpForm renders template with SignUpForm values.
func DisplaySignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	RenderTemplate(w, "./templates/signup.html", s)
}

// ValidateSignUpForm validates user's input SignUpForm values.
func ValidateSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	// Prefill the values that already entered by user.
	PopulateFormField(r, s)

	// Make sure that every field is non-empty.
	if r.FormValue("username") == "" {
		s.Errors["usernameError"] = "Username is required."
	}

	if r.FormValue("firstName") == "" {
		s.Errors["firstNameError"] = "First name is required."
	}

	if r.FormValue("lastName") == "" {
		s.Errors["lastNameError"] = "Last name is required."
	}

	if r.FormValue("email") == "" {
		s.Errors["emailError"] = "Email is required."
	}

	if r.FormValue("password") == "" {
		s.Errors["passwordError"] = "Password is required."
	}

	if r.FormValue("confirmPassword") == "" {
		s.Errors["confirmPasswordError"] = "Confirm Password is required."
	}

	// Check username and email syntax using validationkit.
	if validationkit.CheckUsernameSyntax(r.FormValue("username")) == false {
		s.Errors["usernameError"] = "The username entered has an improper syntax"
	}

	if validationkit.CheckEmailSyntax(r.FormValue("email")) == false {
		s.Errors["emailError"] = "The email entered has an improper syntax"
	}

	// Check if password and confirm password match
	if r.FormValue("password") != r.FormValue("confirmPassword") {
		s.Errors["confirmPasswordError"] = "Password and confirm password don't match"
	}

	// Only process if there exists no error.
	if len(s.Errors) > 0 {
		DisplaySignUpForm(w, r, s)
	} else {
		ProcessSignUpForm(w, r, s)
	}
}

// PopulateFormField is responsible for prefilling the form.
func PopulateFormField(r *http.Request, s *SignUpForm) {
	for _, fieldName := range s.FieldNames {
		s.Fields[fieldName] = r.FormValue(fieldName)
	}
}

// ProcessSignUpFrom inserts values into database and displays confirmation message.
func ProcessSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	// TODO: Insert to database. For now, only display confirmation message.
	DisplayConfirmation(w, r, s)
}

// DisplayConfirmation notifies a successful registration.
func DisplayConfirmation(w http.ResponseWriter, r *http.Request, s *SignUpForm) {
	RenderUnsafeTemplate(w, "./templates/signupconfirmation.html", s)
}
