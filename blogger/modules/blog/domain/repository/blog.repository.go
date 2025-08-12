// Package repository
package repository

import (
	"blogger/modules/blog/domain/entity"
)

type BlogRepository interface {
	CreateBlog(blog entity.Blog) error
	GetBlogByID(id string) (entity.Blog, error)
	UpdateBlog(blog entity.Blog) error
	DeleteBlog(id string) error
	ListBlogs() ([]entity.Blog, error)
	FilterBlogs(
		page int,
		limit int,
		sortBy string,
		order string,
		tagFilter string,
		authorFilter string,
	) ([]entity.Blog, int, error)
}
