package handlers

import (
	"context"
	"log"

	"github.com/hoanhan101/medium/client/common"

	"go.isomorphicgo.org/go/isokit"
)

// FriendsHandler handles request for /friends route from client-side.
func FriendsHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		log.Println("Client-side Friends Handler")
		m := make(map[string]string)
		m["PageTitle"] = "Friends"
		env.TemplateSet.Render(
			"friends_content",
			&isokit.RenderParams{
				Data:        m,
				Disposition: isokit.PlacementReplaceInnerContents,
				Element:     env.PrimaryContent,
			})
	})
}
