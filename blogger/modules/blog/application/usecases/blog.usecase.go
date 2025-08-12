// Package usecases
package usecases

import (
	"blogger/modules/blog/application/dto"
	"blogger/modules/blog/domain/entity"
	"blogger/modules/blog/domain/repository"
	"blogger/modules/blog/infrastructure/ai"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BlogUseCase struct {
	blogRepository  repository.BlogRepository
}

func NewBlogUseCase(
	blogRepository  repository.BlogRepository,
) *BlogUseCase {
	return &BlogUseCase{
		blogRepository:  blogRepository,
	}
}

func (uc *BlogUseCase) GenerateBlog(generateBlogDto dto.GenerateBlogRequest) (entity.Blog, error) {
	prompt := fmt.Sprintf("Write a detailed blog post titled \"%s\".", generateBlogDto.Title)
	if generateBlogDto.Schema != "" {
		prompt += fmt.Sprintf(" Follow this structure: %s.", generateBlogDto.Schema)
	} else {
		prompt += " Include an introduction, main body, and conclusion."
	}	
	blogContent, err := ai.GenerateGeminiContent(prompt)
	if err != nil {
		return entity.Blog{}, err
	}

	blog := entity.NewBlog(
		uuid.NewString(), generateBlogDto.Title, blogContent, "", generateBlogDto.Tags,
	)
	return blog, nil
}

func (uc *BlogUseCase) ImproveBlog(id string) (entity.Blog, error) {
	blog, err := uc.blogRepository.GetBlogByID(id)
	if err != nil {
		return entity.Blog{}, err
	}

	prompt := fmt.Sprintf("Write a detailed blog post titled \"%s\".", blog.Title)
		prompt += fmt.Sprintf(" With this content : %s.", blog.Content)	
		prompt += " Include an introduction, main body, and conclusion."

	blogContent, err := ai.GenerateGeminiContent(prompt)
	if err != nil {
		return entity.Blog{}, err
	}

	blog.Update(blog.Title, blogContent)
	// if err := uc.blogRepository.UpdateBlog(blog); err != nil {
	// 	return entity.Blog{}, err
	// }

	return blog, nil
}

func (uc *BlogUseCase) CommentOnBlog(blogID string, commentDto dto.CommentDto) error {
	blog, err := uc.blogRepository.GetBlogByID(blogID)
	if err != nil {
		return err
	}

	blog.Comments = append(blog.Comments, entity.Comment{
		User:    commentDto.User,
		Message: commentDto.Message,
		Timestamp:  time.Now(),
	})

	return uc.blogRepository.UpdateBlog(blog)
}

func (uc *BlogUseCase) CreateBlog(createBlogDto dto.CreateBlogDto) (entity.Blog, error) {
	blog := entity.NewBlog(
		uuid.NewString(), createBlogDto.Title, createBlogDto.Content, createBlogDto.User, createBlogDto.Tags,
	)
	err := uc.blogRepository.CreateBlog(blog)
	return blog, err
}
	
func (uc *BlogUseCase) UpdateBlog(id string, updateBlogDto dto.UpdateBlogDto) (entity.Blog, error) {
	blog, err := uc.blogRepository.GetBlogByID(id)
	if err != nil {
		return  entity.Blog{}, err
	}

	blog.Update(updateBlogDto.Title, updateBlogDto.Content)

	if err := uc.blogRepository.UpdateBlog(blog); err != nil {
		return  entity.Blog{}, err
	}

	return blog, nil
}

func (uc *BlogUseCase) DeleteBlog(id string) error {
	return uc.blogRepository.DeleteBlog(id)
}


func (uc *BlogUseCase) GetBlogByID(id string) (entity.Blog, error) {
	blog, err := uc.blogRepository.GetBlogByID(id)
	if err != nil {
		return  entity.Blog{}, err
	}

	blog.ViewCount += 1
	if err := uc.blogRepository.UpdateBlog(blog); err != nil {
		return  entity.Blog{}, err
	}

	return blog, nil
}

func (uc *BlogUseCase) LikeBlog(id string) (entity.Blog, error) {
	blog, err := uc.blogRepository.GetBlogByID(id)
	if err != nil {
		return  entity.Blog{}, err
	}

	blog.Like()
	if err := uc.blogRepository.UpdateBlog(blog); err != nil {
		return  entity.Blog{}, err
	}

	return blog, nil
}

func (uc *BlogUseCase) DislikeBlog(id string) (entity.Blog, error) {
	blog, err := uc.blogRepository.GetBlogByID(id)
	if err != nil {
		return entity.Blog{}, err
	}

	blog.Dislike()
	if err := uc.blogRepository.UpdateBlog(blog); err != nil {
		return entity.Blog{}, err
	}

	return blog, nil
}

func (uc *BlogUseCase) GetAllBlogs() ([]entity.Blog, error) {
	return uc.blogRepository.ListBlogs()
}	

func (uc *BlogUseCase) FilterBlogs(
    page int, 
    limit int, 
    sortBy string, 
    order string, 
    tagFilter string, 
    authorFilter string,
) ([]entity.Blog, int, error) {
    blogs, totalCount, err := uc.blogRepository.FilterBlogs(page, limit, sortBy, order, tagFilter, authorFilter)
    return blogs, totalCount, err
}

// ...existing code...