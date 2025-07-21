package data

import (
	"context"
	"fmt"
	"task_manager/controllers/dto"
	"task_manager/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	CreateUser(userPayload dto.RegisterUserDto, userRole string) (models.User, error)
	FindUserByUserName(userName string) (models.User, error)
	PromoteUser(userId string) (models.User, error)
	FindUserById(userId string) (models.User, error)
	GetUserCount() int64
}

type UserSerivce struct {
	collection *mongo.Collection
}

func EncryptPassword(password string) string {
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(pass)
}


func ComparePassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func NewUserService(collection *mongo.Collection) *UserSerivce {
	return &UserSerivce{
		collection: collection,
	}
}

func (us *UserSerivce) PromoteUser(userId string) (models.User, error) {
	// replaceOneOpts := options.FindOneAndReplace()	
	filters := bson.D{{Key: "id", Value: userId}}
	user, err := us.FindUserById(userId)  
	if err != nil {
		return user, err
	}

	user.Role = models.UserRoleAdmin
	var newUser models.User
	if err := us.collection.FindOneAndReplace(context.TODO(), filters, user).Decode(&newUser); err != nil {
		return models.User{}, err
	}

	return newUser, err
}

func (us *UserSerivce) CreateUser(userPayload dto.RegisterUserDto, userRole string) (models.User, error) {
	user := models.User{
		Id: uuid.NewString(),
		UserName: userPayload.UserName,
		Password: EncryptPassword(userPayload.Password),
		Role: userRole,
	}	

	if _, err := us.collection.InsertOne(context.TODO(), user); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (us *UserSerivce) GetUserCount() int64 {
    // nil filter matches all documents
    count, err := us.collection.CountDocuments(context.Background(), bson.D{})
    if err != nil {
        return 0
    }

    return count
}


func (us *UserSerivce) FindUserById(userID string) (models.User, error) {
	findOneOpts := options.FindOne()
	filters := bson.D{{Key: "id", Value: userID}}
	var user models.User
	if err := us.collection.FindOne(context.TODO(), filters, findOneOpts).Decode(&user); err != nil {
		return models.User{}, fmt.Errorf("user not found")
	}

	return user, nil
}

func (us *UserSerivce) FindUserByUserName(UserName string) (models.User, error) {
	findOneOpts := options.FindOne()
	filters := bson.D{{Key: "username", Value: UserName}}
	var user models.User
	if err := us.collection.FindOne(context.TODO(), filters, findOneOpts).Decode(&user); err != nil {
		return models.User{}, fmt.Errorf("user not found")
	}

	return user, nil
}