package user_roles

import (
	"github.com/moshfiq123456/ums-backend/internal/models"
)

func ToUserRoleResponses(userID string, roles []models.Role) []UserRoleResponse {
	resp := make([]UserRoleResponse, 0, len(roles))
	for _, r := range roles {
		resp = append(resp, UserRoleResponse{
			UserID: userID,
			RoleID: uint(r.ID), // <-- convert int64 -> uint explicitly
			Role:   r.Name,
		})
	}
	return resp
}
