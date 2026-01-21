package permissions

import "github.com/moshfiq123456/ums-backend/internal/models"

func ToResponse(p models.Permission) PermissionResponse {
	return PermissionResponse{
		ID:          p.ID,
		Code:        p.Code,
		Name:        p.Name,
		Description: p.Description,
	}
}

func ToResponseList(perms []models.Permission) []PermissionResponse {
	res := make([]PermissionResponse, 0, len(perms))
	for _, p := range perms {
		res = append(res, ToResponse(p))
	}
	return res
}
