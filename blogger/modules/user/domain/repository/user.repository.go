// Package repository
package repository

import (
	"blogger/modules/user/domain/entity"
)

type UserRepository interface {
	CreateUser(user entity.User) error
	GetUserByID(id string) (entity.User, error)
	GetUserByEmail( email string) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id string) error
	ListUsers() ([]entity.User, error)
}
