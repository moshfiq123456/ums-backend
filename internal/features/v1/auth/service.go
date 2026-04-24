package auth

import (
	"context"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/moshfiq123456/ums-backend/internal/models"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
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

	// CreateSession (DB write), GenerateAccessToken and GenerateRefreshToken (JWT signing)
	// are all independent — run them concurrently.
	var (
		accessToken  string
		accessExp    time.Time
		refreshToken string
	)

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return s.repo.CreateSession(gctx, session)
	})

	g.Go(func() error {
		var err error
		accessToken, accessExp, err = GenerateAccessToken(user.ID)
		return err
	})

	g.Go(func() error {
		var err error
		refreshToken, _, err = GenerateRefreshToken(user.ID, sessionID)
		return err
	})

	if err := g.Wait(); err != nil {
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

	// FindSessionByID (DB read) and GenerateAccessToken (JWT signing) only need parsed
	// claims — run them concurrently, then generate the refresh token using session.ID.
	var (
		session     models.LoginSession
		accessToken string
		accessExp   time.Time
	)

	g, gctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		session, err = s.repo.FindSessionByID(gctx, claims.SessionID)
		return err
	})

	g.Go(func() error {
		var err error
		accessToken, accessExp, err = GenerateAccessToken(claims.UserID)
		return err
	})

	if err := g.Wait(); err != nil {
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
