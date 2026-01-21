package user_hierarchy

import "github.com/google/uuid"

type AssignChildRequest struct {
	ChildUserID uuid.UUID `json:"child_user_id" validate:"required"`
}

type CheckHierarchyRequest struct {
	ParentUserID uuid.UUID `json:"parent_user_id" validate:"required"`
	ChildUserID  uuid.UUID `json:"child_user_id" validate:"required"`
}

type CheckHierarchyResponse struct {
	IsRelated bool `json:"is_related"`
}
