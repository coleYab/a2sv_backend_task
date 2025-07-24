// Package controllers : controller module
package controllers

import (
	"fmt"
	"net/http"
	"task_manager/delivery/controllers/dto"
	"task_manager/domain"
	"task_manager/infrastructure"
	"task_manager/usecases"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type JwtData struct {
	UserName string
	ID       string
}

type IUserController interface {
	RegisterUser(ctx *gin.Context)
	Login(ctx *gin.Context)
	PromoteUser(ctx *gin.Context)
}

type UserController struct {
	ps          infrastructure.IPasswordService
	jwtService  infrastructure.IJwtService
	userUseCase usecases.IUserUseCase
}

func NewUserController(uc usecases.IUserUseCase, jwtService infrastructure.IJwtService) *UserController {
	return &UserController{
		ps:          infrastructure.NewPasswordService(),
		jwtService:  jwtService,
		userUseCase: uc,
	}
}

func (uc *UserController) Login(ctx *gin.Context) {
	var loginPayload dto.LoginDto
	if err := ctx.ShouldBind(&loginPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := uc.userUseCase.FindUserByUserName(loginPayload.UserName)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, "invalid credentials")
		return
	}

	if err := uc.ps.ComparePassword(user.Password, loginPayload.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, "invalid credentials")
		return
	}

	// create a jwt token and return the result
	token, err := uc.jwtService.Generate(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to create the jwt token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": token,
		"user":        user,
	})
}

func getUserFromContext(ctx *gin.Context) (domain.User, error) {
	oUser, ok := ctx.Get("user")
	if !ok {
		return domain.User{}, fmt.Errorf("unauthorized")
	}

	user, ok := oUser.(domain.User)
	if !ok {
		return domain.User{}, fmt.Errorf("unauthorized")
	}

	return user, nil
}

func (uc *UserController) PromoteUser(ctx *gin.Context) {
	adminUser, err := getUserFromContext(ctx)
	if err != nil || adminUser.Role != domain.UserRoleAdmin {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID := ctx.Param("id")
	user, err := uc.userUseCase.PromoteUser(userID)
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

func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var registerPayload dto.RegisterUserDto
	if err := ctx.ShouldBind(&registerPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if _, err := uc.userUseCase.FindUserByUserName(registerPayload.UserName); err == nil { // if user exists
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username already taken",
		})
		return
	}

	userRole := domain.UserRoleUser
	if cnt := uc.userUseCase.GetUserCount(); cnt == 0 {
		userRole = domain.UserRoleAdmin
	}

	registerPayload.Password = uc.ps.HashPassword(registerPayload.Password)
	user, err := uc.userUseCase.CreateUser(registerPayload, userRole)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
