package models

import "time"

type (
	Posts struct {
		Title     string
		Content   string
		UserID    int
		Published bool
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt time.Time
	}
)
