package user_permissions

import "github.com/moshfiq123456/ums-backend/internal/models"

func ToUserPermissionResponses(userID string, permissions []models.Permission) []UserPermissionResponse {
	resp := make([]UserPermissionResponse, 0, len(permissions))
	for _, p := range permissions {
		resp = append(resp, UserPermissionResponse{
			UserID:       userID,
			PermissionID: p.ID,
			Code:         p.Code,
			Name:         p.Name,
			Description:  p.Description,
		})
	}
	return resp
}
