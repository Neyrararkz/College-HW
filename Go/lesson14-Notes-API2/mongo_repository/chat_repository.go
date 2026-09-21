package mongo_repository

import (
	"context"
	"time"

	"notes-api/config"
	"notes-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var chatCollection = config.MongoDB.Collection("chats")

func CreateChat(userIDs []int) (*models.Chat, error) {
	chat := models.Chat{
		UserIDs: userIDs,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := chatCollection.InsertOne(ctx, chat)
	if err != nil {
		return nil, err
	}

	chat.ID = res.InsertedID.(primitive.ObjectID)
	return &chat, nil
}

func GetUserChats(userID int) ([]models.Chat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := chatCollection.Find(ctx, bson.M{
		"user_ids": userID,
	})
	if err != nil {
		return nil, err
	}

	var chats []models.Chat
	if err = cursor.All(ctx, &chats); err != nil {
		return nil, err
	}

	return chats, nil
}