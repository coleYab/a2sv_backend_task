package infrastructure

import (
	"fmt"
	"os"
	"task_manager/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IJwtService interface {
	Generate(user domain.User) (string, error)
	GetUserIDFromToken(tokenString string) (string, error)
}

type JwtService struct {
	secretKey string
}

func NewJwtService(secretKey string) *JwtService {
	return &JwtService{
		secretKey: secretKey,
	}
}

func (j *JwtService) GetUserIDFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secretKey, nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid jwt token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userID := claims["sub"]
		id, _ := userID.(string)
		return id, nil
	}

	return "", fmt.Errorf("unable to parse claims from the token")
}

func (j *JwtService) Generate(user domain.User) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")

	// Create token with standard claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID, // subject (e.g., user ID)
		"name": user.UserName,
		"iat":  time.Now().Unix(),                    // issued at
		"exp":  time.Now().Add(time.Hour * 2).Unix(), // expiration
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
