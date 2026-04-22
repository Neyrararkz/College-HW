package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID         primitive.ObjectID `bson:"id.omitempty" json:"id"`
	RoomID     string             `bson:"room_id" json:"room_id"`
	SenderID   int                `bson:"sender_id" json:"sender_id"`
	SenderName string             `bson:"sender_name" json:"sender_name"`
	Text       string             `bson:"text" json:"text"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
}
