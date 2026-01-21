package auth

import (
	"context"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/moshfiq123456/ums-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// -----------------------------
// LOGIN
// -----------------------------
func (s *Service) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	user, err := s.repo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return LoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return LoginResponse{}, err
	}

	// Create DB session first
	sessionID := uuid.New()
	session := models.LoginSession{
		ID:               sessionID,
		UserID:           user.ID,
		RefreshExpiresAt: time.Now().Add(parseDuration(os.Getenv("REFRESH_TOKEN_TTL"))),
	}
	if err := s.repo.CreateSession(ctx, session); err != nil {
		return LoginResponse{}, err
	}

	// Generate tokens
	accessToken, accessExp, err := GenerateAccessToken(user.ID)
	if err != nil {
		return LoginResponse{}, err
	}

	refreshToken, _, err := GenerateRefreshToken(user.ID, sessionID)
	if err != nil {
		return LoginResponse{}, err
	}

	return toLoginResponse(user, accessToken, refreshToken, accessExp), nil
}

// -----------------------------
// REFRESH
// -----------------------------
func (s *Service) Refresh(ctx context.Context, req RefreshTokenRequest) (RefreshResponse, error) {
	claims := &RefreshClaims{}
	token, err := jwt.ParseWithClaims(req.RefreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return RefreshResponse{}, err
	}

	// Check DB session for revocation
	session, err := s.repo.FindSessionByID(ctx, claims.SessionID)
	if err != nil {
		return RefreshResponse{}, err
	}

	// Generate new access token
	accessToken, accessExp, err := GenerateAccessToken(claims.UserID)
	if err != nil {
		return RefreshResponse{}, err
	}

	// Optional: rotate refresh token
	refreshToken, _, err := GenerateRefreshToken(claims.UserID, session.ID)
	if err != nil {
		return RefreshResponse{}, err
	}

	return RefreshResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    accessExp.Format(time.RFC3339),
	}, nil
}

// -----------------------------
// LOGOUT
// -----------------------------
func (s *Service) Logout(ctx context.Context, req LogoutRequest) error {
	claims := &RefreshClaims{}
	token, err := jwt.ParseWithClaims(req.RefreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return err
	}

	return s.repo.LogoutSession(ctx, claims.SessionID)
}

// -----------------------------
// HELPERS
// -----------------------------
func parseDuration(str string) time.Duration {
	d, _ := time.ParseDuration(str)
	return d
}
