// Package repository
package repository

import (
	"blogger/modules/user/domain/entity"
	"blogger/modules/user/domain/repository"
	"blogger/modules/user/infrastructure/persistence/mongoentity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTokenRepository struct {
	collection *mongo.Collection
}

// CreateToken implements repository.TokenRepository.
func (m *MongoTokenRepository) CreateToken(token entity.RefreshToken) error {
	tokenEntity := mongoentity.FromDomainToken(token)
	_, err := m.collection.InsertOne(defaultContext, tokenEntity)
	return err
}

// DeleteRefreshToken implements repository.TokenRepository.
func (m *MongoTokenRepository) DeleteRefreshToken(id string) error {
	_, err := m.collection.DeleteOne(defaultContext, bson.M{"_id": id})
	return err
}

// GetRefreshTokenByID implements repository.TokenRepository.
func (m *MongoTokenRepository) GetRefreshTokenByID(id string) (entity.RefreshToken, error) {
	var tokenEntity mongoentity.MongoRefreshToken
	err := m.collection.FindOne(defaultContext, bson.M{"id": id}).Decode(&tokenEntity)
	if err != nil {
		return entity.RefreshToken{}, err
	}
	return mongoentity.ToDomainToken(tokenEntity), nil
}

// ListRefreshTokens implements repository.TokenRepository.
func (m *MongoTokenRepository) ListRefreshTokens() ([]entity.RefreshToken, error) {
	cursor, err := m.collection.Find(defaultContext, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(defaultContext)

	var tokens []entity.RefreshToken
	for cursor.Next(defaultContext) {
		var tokenEntity mongoentity.MongoRefreshToken
		if err := cursor.Decode(&tokenEntity); err != nil {
			return nil, err
		}
		tokens = append(tokens, mongoentity.ToDomainToken(tokenEntity))
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tokens, nil
}

// UpdateRefreshToken implements repository.TokenRepository.
func (m *MongoTokenRepository) UpdateRefreshToken(token entity.RefreshToken) error {
	tokenEntity := mongoentity.FromDomainToken(token)
	_, err := m.collection.UpdateOne(defaultContext, bson.M{"id": token.ID}, bson.M{"$set": tokenEntity})
	return err
}

func NewTokenRepository(collection *mongo.Collection) repository.TokenRepository {
	return &MongoTokenRepository{
		collection: collection,
	}
}
