package mongo_repository

import (
	"context"
	"time"

	"notes-api/config"
	"notes-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var messageCollection = config.MongoDB.Collection("messages")

func CreateMessage(msg models.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := messageCollection.InsertOne(ctx, msg)
	return err
}

func GetMessages(chatID primitive.ObjectID) ([]models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := messageCollection.Find(ctx, bson.M{
		"chat_id": chatID,
	})
	if err != nil {
		return nil, err
	}

	var messages []models.Message
	if err = cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}