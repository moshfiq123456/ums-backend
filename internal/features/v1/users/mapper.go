package users

import "github.com/moshfiq123456/ums-backend/internal/models"

// Mapper: convert models.User -> UserResponse
func toResponse(user models.User) UserResponse {
	return UserResponse{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Status:    user.Status,
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
