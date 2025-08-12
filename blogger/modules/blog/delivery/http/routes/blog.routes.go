// Package routes: blog routes
package routes

import (
	"blogger/modules/blog/application/usecases"
	"blogger/modules/blog/delivery/http/controllers"

	"github.com/gin-gonic/gin"
)

type BlogRoutes struct {
	authMiddleware gin.HandlerFunc
	controller     *controllers.BlogController
}

func (r *BlogRoutes) RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/blogs", r.controller.CreateBlog)
	router.POST("/blogs/:id/comment", r.controller.Comment)
	router.PUT("/blogs/:id", r.controller.UpdateBlog)
	router.DELETE("/blogs/:id", r.controller.DeleteBlog)
	router.GET("/blogs", r.controller.GetAllBlogs)
	router.GET("/blogs/:id/like", r.controller.LikeBlog)
	router.GET("/blogs/:id/dislike", r.controller.UnlikeBlog)
	router.GET("/blogs/:id", r.controller.GetBlogByID)
	router.POST("/blogs/generate/", r.controller.GenerateBlog)
	router.POST("/blogs/:id/improve", r.controller.ImproveBlog)
}

func NewBlogRoutes(usecase *usecases.BlogUseCase, authMiddleware gin.HandlerFunc) *BlogRoutes {
	return &BlogRoutes{
		authMiddleware: authMiddleware,
		controller:     controllers.NewBlogController(usecase),
	}
}
