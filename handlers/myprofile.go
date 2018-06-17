package handlers

import (
	"net/http"

	"github.com/hoanhan101/medium/common"

	"go.isomorphicgo.org/go/isokit"
)

// MyProfileHandler handles request for /profile route.
func MyProfileHandler(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := make(map[string]string)
		m["PageTitle"] = "My Profile"
		env.TemplateSet.Render("profile_page", &isokit.RenderParams{Writer: w, Data: m})
	})
}
