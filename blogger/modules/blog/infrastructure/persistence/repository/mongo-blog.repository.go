// Package repository
package repository

import (
	"blogger/modules/blog/domain/entity"
	"blogger/modules/blog/infrastructure/persistence/mongoentity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoBlogRepository struct{
	collection *mongo.Collection
}

var defaultContext = context.TODO()

// CreateBlog implements repository.BlogRepository.
func (m *MongoBlogRepository) CreateBlog(blog entity.Blog) error {
	blogEntity := mongoentity.FromDomainEntity(blog)
	_, err := m.collection.InsertOne(defaultContext, blogEntity)
	return err
}

// DeleteBlog implements repository.BlogRepository.
func (m *MongoBlogRepository) DeleteBlog(id string) error {
	_, err := m.collection.DeleteOne(defaultContext, bson.M{"_id": id})
	return err
}

// GetBlogByID implements repository.BlogRepository.
func (m *MongoBlogRepository) GetBlogByID(id string) (entity.Blog, error) {
	var blogEntity mongoentity.MongoBlog
	err := m.collection.FindOne(defaultContext, bson.M{"id": id}).Decode(&blogEntity)
	if err != nil {
		return entity.Blog{}, err
	}
	return mongoentity.ToDomainEntity(blogEntity), nil
}

// ListBlogs implements repository.BlogRepository.
func (m *MongoBlogRepository) ListBlogs() ([]entity.Blog, error) {
	cursor, err := m.collection.Find(defaultContext, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(defaultContext)

	var blogs []entity.Blog
	for cursor.Next(defaultContext) {
		var blogEntity mongoentity.MongoBlog
		if err := cursor.Decode(&blogEntity); err != nil {
			return nil, err
		}
		blogs = append(blogs, mongoentity.ToDomainEntity(blogEntity))
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return blogs, nil
}

func (m *MongoBlogRepository) FilterBlogs(
	page int,
	limit int,
	sortBy string,
	order string,
	tagFilter string,
	authorFilter string,
) ([]entity.Blog, int, error) {
	var filters = make(primitive.M)
	if tagFilter != "" {
		filters["tags"] = bson.M{"$in": []string{tagFilter}}
	}
	if authorFilter != "" {
		filters["author"] = authorFilter
	}

	// 1. Count total matching documents
	totalCount, err := m.collection.CountDocuments(defaultContext, filters)
	if err != nil {
		return nil, 0, err
	}

	// 2. Apply sorting
	sort := bson.M{}
	switch sortBy {
	case "popularity":
		sort["likes"] = -1
	case "title":
		sort["title"] = 1
	case "updatedAt":
		sort["updatedAt"] = -1
	default:
		sort["createdAt"] = -1
	}

	// 3. Paginate results
	skip := int64((page - 1) * limit)
	lim := int64(limit)
	cursor, err := m.collection.Find(defaultContext, filters, &options.FindOptions{
		Sort:  sort,
		Skip:  &skip,
		Limit: &lim,
	})
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(defaultContext)

	var blogs []entity.Blog
	for cursor.Next(defaultContext) {
		var blogEntity mongoentity.MongoBlog
		if err := cursor.Decode(&blogEntity); err != nil {
			return nil, 0, err
		}
		blogs = append(blogs, mongoentity.ToDomainEntity(blogEntity))
	}

	if err := cursor.Err(); err != nil {
		return nil, 0, err
	}

	return blogs, int(totalCount), nil
}

// UpdateBlog implements repository.BlogRepository.
func (m *MongoBlogRepository) UpdateBlog(blog entity.Blog) error {
	blogEntity := mongoentity.FromDomainEntity(blog)
	_, err := m.collection.UpdateOne(defaultContext, bson.M{"id": blog.ID}, bson.M{"$set": blogEntity})
	return err
}

func NewBlogRepository(collection *mongo.Collection) *MongoBlogRepository {
	return &MongoBlogRepository{
		collection: collection,
	}
}
