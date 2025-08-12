// Package routes: user routes
package routes

import (
	"blogger/modules/user/application/usecases"
	"blogger/modules/user/delivery/http/controllers"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	authMiddleware gin.HandlerFunc
	controller *controllers.UserController
}

func (r *UserRoutes)RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", r.controller.Register)
	router.POST("/auth/login", r.controller.Login)
	router.GET("/auth/logout", r.controller.Logout)
	router.POST("/user/:id/profile", r.authMiddleware, r.controller.UpdateProfile)
	router.PUT("/user/:id/promote", r.authMiddleware, r.controller.Promote)
}


func NewUserRoutes(usecase *usecases.UserUseCase, authMiddelware gin.HandlerFunc) *UserRoutes {
	return &UserRoutes{
		authMiddleware: authMiddelware,
		controller: controllers.NewUserController(usecase),
	}
}
