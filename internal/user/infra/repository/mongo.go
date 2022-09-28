package repository

import (
	"context"
	"github.com/guil95/chat-go/internal/user/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(db *mongo.Database) domain.Repository {
	return mongoRepository{db}
}

func (r mongoRepository) SaveUser(ctx context.Context, user *domain.User) error {
	collection := r.db.Collection("users")

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (r mongoRepository) FindByUserName(ctx context.Context, userName string) (*domain.User, error) {
	collection := r.db.Collection("users")

	var result domain.User

	err := collection.FindOne(ctx, bson.D{{"username", userName}}).Decode(&result)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}

	return &result, nil
}

func (r mongoRepository) FindByUserAndPassword(ctx context.Context, userName, password string) (*domain.User, error) {
	collection := r.db.Collection("users")

	var result domain.User

	err := collection.FindOne(ctx, bson.D{{"username", userName}, {"password", password}}).Decode(&result)
	if err != nil {
		zap.S().Error(err)
		return nil, err
	}

	return &result, nil
}
