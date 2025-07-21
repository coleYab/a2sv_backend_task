package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"task_manager/data"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func JwtAuthMiddleware(us data.IUserService) gin.HandlerFunc {
    secretKey := []byte(os.Getenv("SECRET_KEY"))
	return func (ctx *gin.Context)  {
		authHeader := ctx.GetHeader("Authorization")
        if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return secretKey, nil
        })


        if err != nil || !token.Valid {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok {
            userID := claims["sub"] 
            id, _ := userID.(string)
            user, err := us.FindUserById(id)
             if err != nil {
                 ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
                 return
             }
            ctx.Set("user", user)
        } else {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            return
        }
		
		ctx.Next()
	}
}


