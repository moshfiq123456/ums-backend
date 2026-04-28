package users

import "time"

type UserFilter struct {
	OrgID  string `form:"-"` // set programmatically from JWT, not from query param
	Search string `form:"search"`
	Status string `form:"status"`
}

type CreateUserRequest struct {
	Name     string  `json:"name"     binding:"required,min=2,max=50"`
	Email    string  `json:"email"    binding:"required,email"`
	Password string  `json:"password" binding:"required,min=8,max=32"`
	Phone    *string `json:"phone"    binding:"omitempty,e164"`
	UserType string  `json:"user_type" binding:"omitempty,oneof=admin customer member"`
}


type UpdateUserRequest struct {
	Name  *string `json:"name" binding:"omitempty,min=2,max=50"`
	Phone *string `json:"phone" binding:"omitempty,e164"`
}

type UpdateStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=active inactive blocked"`
}

type ChangePasswordRequest struct {
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type UserResponse struct {
	ID        string     `json:"id"`
	OrgID     string     `json:"org_id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     *string    `json:"phone,omitempty"`
	AvatarURL *string    `json:"avatar_url,omitempty"`
	UserType  string     `json:"user_type"`
	Status    string     `json:"status"`
	Roles     []RoleInfo `json:"roles"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type RoleInfo struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
