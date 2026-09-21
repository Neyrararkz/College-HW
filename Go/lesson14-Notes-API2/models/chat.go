package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserIDs []int              `bson:"user_ids" json:"user_ids"`
}