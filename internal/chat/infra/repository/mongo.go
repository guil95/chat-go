package repository

import (
	"context"

	"github.com/guil95/chat-go/internal/chat"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	db *mongo.Database
}

func NewMongoRepository(db *mongo.Database) chat.ChatRepository {
	return mongoRepository{db}
}

func (r mongoRepository) SaveChat(ctx context.Context, chat chat.Chat) error {
	collection := r.db.Collection("chats")

	_, err := collection.InsertOne(ctx, chat)
	if err != nil {
		return err
	}

	return nil
}
