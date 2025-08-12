// Package mongoentity
package mongoentity

import (
	"blogger/modules/blog/domain/entity"
	"time"
)

type MongoBlog struct {
    ID        string           `bson:"id"`
    Title     string           `bson:"title"`
    Content   string           `bson:"content"`
    User      string           `bson:"user"`
    Tags      []string         `bson:"tags"`
    CreatedAt time.Time        `bson:"createdAt"`
    UpdatedAt time.Time        `bson:"updatedAt"`
    ViewCount int              `bson:"viewCount"`
    Likes     int              `bson:"likes"`
    Dislikes  int              `bson:"dislikes"`
    Comments  []MongoComment   `bson:"comments"`
}

type MongoComment struct {
    User      string    `bson:"user"`
    Message   string    `bson:"message"`
    Timestamp time.Time `bson:"timestamp"`
}

func FromDomainEntity(blog entity.Blog) MongoBlog {
    comments := make([]MongoComment, len(blog.Comments))
    for i, c := range blog.Comments {
        comments[i] = MongoComment{
            User:      c.User,
            Message:   c.Message,
            Timestamp: c.Timestamp,
        }
    }
    return MongoBlog{
        ID:        blog.ID,
        Title:     blog.Title,
        Content:   blog.Content,
        User:      blog.User,
        Tags:      blog.Tags,
        CreatedAt: blog.CreatedAt,
        UpdatedAt: blog.UpdatedAt,
        ViewCount: blog.ViewCount,
        Likes:     blog.Likes,
        Dislikes:  blog.Dislikes,
        Comments:  comments,
    }
}

func ToDomainEntity(blog MongoBlog) entity.Blog {
    comments := make([]entity.Comment, len(blog.Comments))
    for i, c := range blog.Comments {
        comments[i] = entity.Comment{
            User:      c.User,
            Message:   c.Message,
            Timestamp: c.Timestamp,
        }
    }
    return entity.Blog{
        ID:        blog.ID,
        Title:     blog.Title,
        Content:   blog.Content,
        User:      blog.User,
        Tags:      blog.Tags,
        CreatedAt: blog.CreatedAt,
        UpdatedAt: blog.UpdatedAt,
        ViewCount: blog.ViewCount,
        Likes:     blog.Likes,
        Dislikes:  blog.Dislikes,
        Comments:  comments,
    }
}