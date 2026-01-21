package user_roles

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
