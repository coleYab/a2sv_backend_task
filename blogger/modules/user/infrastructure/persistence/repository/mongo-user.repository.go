// Package repository
package repository

import (
	"blogger/modules/user/domain/entity"
	"blogger/modules/user/infrastructure/persistence/mongoentity"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct{
	collection *mongo.Collection
}

var defaultContext = context.TODO()

// CreateUser implements repository.UserRepository.
func (m *MongoUserRepository) CreateUser(user entity.User) error {
	userEntity := mongoentity.FromDomainEntity(user)
	_, err := m.collection.InsertOne(defaultContext, userEntity)
	return err
}

// DeleteUser implements repository.UserRepository.
func (m *MongoUserRepository) DeleteUser(id string) error {
	_, err := m.collection.DeleteOne(defaultContext, bson.M{"_id": id})
	return err
}

// GetUserByEmail implements repository.UserRepository.
func (m *MongoUserRepository) GetUserByEmail(email string) (entity.User, error) {
	var userEntity mongoentity.MongoUser
	err := m.collection.FindOne(defaultContext, bson.M{"email": email}).Decode(&userEntity)
	if err != nil {
		return entity.User{}, err
	}
	return mongoentity.ToDomainEntity(userEntity), nil
}

// GetUserByID implements repository.UserRepository.
func (m *MongoUserRepository) GetUserByID(id string) (entity.User, error) {
	var userEntity mongoentity.MongoUser
	err := m.collection.FindOne(defaultContext, bson.M{"id": id}).Decode(&userEntity)
	if err != nil {
		return entity.User{}, err
	}
	return mongoentity.ToDomainEntity(userEntity), nil
}

// GetUserByUsername implements repository.UserRepository.
func (m *MongoUserRepository) GetUserByUsername(username string) (entity.User, error) {
	var userEntity mongoentity.MongoUser
	err := m.collection.FindOne(defaultContext, bson.M{"username": username}).Decode(&userEntity)
	if err != nil {
		return entity.User{}, err
	}
	return mongoentity.ToDomainEntity(userEntity), nil
}

// ListUsers implements repository.UserRepository.
func (m *MongoUserRepository) ListUsers() ([]entity.User, error) {
	cursor, err := m.collection.Find(defaultContext, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(defaultContext)

	var users []entity.User
	for cursor.Next(defaultContext) {
		var userEntity mongoentity.MongoUser
		if err := cursor.Decode(&userEntity); err != nil {
			return nil, err
		}
		users = append(users, mongoentity.ToDomainEntity(userEntity))
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser implements repository.UserRepository.
func (m *MongoUserRepository) UpdateUser(user entity.User) error {
	userEntity := mongoentity.FromDomainEntity(user)
	_, err := m.collection.UpdateOne(defaultContext, bson.M{"id": user.ID}, bson.M{"$set": userEntity})
	return err
}

func NewUserRepository(collection *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{
		collection: collection,
	}
}
