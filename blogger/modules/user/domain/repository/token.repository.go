// Package repository
package repository

import (
	"blogger/modules/user/domain/entity"
)

type TokenRepository interface {
	CreateToken(token entity.RefreshToken) error
	GetRefreshTokenByID(id string) (entity.RefreshToken, error)
	UpdateRefreshToken(token entity.RefreshToken) error
	DeleteRefreshToken(id string) error
	ListRefreshTokens() ([]entity.RefreshToken, error)
}
