package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hoanhan101/medium/common"
	"github.com/hoanhan101/medium/common/datastore"
	"github.com/hoanhan101/medium/endpoints"
	"github.com/hoanhan101/medium/handlers"
	"github.com/hoanhan101/medium/middleware"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	// Web server port.
	PORT = ":8080"
)

func main() {
	// Create a new datastore instance.
	db, err := datastore.NewDatastore(datastore.MYSQL, "medium:medium@/mediumdb")
	// db, err := datastore.NewDatastore(datastore.MONGODB, "localhost:27017")
	// db, err := datastore.NewDatastore(datastore.REDIS, "localhost:6379")
	defer db.Close()

	if err != nil {
		log.Print(err)
	}

	// Use env variable to pass the datastore connection as the dependency
	// injection to the SignUpHandler.
	env := common.Env{DB: db}

	// Create a new mux router.
	r := mux.NewRouter()

	// Core routes.
	r.Handle("/signup", handlers.SignUpHandler(&env)).Methods("GET", "POST")

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("GET", "POST")
	r.HandleFunc("/postpreview", handlers.PostPreviewHandler).Methods("GET", "POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	r.HandleFunc("/feed", handlers.FeedHandler).Methods("GET")
	r.HandleFunc("/friends", handlers.FriendsHandler).Methods("GET")
	r.HandleFunc("/find", handlers.FindHandler).Methods("GET", "POST")
	r.HandleFunc("/profile", handlers.MyProfileHandler).Methods("GET")
	r.HandleFunc("/profile/{username}", handlers.ProfileHandler).Methods("GET")
	r.HandleFunc("/upload/image", handlers.UploadImageHandler).Methods("GET", "POST")

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
	http.Handle("/", middleware.ContextHandler(middleware.RecoverPanicHandler(ghandlers.LoggingHandler(os.Stdout, r))))

	// Fix path to static folder.
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Pass the context value through the request.
	http.ListenAndServe(PORT, nil)
}
