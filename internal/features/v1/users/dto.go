package users

import "time"

type CreateUserRequest struct {
	Name     string  `json:"name" binding:"required,min=2,max=50"`
	Email    string  `json:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required,min=8,max=32"`
	Phone    *string `json:"phone" binding:"omitempty,e164"` // E.164 phone format, optional
}


type UpdateUserRequest struct {
	Name  *string `json:"name" binding:"omitempty,min=2,max=50"`
	Phone *string `json:"phone" binding:"omitempty,e164"`
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
