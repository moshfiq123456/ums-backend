package role_permissions

import "github.com/moshfiq123456/ums-backend/internal/models"

type PermissionResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func toPermissionResponse(p models.Permission) PermissionResponse {
	return PermissionResponse{
		ID:   p.ID,
		Name: p.Name,
		Code: p.Code,
	}
}
