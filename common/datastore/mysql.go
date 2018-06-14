package datastore

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hoanhan101/medium/models"
)

// MySQLDatastore embeds sql.DB pointer.
type MySQLDatastore struct {
	*sql.DB
}

// NewMySQLDatastore creates a new MySQL datastore.
func NewMySQLDatastore(dataSourceName string) (*MySQLDatastore, error) {
	connection, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &MySQLDatastore{DB: connection}, nil
}

// CreateUser inserts an user into database.
func (m *MySQLDatastore) CreateUser(user *models.User) error {
	// Begin the transaction and defer to rollback if anything goes wrong.
	tx, err := m.Begin()
	defer tx.Rollback()

	if err != nil {
		log.Println(err)
	}

	// Create a prepare statement for the query that helps us prevent sql injection.
	stmt, err := tx.Prepare("INSERT INTO user(uuid, username, first_name, last_name, email, password_hash) VALUES (?, ?, ?, ?, ?, ?)")
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return err
	}

	// Execute the statement.
	_, err = stmt.Exec(user.UUID, user.Username, user.FirstName, user.LastName, user.Email, user.PasswordHash)
	if err != nil {
		log.Println(err)
		return err
	}

	// Commit the transaction.
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// GetUser returns an user object for a given username.
func (m *MySQLDatastore) GetUser(username string) (*models.User, error) {
	// Prepare a sql statement and defer to close.
	stmt, err := m.Prepare("SELECT uuid, username, first_name, last_name, email, password_hash, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM user WHERE username = ?")
	defer stmt.Close()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Get the first row returned from the query.
	row := stmt.QueryRow(username)

	// Create a user object.
	u := models.User{}

	// Copy the column from the matched row into the corresponding values.
	err = row.Scan(&u.UUID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.PasswordHash, &u.TimestampCreated, &u.TimestampModified)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, err
}

// Close connection.
func (m *MySQLDatastore) Close() {
	m.Close()
}
