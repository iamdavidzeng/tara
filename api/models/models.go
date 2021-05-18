package models

import "github.com/iamdavidzeng/tara/api/services"

type (
	Users struct {
		services.GormBase
		Email    string
		Phone    string
		Password string
	}

	Posts struct {
		services.GormBase
		UserID    int
		Title     string
		Content   string
		Published bool
	}
)
