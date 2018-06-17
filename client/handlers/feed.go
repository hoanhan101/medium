package handlers

import (
	"context"
	"log"

	"github.com/hoanhan101/medium/client/common"

	"go.isomorphicgo.org/go/isokit"
)

// FeedHandler handles request for /feed route from client-side.
func FeedHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		log.Println("Client-side Feed Handler")
		m := make(map[string]string)
		m["PageTitle"] = "Feed"
		env.TemplateSet.Render(
			"feed_content",
			&isokit.RenderParams{
				Data:        m,
				Disposition: isokit.PlacementReplaceInnerContents,
				Element:     env.PrimaryContent,
			})
	})
}
