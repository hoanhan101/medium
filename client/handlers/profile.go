package handlers

import (
	"context"
	"log"

	"github.com/hoanhan101/medium/client/common"

	"go.isomorphicgo.org/go/isokit"
)

// TODO.
func ProfileHandler(env *common.Env) isokit.Handler {
	return isokit.HandlerFunc(func(ctx context.Context) {
		log.Println("Client-side Profile Handler")
	})
}
