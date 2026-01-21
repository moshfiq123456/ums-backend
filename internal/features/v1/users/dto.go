package users

import "time"

type CreateUserRequest struct {
	Name     string  `json:"name" binding:"required,min=2"`
	Email    string  `json:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required,min=8"`
	Phone    *string `json:"phone"`
}

type UpdateUserRequest struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
}

type UpdateStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=active inactive blocked"`
}

type UserResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     *string    `json:"phone,omitempty"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
