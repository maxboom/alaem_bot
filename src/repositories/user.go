package repositories

import (
	"database/sql"
	"entity"
	"fmt"
)

// UserRepositoryI Interface
type UserRepositoryI interface {
	GetUser(db *sql.DB, username string) entity.DBUserT
	CreateUser(db *sql.DB, username string)
}

// UserRepositoryT Type
type UserRepositoryT struct {
}

// GetUser Method
func (s UserRepositoryT) GetUser(db *sql.DB, username string) entity.DBUserT {
	var user entity.DBUserT

	err := db.QueryRow("SELECT id, username, is_authorized FROM users where username = ?", username).Scan(&user.ID, &user.Username, &user.IsAuthorized)

	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	defer db.Close()

	return user
}

// CreateUser Method
func (s UserRepositoryT) CreateUser(db *sql.DB, username string) {

	insert, err := db.Query("INSERT INTO users (username, is_authorized) VALUES ('" + username + "', true )")

	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	defer insert.Close()
}
