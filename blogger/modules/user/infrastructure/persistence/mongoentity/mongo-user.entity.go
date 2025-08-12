// Package mongoentity
package mongoentity

import (
	"blogger/modules/user/domain/entity"
	"time"
)

type MongoUser struct {
	ID        string    `bson:"id"`
	Username  string    `bson:"username"`
	Email     string    `bson:"email"`
	Role 	  string    `bson:"role"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func FromDomainEntity(user entity.User) MongoUser {
	return MongoUser{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		Role: user.Role,
		Password: user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToDomainEntity(user MongoUser) entity.User {
	return entity.User{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		Role: user.Role,
		Password: user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}