// Package controllers
package controllers

import (
	"blogger/modules/blog/application/dto"
	"blogger/modules/blog/application/usecases"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	usecase *usecases.BlogUseCase
}

func NewBlogController(usecase *usecases.BlogUseCase) *BlogController {
	return &BlogController{
		usecase: usecase,
	}
}

func (c *BlogController) ImproveBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	blog, err := c.usecase.ImproveBlog(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, blog)
}

func (c *BlogController) GenerateBlog(ctx *gin.Context) {
	var generateBlogSchema dto.GenerateBlogRequest
	if err := ctx.ShouldBind(&generateBlogSchema); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	blog, err := c.usecase.GenerateBlog(generateBlogSchema)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, blog)
}

func (c *BlogController) Comment(ctx *gin.Context) {
	var commentDto dto.CommentDto
	if err := ctx.ShouldBind(&commentDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return  
	}

	blogID := ctx.Param("id")

	if err := c.usecase.CommentOnBlog(blogID, commentDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
}

func (c *BlogController) CreateBlog(ctx *gin.Context) {
	var createBlogDto dto.CreateBlogDto
	if err := ctx.ShouldBind(&createBlogDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	blog, err := c.usecase.CreateBlog(createBlogDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, blog)
}

func (c *BlogController) UpdateBlog(ctx *gin.Context) {
	var updateBlogDto dto.UpdateBlogDto
	if err := ctx.ShouldBind(&updateBlogDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := ctx.Param("id")

	blog, err := c.usecase.UpdateBlog(id, updateBlogDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, blog)
}

func (c *BlogController) DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.usecase.DeleteBlog(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (c *BlogController) GetAllBlogs(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	sortBy := ctx.DefaultQuery("sortBy", "createdAt")  // createdAt, popularity, title, updatedAt
	order := ctx.DefaultQuery("order", "desc")         // asc or desc
	tagFilter := ctx.Query("tag")                      // e.g., "tech"
	authorFilter := ctx.Query("author")                // e.g., "john"

	// Ensure page and limit are valid
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	// 2. Call usecase with filter/sort/pagination
	blogs, totalCount, err := c.usecase.FilterBlogs(page, limit, sortBy, order, tagFilter, authorFilter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. Return paginated response
	ctx.JSON(http.StatusOK, gin.H{
		"blogs":      blogs,
		"page":       page,
		"limit":      limit,
		"total":      totalCount,
		"totalPages": (totalCount + limit - 1) / limit, // round up
	})
}


func (c *BlogController) GetBlogByID(ctx *gin.Context) {
	id := ctx.Param("id")

	blog, err := c.usecase.GetBlogByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}


func (c *BlogController) LikeBlog(ctx *gin.Context) {
	id := ctx.Param("id")

	blog, err := c.usecase.LikeBlog(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}	

func (c *BlogController) UnlikeBlog(ctx *gin.Context) {
	id := ctx.Param("id")

	blog, err := c.usecase.DislikeBlog(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}
