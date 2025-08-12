// Package middleware
package middleware

import (
	"blogger/modules/user/domain/entity"
	"blogger/modules/user/domain/repository"
	"blogger/modules/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(ctx *gin.Context) (entity.User, error) {
	oUser, ok := ctx.Get("user")
	if !ok {
		return entity.User{}, fmt.Errorf("unauthorized")
	}

	user, ok := oUser.(entity.User)
	if !ok {
		return entity.User{}, fmt.Errorf("unauthorized")
	}

	return user, nil
}


func VerifyRole(ctx *gin.Context, role string) error {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return err
	}

	if user.Role == role {
		return nil
	}

	return fmt.Errorf("forbidden resource")
}

func JwtAuthMiddleware(us repository.UserRepository, authService utils.AuthToken) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := authService.Decode(tokenString)
		if err == nil {
			user, err := us.GetUserByID(userID)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
				return
			}
			ctx.Set("user", user)
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Next()
	}
}