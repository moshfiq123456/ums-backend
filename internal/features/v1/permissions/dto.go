package permissions

type CreatePermissionRequest struct {
	Code        string `json:"code" validate:"required,lowercase"`
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description"`
}

type UpdatePermissionRequest struct {
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description"`
}

type PermissionResponse struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
