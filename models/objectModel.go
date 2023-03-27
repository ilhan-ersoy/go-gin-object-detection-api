package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Object struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name"`
	User_id    *string            `bson:"user_id"`
	Image      *string            `json:"image"`
	Labels     []string           `json:"labels"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}
