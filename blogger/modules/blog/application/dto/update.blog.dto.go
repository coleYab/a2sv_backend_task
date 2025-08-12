// Package dto
package dto

type UpdateBlogDto struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}