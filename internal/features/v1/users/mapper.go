package users

import "github.com/moshfiq123456/ums-backend/internal/models"

// Mapper: convert models.User -> UserResponse
func toResponse(user models.User) UserResponse {
	roles := make([]RoleInfo, 0, len(user.Roles))
	for _, r := range user.Roles {
		roles = append(roles, RoleInfo{ID: r.ID, Name: r.Name, Code: r.Code})
	}
	return UserResponse{
		ID:        user.ID.String(),
		OrgID:     user.OrganizationID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		AvatarURL: user.AvatarURL,
		UserType:  user.UserType,
		Status:    user.Status,
		Roles:     roles,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func toResponseList(users []models.User) []UserResponse {
	res := make([]UserResponse, 0, len(users))
	for _, u := range users {
		res = append(res, toResponse(u))
	}
	return res
}
