package entity

import (
	"time"
)

type RefreshToken struct {
	ID        string 			`json:"id"`
	UserID    string 			`json:"userId"`
	Token     string             `json:"token"`
	ExpiresAt time.Time          `json:"expires_at"`
	CreatedAt time.Time          `json:"created_at"`
	Revoked   bool               `json:"revoked"`
}

func NewToken(id, userID, token string, expiresAt, createdAt time.Time, revoked bool) RefreshToken {
	return RefreshToken{
		ID: id,
		UserID: userID,
		Token: token,
		ExpiresAt: expiresAt,
		CreatedAt: createdAt,
		Revoked: revoked,
	}
}

func (t *RefreshToken)Revoke() {
	t.Revoked = true
} 