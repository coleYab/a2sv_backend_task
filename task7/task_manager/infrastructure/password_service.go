// Package infrastructure: the infrastructure module
package infrastructure

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type IPasswordService interface {
	HashPassword(string) string
	ComparePassword(string, string) error
}

type PasswordService struct{}

func NewPasswordService() *PasswordService {
	return &PasswordService{}
}

func (j *PasswordService) HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (j *PasswordService) ComparePassword(hash string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		log.Println(err, hash, password)
		return fmt.Errorf("invalid credentials")
	}

	return nil
}
