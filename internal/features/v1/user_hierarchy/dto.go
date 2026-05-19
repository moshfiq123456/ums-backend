package user_hierarchy

import "github.com/google/uuid"

type HierarchyFilter struct {
	OrgID    string `form:"org_id"`
	Search   string `form:"search"`
	ParentID string `form:"parent_id"`
	ChildID  string `form:"child_id"`
}

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

type HierarchyDetail struct {
	OrgID       string `json:"org_id"`
	OrgName     string `json:"org_name"`
	ParentID    string `json:"parent_id"`
	ParentName  string `json:"parent_name"`
	ParentEmail string `json:"parent_email"`
	ChildID     string `json:"child_id"`
	ChildName   string `json:"child_name"`
	ChildEmail  string `json:"child_email"`
}
