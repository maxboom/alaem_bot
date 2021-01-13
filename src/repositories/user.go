package repositories

import (
	"entity"

	"gorm.io/gorm"
)

// UserRepositoryI Interface
type UserRepositoryI interface {
	GetUser(db *gorm.DB, username string) entity.DBUserT
	UpdateUser(db *gorm.DB, user entity.DBUserT) entity.DBUserT
}

// UserRepositoryT Type
type UserRepositoryT struct {
}

// GetUser Method
func (UserRepositoryT) GetUser(db *gorm.DB, username string) entity.DBUserT {
	var user entity.DBUserT

	db.FirstOrCreate(&user, entity.DBUserT{
		Username:     username,
		IsAuthorized: false,
	})

	return user
}

// UpdateUser Method
func (UserRepositoryT) UpdateUser(db *gorm.DB, user entity.DBUserT) entity.DBUserT {
	db.Save(user)

	return UserRepositoryT{}.GetUser(db, user.Username)
}
