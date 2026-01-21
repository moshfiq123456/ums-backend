package role_permissions

type AssignPermissionRequest struct {
	PermissionID uint `json:"permission_id" validate:"required"`
}

type BulkAssignPermissionRequest struct {
	PermissionIDs []uint `json:"permission_ids" validate:"required,min=1"`
}
