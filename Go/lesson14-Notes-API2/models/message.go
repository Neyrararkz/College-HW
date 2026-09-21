package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ChatID  primitive.ObjectID `bson:"chat_id" json:"chat_id"`
	UserID  int                `bson:"user_id" json:"user_id"`
	Text    string             `bson:"text" json:"text"`
}