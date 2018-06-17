package main

import (
	"log"
	"strings"

	"github.com/hoanhan101/medium/client/common"
	"github.com/hoanhan101/medium/client/handlers"

	"go.isomorphicgo.org/go/isokit"
	"honnef.co/go/js/dom"
)

var D = dom.GetWindow().Document().(dom.HTMLDocument)

func run() {
	log.Println("GopherJS is running on the client-side.")

	// isokit.FetchTemplateBundle make XHR calls to the server to fetch the
	// template bundle. Run as a Goroutine to prevent the function from
	// blocking. After receiving the templateSet, attach it to the env
	// variable so it is accessible throughout the client-side.
	templateSetChannel := make(chan *isokit.TemplateSet)
	go isokit.FetchTemplateBundle(templateSetChannel)
	ts := <-templateSetChannel

	// env is conceptually the same as the one we defined in the back-end.
	env := common.Env{}
	env.TemplateSet = ts
	env.Window = dom.GetWindow()
	env.Document = dom.GetWindow().Document()
	env.PrimaryContent = env.Document.GetElementByID("primaryContent")

	// isokit.NewRouter is pretty much like the gorilla's mux router, except it
	// is used on the client-side.
	r := isokit.NewRouter()
	r.Handle("/feed", handlers.FeedHandler(&env))
	r.Handle("/friends", handlers.FriendsHandler(&env))
	r.Handle("/profile", handlers.MyProfileHandler(&env))
	r.Listen()

	initialEventHandlers(&env)
}

// initialEventHandlers initializes various event handlers using the common env
// variable. Since there is only one page in the single-page web app, we need
// a way to initialize envent handlers for each page. TODO.
func initialEventHandlers(env *common.Env) {
	log.Println("Location:", env.Window.Location().Href)

	// Get the current page name.
	l := strings.Split(env.Window.Location().Href, "/")
	pageName := l[len(l)-1]

	// TODO.
	switch pageName {
	}
}

func main() {
	switch readyState := D.ReadyState(); readyState {
	case "loading":
		D.AddEventListener("DOMContentLoaded", false, func(dom.Event) {
			go run()
		})
	case "interactive", "complete":
		run()
	default:
		log.Println("Encountered different value for dom.ReadyState:", readyState)
	}
}
