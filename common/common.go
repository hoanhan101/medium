package common

import (
	"github.com/hoanhan101/medium/common/datastore"
)

// Env holds the datastore connection.
type Env struct {
	DB datastore.Datastore
}
