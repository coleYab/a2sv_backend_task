package mongoentity

import (
	"blogger/modules/user/domain/entity"
	"time"
)

type MongoRefreshToken struct {
	ID        string `bson:"id" json:"id"`
	UserID    string `bson:"userId" json:"user_id"`
	Token     string             `bson:"token" json:"token"`
	ExpiresAt time.Time          `bson:"expires_at" json:"expires_at"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	Revoked   bool               `bson:"revoked" json:"revoked"`
}

func FromDomainToken(token entity.RefreshToken) MongoRefreshToken {
	return MongoRefreshToken{
		ID: token.ID,
		UserID: token.UserID,
		Token: token.Token,
		ExpiresAt: token.ExpiresAt,
		CreatedAt: token.CreatedAt,
		Revoked: token.Revoked,
	}
}

func ToDomainToken(token MongoRefreshToken) entity.RefreshToken {
	return entity.RefreshToken{
		ID: token.ID,
		UserID: token.UserID,
		Token: token.Token,
		ExpiresAt: token.ExpiresAt,
		CreatedAt: token.CreatedAt,
		Revoked: token.Revoked,
	}
}