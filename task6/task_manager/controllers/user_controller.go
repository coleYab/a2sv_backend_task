package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"task_manager/controllers/dto"
	"task_manager/data"
	"task_manager/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

type JwtData struct {
	UserName string
	Id string
}

type IUserController interface {
	RegisterUser(ctx *gin.Context)
	Login(ctx *gin.Context)
	PromoteUser(ctx *gin.Context)
}

type UserController struct {
	userService data.IUserService
}

func NewUserController(service data.IUserService) *UserController {
	return &UserController{
		userService : service,
	}
}

func generateAccessToken(user models.User) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")

    // Create token with standard claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": user.Id,              // subject (e.g., user ID)
        "name": user.UserName,
        "iat": time.Now().Unix(),         // issued at
        "exp": time.Now().Add(time.Hour * 2).Unix(), // expiration
    })

    // Sign the token with the secret key
    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
		return "", err
    }

	return tokenString, nil
}

func (uc *UserController)Login(ctx *gin.Context) {
	var loginPayload dto.LoginDto
	if err := ctx.ShouldBind(&loginPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}	

	user, err := uc.userService.FindUserByUserName(loginPayload.UserName)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "invalid credentials")
		return
	}

	if err := data.ComparePassword(loginPayload.Password, user.Password); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusUnauthorized, "invalid credentials")
		return
	}

	// create a jwt token and return the result
	token, err := generateAccessToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to create the jwt token",
		})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": token,
		"user": user,
	})
}

func getUserFromContext(ctx *gin.Context) (models.User, error) {
	oUser, ok := ctx.Get("user")	
	if !ok {
		return models.User{}, fmt.Errorf("unauthorized")
	}

	user, ok := oUser.(models.User)
	if !ok {
		return models.User{}, fmt.Errorf("unauthorized")
	}

	return user, nil
}

func (uc *UserController) PromoteUser(ctx *gin.Context) {
	adminUser, err := getUserFromContext(ctx)	
	if err != nil || adminUser.Role != models.UserRoleAdmin {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId := ctx.Param("id")
	user, err := uc.userService.PromoteUser(userId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (uc *UserController)RegisterUser(ctx *gin.Context) {
	var registerPayload dto.RegisterUserDto
	if err := ctx.ShouldBind(&registerPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}


	if _, err := uc.userService.FindUserByUserName(registerPayload.UserName); err == nil { // if user exists
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username already taken",
		})
		return 
	}

	userRole := models.UserRoleUser
	if cnt := uc.userService.GetUserCount(); cnt == 0 {
		userRole = models.UserRoleAdmin
	}

	user, err := uc.userService.CreateUser(registerPayload, userRole)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return 
	}

	ctx.JSON(http.StatusCreated, user)
}
