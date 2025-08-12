// Package controllers
package controllers

import (
	"blogger/modules/user/application/dto"
	"blogger/modules/user/application/usecases"
	"blogger/modules/utils/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usecase *usecases.UserUseCase
}

func NewUserController(usecase *usecases.UserUseCase) *UserController {
	return &UserController{
		usecase: usecase,
	}
}

func (c *UserController) UpdateProfile(ctx *gin.Context) {
	user, err := middleware.GetUserFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	var updateProfileDto dto.UpdateProfileDTO
	if err := ctx.ShouldBind(&updateProfileDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err = c.usecase.UpdateProfile(user.ID, updateProfileDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) Login(ctx *gin.Context) {
	var loginDto dto.LoginUserDTO
	if err := ctx.ShouldBind(&loginDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := c.usecase.Login(loginDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (c *UserController) Register(ctx *gin.Context) {
	var registerDto dto.RegisterUserDTO
	if err := ctx.ShouldBind(&registerDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := c.usecase.CreateUser(registerDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (c *UserController) Promote(ctx *gin.Context) {
	var promoteUserDto dto.PromoteUserDTO
	if err := ctx.ShouldBind(&promoteUserDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	promoter, err := middleware.GetUserFromContext(ctx)
	if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	if err := c.usecase.PromoteUser(promoter.ID, promoteUserDto); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user promoted successfully",
	})
}
