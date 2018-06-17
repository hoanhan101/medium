package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hoanhan101/medium/common"
	"github.com/hoanhan101/medium/common/asyncq"
	"github.com/hoanhan101/medium/common/datastore"
	"github.com/hoanhan101/medium/endpoints"
	"github.com/hoanhan101/medium/handlers"
	"github.com/hoanhan101/medium/middleware"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"go.isomorphicgo.org/go/isokit"
)

const (
	// TLS port.
	PORT = ":8443"
)

var WebAppRoot = os.Getenv("MEDIUM_APP_ROOT")

func main() {
	// Start a task dispatcher with 9 workers.
	asyncq.StartTaskDispatcher(9)

	// Create a new datastore instance.
	db, err := datastore.NewDatastore(datastore.MYSQL, "medium:medium@/mediumdb")
	// db, err := datastore.NewDatastore(datastore.MONGODB, "localhost:27017")
	// db, err := datastore.NewDatastore(datastore.REDIS, "localhost:6379")
	defer db.Close()

	if err != nil {
		log.Print(err)
	}

	// Use env variable to pass the datastore connection as the dependency
	// injection.
	env := common.Env{}

	// Create a new template set. It allows to persist all the templates in
	// memory and render a particular one given its name. Here, we bundle all
	// templates and send to client-side in one payload. After the intial page
	// load, all subsequent template rendering can be perform on the client side.
	isokit.TemplateFilesPath = WebAppRoot + "/templates"
	isokit.TemplateFileExtension = ".html"
	ts := isokit.NewTemplateSet()
	ts.GatherTemplates()
	env.TemplateSet = ts
	env.DB = db

	// Create a new mux router.
	r := mux.NewRouter()

	// Global routes.
	r.HandleFunc("/", handlers.HomeHandler)
	r.Handle("/signup", handlers.SignUpHandler(&env)).Methods("GET", "POST")

	r.Handle("/login", handlers.LoginHandler(&env)).Methods("GET", "POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET", "POST")

	// Gated routes that are only available for authenticated users.
	r.Handle("/feed", middleware.GatedContentHandler(handlers.FeedHandler(&env))).Methods("GET")
	r.Handle("/friends", middleware.GatedContentHandler(handlers.FriendsHandler(&env))).Methods("GET")
	r.Handle("/profile", middleware.GatedContentHandler(handlers.MyProfileHandler(&env))).Methods("GET")

	r.Handle("/find", middleware.GatedContentHandler(handlers.FindHandler)).Methods("GET", "POST")
	r.Handle("/profile/{username}", middleware.GatedContentHandler(handlers.ProfileHandler)).Methods("GET")
	r.Handle("/upload/image", middleware.GatedContentHandler(handlers.UploadImageHandler)).Methods("GET", "POST")
	r.Handle("/postpreview", middleware.GatedContentHandler(handlers.PostPreviewHandler)).Methods("GET", "POST")

	// Temporary routes simulate different scenarios that are handled by
	// middleware functions:
	// - panic simulates panic recovery
	// - foo simulates persistent context value
	r.HandleFunc("/panic", handlers.TriggerPanicHandler).Methods("GET")
	r.HandleFunc("/foo", handlers.FooHandler).Methods("GET")

	// Client-Side routing.
	// client.js and client.js.map are both produced by GopherJS.
	r.Handle("/js/client.js", isokit.GopherjsScriptHandler(WebAppRoot))
	r.Handle("/js/client.js.map", isokit.GopherjsScriptMapHandler(WebAppRoot))

	// This route will be used to fetch the template from the server-side to
	// the client-side.
	r.Handle("/template-bundle", handlers.TemplateBundleHandler(&env))

	// CRUD APIs for social media posts.
	r.HandleFunc("/api/{username}", endpoints.FetchPosts).Methods("GET")
	r.HandleFunc("/api/{postid}", endpoints.CreatePost).Methods("POST")
	r.HandleFunc("/api/{postid}", endpoints.UpdatePost).Methods("PUT")
	r.HandleFunc("/api/{postid}", endpoints.DeletePost).Methods("DELETE")

	// Fix path to static folder.
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(WebAppRoot+"/static"))))

	// ghandlers.LoggingHandler(os.Stdout, r) is the default gorilla's logging
	// handler. middleware.RecoverPanicHandler() chains the ghandlers to catch
	// any panic causes. Finally, middleware.ContextHandler persists the
	// context value, which is foo in this situation. Instead of
	// http.Handle("/", middleware.ContextHandler(middleware.RecoverPanicHandler(ghandlers.LoggingHandler(os.Stdout, r))))
	// can use alice package to chain middle function in a pretty way.
	loggedRouter := ghandlers.LoggingHandler(os.Stdout, r)
	stdChain := alice.New(middleware.RecoverPanicHandler)
	http.Handle("/", stdChain.Then(loggedRouter))

	err = http.ListenAndServeTLS(PORT, WebAppRoot+"/certs/mediumcert.pem", WebAppRoot+"/certs/mediumkey.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
