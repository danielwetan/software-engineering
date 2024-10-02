package model

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

// Request
type (
	CreateUserRequest struct {
		Email    string `json:"email" validate:"email,required"`
		Password string `json:"password" validate:"min=8,required"`
	}
)
