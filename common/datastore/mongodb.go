package datastore

import (
	"github.com/hoanhan101/medium/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDBDatastore embeds mgo.Session pointer.
type MongoDBDatastore struct {
	*mgo.Session
}

// NewMongoDBDatastore creates a new MongoDB datastore.
func NewMongoDBDatastore(url string) (*MongoDBDatastore, error) {
	// Create the connection.
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	return &MongoDBDatastore{Session: session}, nil
}

// CreateUser inserts an user into database.
func (m *MongoDBDatastore) CreateUser(user *models.User) error {
	// Return the MongoDB session by issuing a Copy.
	session := m.Copy()
	defer session.Close()

	// Get the User collection in the mediumdb database.
	userCollection := session.DB("mediumdb").C("User")

	// Insert the new user record to the session.
	err := userCollection.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

// GetUser returns an user object for a given username.
func (m *MongoDBDatastore) GetUser(username string) (*models.User, error) {
	// Get the session/
	session := m.Copy()
	defer session.Close()

	// Get the User collection.
	userCollection := session.DB("mediumdb").C("User")

	// Create a new user instance/
	u := models.User{}

	// Find the matching username and return 1 record.
	err := userCollection.Find(bson.M{"username": username}).One(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// Close connection.
func (m *MongoDBDatastore) Close() {
	m.Close()
}
