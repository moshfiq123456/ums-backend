package auth

import (
	"context"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/moshfiq123456/ums-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// Service handles auth logic
type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// Login authenticates a user and returns access token + refresh token
func (s *Service) Login(ctx context.Context, req LoginRequest) (LoginResponse, string, error) {
	user, err := s.repo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return LoginResponse{}, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return LoginResponse{}, "", err
	}

	sessionID := uuid.New()
	session := models.LoginSession{
		ID:               sessionID,
		UserID:           user.ID,
		RefreshExpiresAt: time.Now().Add(parseDuration(os.Getenv("REFRESH_TOKEN_TTL"))),
	}

	if err := s.repo.CreateSession(ctx, session); err != nil {
		return LoginResponse{}, "", err
	}

	accessToken, accessExp, err := GenerateAccessToken(user.ID)
	if err != nil {
		return LoginResponse{}, "", err
	}

	refreshToken, _, err := GenerateRefreshToken(user.ID, sessionID)
	if err != nil {
		return LoginResponse{}, "", err
	}

	resp := LoginResponse{
		User:        toUserAuthResponse(user),
		AccessToken: accessToken,
		ExpiresAt:   accessExp.Format(time.RFC3339),
	}

	return resp, refreshToken, nil
}

// Refresh validates refresh token and returns new access + refresh token
func (s *Service) Refresh(ctx context.Context, refreshToken string) (RefreshResponse, string, error) {
	claims, err := ParseRefreshToken(refreshToken)
	if err != nil {
		return RefreshResponse{}, "", err
	}

	session, err := s.repo.FindSessionByID(ctx, claims.SessionID)
	if err != nil {
		return RefreshResponse{}, "", err
	}

	accessToken, accessExp, err := GenerateAccessToken(claims.UserID)
	if err != nil {
		return RefreshResponse{}, "", err
	}

	newRefreshToken, _, err := GenerateRefreshToken(claims.UserID, session.ID)
	if err != nil {
		return RefreshResponse{}, "", err
	}

	resp := RefreshResponse{
		AccessToken: accessToken,
		ExpiresAt:   accessExp.Format(time.RFC3339),
	}

	return resp, newRefreshToken, nil
}

// Logout invalidates a refresh token
func (s *Service) Logout(ctx context.Context, refreshToken string) error {
	claims, err := ParseRefreshToken(refreshToken)
	if err != nil {
		return err
	}
	return s.repo.LogoutSession(ctx, claims.SessionID)
}

// Helper
func parseDuration(str string) time.Duration {
	if str == "" {
		return 7 * 24 * time.Hour
	}
	d, err := time.ParseDuration(str)
	if err != nil {
		return 7 * 24 * time.Hour
	}
	return d
}
