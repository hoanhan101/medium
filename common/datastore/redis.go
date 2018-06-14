package datastore

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/hoanhan101/medium/models"

	"github.com/mediocregopher/radix.v2/pool"
)

// RedisDatastore embeds pool pointer.
type RedisDatastore struct {
	*pool.Pool
}

// NewRedisDatastore creates a new MongoDB datastore.
func NewRedisDatastore(address string) (*RedisDatastore, error) {
	// Connect via tcp with a size of 10.
	connectionPool, err := pool.New("tcp", address, 10)
	if err != nil {
		return nil, err
	}

	return &RedisDatastore{Pool: connectionPool}, nil
}

// CreateUser inserts an user into database.
func (r *RedisDatastore) CreateUser(user *models.User) error {
	// Store user object as an JSON encoded string.
	userJSON, err := json.Marshal(*user)
	if err != nil {
		return err
	}

	// Issue a SET command.
	if r.Cmd("SET", "user:"+user.Username, string(userJSON)).Err != nil {
		return errors.New("Failed to execute Redis SET command")
	}

	return nil
}

// GetUser returns an user object for a given username.
func (r *RedisDatastore) GetUser(username string) (*models.User, error) {
	// Check if the user exists.
	exists, err := r.Cmd("EXISTS", "user:"+username).Int()
	if err != nil {
		return nil, err
	} else if exists == 0 {
		return nil, nil
	}

	var u models.User

	// Get the string values representation of data.
	userJSON, err := r.Cmd("GET", "user:"+username).Str()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Unmarshal the JSON represetation into variable u.
	if err := json.Unmarshal([]byte(userJSON), &u); err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}

// Close connection.
func (r *RedisDatastore) Close() {
	r.Empty()
}
