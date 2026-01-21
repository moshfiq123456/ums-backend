package auth

import (
	"time"

	"github.com/moshfiq123456/ums-backend/internal/models"
)

func toLoginResponse(
	user models.User,
	accessToken string,
	refreshToken string,
	expiresAt time.Time,
) LoginResponse {
	return LoginResponse{
		User:         toUserAuthResponse(user),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    expiresAt.Format(time.RFC3339),
	}
}

func toUserAuthResponse(user models.User) UserAuthResponse {
	return UserAuthResponse{
		ID:     user.ID.String(),
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}
}
