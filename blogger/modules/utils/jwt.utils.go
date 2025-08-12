package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthToken interface {
	Generate(id, username string) (string, error)
	Decode(token string) (string, error)
}

type JwtService struct {
	secretKey string
}

func NewJwtService(secretKey string) *JwtService {
	return &JwtService{
		secretKey: secretKey,
	}
}

func (j *JwtService) Decode(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})

	if err != nil || !token.Valid {
		fmt.Println(err.Error())
		return "", fmt.Errorf("invalid jwt token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userID := claims["sub"]
		id, _ := userID.(string)
		return id, nil
	}

	return "", fmt.Errorf("unable to parse claims from the token")
}

func (j *JwtService) Generate(id, username string) (string, error) {
	// Create token with standard claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  id,
		"name": username,
		"iat":  time.Now().Unix(),                    // issued at
		"exp":  time.Now().Add(time.Hour * 2).Unix(), // expiration
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
