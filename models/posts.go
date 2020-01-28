package models

import "time"

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Posts struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		Title     string             `bson:"title"`
		Content   string             `bson:"content"`
		UserID    int                `bson:"user_id"`
		Published bool               `bson:"published"`
		CreatedAt time.Time          `bson:"created_at"`
		UpdatedAt time.Time          `bson:"updated_at"`
	}
)
