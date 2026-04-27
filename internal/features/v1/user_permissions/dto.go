package user_permissions

type UserPermissionFilter struct {
	Search       string `form:"search"`
	UserID       string `form:"user_id"`
	PermissionID string `form:"permission_id"`
}

type AssignPermissionsRequest struct {
	PermissionIDs []uint `json:"permission_ids" validate:"required,min=1"`
}

type RemovePermissionsRequest struct {
	PermissionIDs []uint `json:"permission_ids" validate:"required,min=1"`
}

type UserPermissionResponse struct {
	UserID       string `json:"user_id"`
	PermissionID uint   `json:"permission_id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}

type UserPermissionDetail struct {
	UserID        string  `json:"user_id"`
	UserName      string  `json:"user_name"`
	UserEmail     string  `json:"user_email"`
	UserAvatarURL *string `json:"user_avatar_url,omitempty"`
	PermissionID  uint    `json:"permission_id"`
	Code          string  `json:"code"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
}
