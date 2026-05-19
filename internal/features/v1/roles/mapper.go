// internal/features/v1/roles/mapper.go
package roles

import (
	"github.com/moshfiq123456/ums-backend/internal/models"
)

func toResponse(role models.Role) RoleResponse {
	return RoleResponse{
		ID:          role.ID,
		OrgID:       role.OrganizationID.String(),
		OrgName:     role.Organization.Name,
		Name:        role.Name,
		Code:        role.Code,
		Description: role.Description,
		IsActive:    role.IsActive,
		CreatedAt:   role.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   role.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}
}

func toResponseList(roles []models.Role) []RoleResponse {
	res := make([]RoleResponse, 0, len(roles))
	for _, r := range roles {
		res = append(res, toResponse(r))
	}
	return res
}
