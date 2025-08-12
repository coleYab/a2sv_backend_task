// Package dto
package dto

type CreateBlogDto struct {
    Title   string   `json:"title" binding:"required"`
    Content string   `json:"content" binding:"required"`
    User    string   `json:"user" binding:"required"`
    Tags    []string `json:"tags"`
}

type CommentDto  struct {
	User string `json:"user" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type GenerateBlogRequest struct {
	Title  string   `json:"title" binding:"required"`
	Schema string   `json:"schema" binding:"required"`
	Tags []string `json:"tags" binding:"required"`
}