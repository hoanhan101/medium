package handlers

import (
	"context"
	"log"

	"github.com/hoanhan101/medium/client/common"

	"go.isomorphicgo.org/go/isokit"
)

// MyProfileHandler handles request for /profile route from client-side.
func MyProfileHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		log.Println("Client-side My Profile Handler")
		m := make(map[string]string)
		m["PageTitle"] = "My Profile"
		env.TemplateSet.Render(
			"profile_content",
			&isokit.RenderParams{
				Data:        m,
				Disposition: isokit.PlacementReplaceInnerContents,
				Element:     env.PrimaryContent,
			})
	})
}
