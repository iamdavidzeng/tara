package models

import "github.com/jinzhu/gorm"

// models: Users
type (
	Users struct {
		gorm.Model
		Email    string
		Phone    string
		Password string
	}
)