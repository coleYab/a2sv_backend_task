package utils

import "golang.org/x/crypto/bcrypt"

type PasswordUtils interface {
	HashPassword(password string) string
	ComparePassword(hash, password string) error
}

type PasswordUtil struct{}

// ComparePassword implements PasswordUtils.
func (p PasswordUtil) ComparePassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// HashPassword implements PasswordUtils.
func (p PasswordUtil) HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func NewPasswordUtil() *PasswordUtil {
	return &PasswordUtil{}
}
