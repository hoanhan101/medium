package datastore

import (
	"errors"

	"github.com/hoanhan101/medium/models"
)

// Datastore describes common characteristics of a datastore.
type Datastore interface {
	CreateUser(user *models.User) error
	GetUser(username string) (*models.User, error)
	Close()
}

// Different types of datastore.
const (
	// MYSQL is responsible for MySQL database.
	MYSQL = iota

	// MONGODB is responsible for MongoDB database.
	MONGODB

	// REDIS is responsible for Redis database.
	REDIS
)

// NewDatastore creates a new datastore of a particular type.
func NewDatastore(datastoreType int, dbConnectionString string) (Datastore, error) {
	switch datastoreType {
	case MYSQL:
		return NewMySQLDatastore(dbConnectionString)
	case MONGODB:
		return NewMongoDBDatastore(dbConnectionString)
	case REDIS:
		return NewRedisDatastore(dbConnectionString)
	default:
		return nil, errors.New("The datastore you specified does not exist.")
	}
}
