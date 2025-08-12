// Package entity
package entity

import (
	"time"
)

type Blog struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	User      string    `json:"user"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ViewCount int       `json:"viewCount"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	Comments  []Comment `json:"comments"`
}

// Comment represents a comment on a blog post.
type Comment struct {
	User      string    `json:"user"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

// NewBlog creates and returns a new Blog entity.
func NewBlog(id, title, content, user string, tags []string) Blog {
	now := time.Now()
	return Blog{
		ID:        id,
		Title:     title,
		Content:   content,
		User:      user,
		Tags:      tags,
		CreatedAt: now,
		UpdatedAt: now,
		ViewCount: 0,
		Likes:     0,
		Dislikes:  0,
		Comments:  []Comment{},
	}
}

// Update modifies the blog post’s title and content.
func (b *Blog) Update(title, content string) {
	b.Title = title
	b.Content = content
	b.UpdatedAt = time.Now()
}

// AddView increments the view count.
func (b *Blog) AddView() {
	b.ViewCount++
}

// Like adds a like to the blog post.
func (b *Blog) Like() {
	b.Likes++
}

// Dislike adds a dislike to the blog post.
func (b *Blog) Dislike() {
	b.Dislikes++
}

// AddComment adds a comment to the blog.
func (b *Blog) AddComment(user, message string) {
	comment := Comment{
		User:      user,
		Message:   message,
		Timestamp: time.Now(),
	}
	b.Comments = append(b.Comments, comment)
}
