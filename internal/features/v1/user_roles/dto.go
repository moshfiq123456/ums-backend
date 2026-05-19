package user_roles

type UserRoleFilter struct {
	OrgID  string `form:"org_id"`
	Search string `form:"search"`
	UserID string `form:"user_id"`
	RoleID string `form:"role_id"`
}

type AssignRolesRequest struct {
	RoleIDs []uint `json:"role_ids" validate:"required,min=1"`
}

type RemoveRolesRequest struct {
	RoleIDs []uint `json:"role_ids" validate:"required,min=1"`
}

type UserRoleResponse struct {
	UserID string `json:"user_id"`
	RoleID uint   `json:"role_id"`
	Role   string `json:"role_name"`
}

type UserRoleDetail struct {
	OrgID         string  `json:"org_id"`
	OrgName       string  `json:"org_name"`
	UserID        string  `json:"user_id"`
	UserName      string  `json:"user_name"`
	UserEmail     string  `json:"user_email"`
	UserAvatarURL *string `json:"user_avatar_url,omitempty"`
	RoleID        uint    `json:"role_id"`
	RoleName      string  `json:"role_name"`
	RoleCode      string  `json:"role_code"`
}
