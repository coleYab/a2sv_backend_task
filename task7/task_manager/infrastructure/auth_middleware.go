// Package infrastructure: infrastructure module
package infrastructure

import (
	"net/http"
	"strings"
	"task_manager/repositories"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(us repositories.IUserRepository, jwtService IJwtService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		userID, err := jwtService.GetUserIDFromToken(tokenString)
		if err == nil {
			user, err := us.FindUserByID(userID)
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
