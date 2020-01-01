package schemas

// Schemas: UserSchema
type (
	UserSchema struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
)