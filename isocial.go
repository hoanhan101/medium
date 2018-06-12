package main

import (
	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/hoanhan101/isocial/endpoints"
	"github.com/hoanhan101/isocial/handlers"
	"github.com/hoanhan101/isocial/middleware"
)

const (
	// Web server port.
	PORT = ":8080"
)

func main() {
	r := mux.NewRouter()

	// Core routes.
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("GET,POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	r.HandleFunc("/feed", handlers.FeedHandler).Methods("GET")
	r.HandleFunc("/friends", handlers.FriendsHandler).Methods("GET")
	r.HandleFunc("/find", handlers.FindHandler).Methods("GET,POST")
	r.HandleFunc("/profile", handlers.MyProfileHandler).Methods("GET")
	r.HandleFunc("/profile/{username}", handlers.ProfileHandler).Methods("GET")

	// Temporary routes simulate different scenarios that are handled by
	// middleware functions:
	// - panic simulates panic recovery
	// - foo simulates persistent context value
	r.HandleFunc("/panic", handlers.TriggerPanicHandler).Methods("GET")
	r.HandleFunc("/foo", handlers.FooHandler).Methods("GET")

	// CRUD APIs for social media posts.
	r.HandleFunc("/api/{username}", endpoints.FetchPosts).Methods("GET")
	r.HandleFunc("/api/{postid}", endpoints.CreatePost).Methods("POST")
	r.HandleFunc("/api/{postid}", endpoints.UpdatePost).Methods("PUT")
	r.HandleFunc("/api/{postid}", endpoints.DeletePost).Methods("DELETE")

	// ghandlers.LoggingHandler(os.Stdout, r) is the default gorilla's logging
	// handler. middleware.RecoverPanicHandler() chains the ghandlers to catch
	// any panic causes. Finally, middleware.ContextHandler persists the
	// context value, which is foo in this situation.
	http.Handle("/", middleware.ContextHandler(
		middleware.RecoverPanicHandler(ghandlers.LoggingHandler(os.Stdout, r))))

	// Pass the context value through the request.
	http.ListenAndServe(PORT, nil)
}
