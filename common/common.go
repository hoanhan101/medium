package common

import (
	"github.com/hoanhan101/medium/common/datastore"

	"go.isomorphicgo.org/go/isokit"
)

// Env holds the datastore connection.
type Env struct {
	DB          datastore.Datastore
	TemplateSet *isokit.TemplateSet
}
