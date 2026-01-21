// internal/features/v1/roles/dto.go
package roles

type CreateRoleRequest struct {
	Name        string `json:"name" validate:"required,min=3"`
	Code        string `json:"code" validate:"required,min=3"`
	Description string `json:"description"`
}

type UpdateRoleRequest struct {
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description"`
}

type UpdateRoleStatusRequest struct {
	IsActive bool `json:"is_active"`
}

type RoleResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
