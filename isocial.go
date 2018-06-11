package main

import (
	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hoanhan101/isocial/handlers"
	"github.com/hoanhan101/isocial/middleware"
)

const (
	// Web server port.
	PORT = ":8080"
)

func main() {
	r := mux.NewRouter()

	// Core routes
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("GET,POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	r.HandleFunc("/feed", handlers.FeedHandler).Methods("GET")
	r.HandleFunc("/friends", handlers.FriendsHandler).Methods("GET")
	r.HandleFunc("/find", handlers.FindHandler).Methods("GET,POST")
	r.HandleFunc("/profile", handlers.MyProfileHandler).Methods("GET")
	r.HandleFunc("/profile/{username}", handlers.ProfileHandler).Methods("GET")

	// Temporary routes simulate different scenarios that are handled by our
	// middleware functions.
	// - panic route simulates panic recovery
	// - foo route simulates context value passing
	r.HandleFunc("/panic", handlers.TriggerPanicHandler).Methods("GET")
	r.HandleFunc("/foo", handlers.FooHandler).Methods("GET")

	// ghandlers.LoggingHandler(os.Stdout, r) is the default gorilla's logging
	// handler. middleware.RecoverPanicHandler() chains the ghandlers to catch
	// any panic causes. Finally, middleware.ContextHandler persists the
	// context value, which is foo in this situation.
	http.Handle("/", middleware.ContextHandler(
		middleware.RecoverPanicHandler(ghandlers.LoggingHandler(os.Stdout, r))))

	// Pass the context value through the request.
	http.ListenAndServe(PORT, nil)
}
