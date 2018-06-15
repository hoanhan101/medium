package handlers

import (
	"log"
	"net/http"

	"github.com/hoanhan101/medium/common"
	"github.com/hoanhan101/medium/common/authenticate"
	"github.com/hoanhan101/medium/common/utility"
	"github.com/hoanhan101/medium/validationkit"
)

// LoginForm is responsible for registering Login information.
type LoginForm struct {
	PageTitle string

	// FieldNames is a collection of all the fields that we want to prefill
	// in case user makes a mistake.
	FieldNames []string

	// Fields map field names to their corresponding values.
	Fields map[string]string

	// Errors map field names to their corresponding errors.
	Errors map[string]string
}

// LoginHandler handles http request for login route.
func LoginHandler(e *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := LoginForm{}
		l.PageTitle = "Log In"
		l.FieldNames = []string{
			"username",
			"password",
		}
		l.Fields = make(map[string]string)
		l.Errors = make(map[string]string)

		switch r.Method {
		case "GET":
			DisplayLoginForm(w, r, &l)
		case "POST":
			ValidateLoginForm(w, r, &l, e)
		default:
			DisplayLoginForm(w, r, &l)
		}
	})
}

// DisplayLoginForm renders template with LoginForm values.
func DisplayLoginForm(w http.ResponseWriter, r *http.Request, l *LoginForm) {
	RenderTemplate(w, "./templates/login.html", l)
}

// ValidateLoginForm validates user's input LoginForm values.
func ValidateLoginForm(w http.ResponseWriter, r *http.Request, l *LoginForm, e *common.Env) {
	// Prefill the values that already entered by user.
	PopulateLoginFormField(r, l)

	// Make sure that every field is non-empty.
	if r.FormValue("username") == "" {
		l.Errors["usernameError"] = "Username is required."
	}

	if r.FormValue("password") == "" {
		l.Errors["passwordError"] = "Password is required."
	}

	// Check username syntax using validationkit.
	if validationkit.CheckUsernameSyntax(r.FormValue("username")) == false {
		l.Errors["usernameError"] = "The username entered has an improper syntax"
	}

	// Only process if there exists no error. Otherwise, display the LoginForm
	// again with entered values.
	if len(l.Errors) > 0 {
		DisplayLoginForm(w, r, l)
	} else {
		ProcessLoginForm(w, r, l, e)
	}
}

// PopulateLoginFormField is responsible for prefilling the form.
func PopulateLoginFormField(r *http.Request, l *LoginForm) {
	for _, fieldName := range l.FieldNames {
		l.Fields[fieldName] = r.FormValue(fieldName)
	}
}

// ProcessLoginForm validates user's identify and create a session.
func ProcessLoginForm(w http.ResponseWriter, r *http.Request, l *LoginForm, e *common.Env) {
	authResult := authenticate.VerifyCredentials(e, r.FormValue("username"), r.FormValue("password"))
	log.Println("Authenticate Result:", authResult)

	// If user successfully login, create a cookie session and redirect to feed route.
	if authResult == true {
		// Generate a session ID.
		sessionID := utility.GenerateUUID()

		// Fetch the user record from the datastore.
		u, err := e.DB.GetUser(r.FormValue("username"))
		if err != nil {
			log.Println("Error encountered while fetching user record:", err)
			http.Redirect(w, r, "/login", 302)
			return
		}

		// Create a secure cookie.
		err = authenticate.CreateSecureCookie(u, sessionID, w, r)
		if err != nil {
			log.Println("Error encountered while creating secure cookie:", err)
			http.Redirect(w, r, "/login", 302)
			return
		}

		// Create a user session.
		err = authenticate.CreateUserSession(u, sessionID, w, r)
		if err != nil {
			log.Println("Error encountered while creating user session:", err)
			http.Redirect(w, r, "/login", 302)
			return
		}

		http.Redirect(w, r, "/feed", 302)
	} else {
		l.Errors["usernameError"] = "Invalid login."
		DisplayLoginForm(w, r, l)
	}
}
