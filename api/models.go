package api

import "gorm.io/gorm"

type (
	Users struct {
		gorm.Model
		Email    string
		Phone    string
		Password string
	}

	Posts struct {
		gorm.Model
		UserID    int
		Title     string
		Content   string
		Published bool
	}
)
