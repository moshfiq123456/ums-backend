package permissions

type PermissionFilter struct {
	OrgID  string `form:"org_id"`
	Search string `form:"search"`
}

type CreatePermissionRequest struct {
	OrgID       string `json:"org_id"       binding:"omitempty,uuid"`
	Code        string `json:"code"         validate:"required,lowercase"`
	Name        string `json:"name"         validate:"required,min=3"`
	Description string `json:"description"`
}

type UpdatePermissionRequest struct {
	Name        string `json:"name" validate:"required,min=3"`
	Description string `json:"description"`
}

type PermissionResponse struct {
	ID          uint   `json:"id"`
	OrgID       string `json:"org_id"`
	OrgName     string `json:"org_name"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
