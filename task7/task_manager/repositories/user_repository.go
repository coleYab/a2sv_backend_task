// Package repositories: repository module
package repositories

import (
	"context"
	"fmt"
	"task_manager/delivery/controllers/dto"
	"task_manager/domain"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IUserRepository interface {
	CreateUser(userPayload dto.RegisterUserDto, userRole string) (domain.User, error)
	FindUserByUserName(userName string) (domain.User, error)
	PromoteUser(userID string) (domain.User, error)
	FindUserByID(userID string) (domain.User, error)
	GetUserCount() int64
}

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: collection,
	}
}

func (us *UserRepository) PromoteUser(userID string) (domain.User, error) {
	// replaceOneOpts := options.FindOneAndReplace()
	filters := bson.D{{Key: "id", Value: userID}}
	user, err := us.FindUserByID(userID)
	if err != nil {
		return user, err
	}

	user.Role = domain.UserRoleAdmin
	var newUser domain.User
	if err := us.collection.FindOneAndReplace(context.TODO(), filters, user).Decode(&newUser); err != nil {
		return domain.User{}, err
	}

	return newUser, err
}

func (us *UserRepository) CreateUser(userPayload dto.RegisterUserDto, userRole string) (domain.User, error) {
	user := domain.User{
		ID:       uuid.NewString(),
		UserName: userPayload.UserName,
		Password: userPayload.Password,
		Role:     userRole,
	}

	if _, err := us.collection.InsertOne(context.TODO(), user); err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (us *UserRepository) GetUserCount() int64 {
	// nil filter matches all documents
	count, err := us.collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		return 0
	}

	return count
}

func (us *UserRepository) FindUserByID(userID string) (domain.User, error) {
	findOneOpts := options.FindOne()
	filters := bson.D{{Key: "id", Value: userID}}
	var user domain.User
	if err := us.collection.FindOne(context.TODO(), filters, findOneOpts).Decode(&user); err != nil {
		return domain.User{}, fmt.Errorf("user not found")
	}

	return user, nil
}

func (us *UserRepository) FindUserByUserName(UserName string) (domain.User, error) {
	findOneOpts := options.FindOne()
	filters := bson.D{{Key: "username", Value: UserName}}
	var user domain.User
	if err := us.collection.FindOne(context.TODO(), filters, findOneOpts).Decode(&user); err != nil {
		return domain.User{}, fmt.Errorf("user not found")
	}

	return user, nil
}
