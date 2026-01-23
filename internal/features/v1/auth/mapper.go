package auth

import (
	"time"

	"github.com/moshfiq123456/ums-backend/internal/models"
)

// Converts internal User model + access token into API response
func toLoginResponse(
	user models.User,
	accessToken string,
	expiresAt time.Time,
) LoginResponse {
	return LoginResponse{
		User:        toUserAuthResponse(user),
		AccessToken: accessToken,
		ExpiresAt:   expiresAt.Format(time.RFC3339),
	}
}

// Converts User model to UserAuthResponse
func toUserAuthResponse(user models.User) UserAuthResponse {
	return UserAuthResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}
}
