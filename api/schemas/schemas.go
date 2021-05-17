package schemas

import (
	"time"
)

type (
	UserSchema struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	PostSchema struct {
		ID        uint      `json:"id"`
		UserID    int       `json:"user_id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		Published bool      `json:"published"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
