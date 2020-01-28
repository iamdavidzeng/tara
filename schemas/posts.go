package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostSchema struct {
	ID        primitive.ObjectID `json:"id"`
	UserID    int                `json:"user_id"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	Published bool               `json:"published"`
	UpdatedAt time.Time          `json:"updated_at"`
}
