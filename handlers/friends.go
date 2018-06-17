package handlers

import (
	"net/http"

	"github.com/hoanhan101/medium/common"

	"go.isomorphicgo.org/go/isokit"
)

// FriendsHandler handles request for /friends route.
func FriendsHandler(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := make(map[string]string)
		m["PageTitle"] = "Friends"
		env.TemplateSet.Render("friends_page", &isokit.RenderParams{Writer: w, Data: m})
	})
}
