package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

/*
|--------------------------------------------------------------------------
| Claims
|--------------------------------------------------------------------------
*/

// Access token claims
type AccessClaims struct {
	UserID uuid.UUID `json:"sub"`
	jwt.RegisteredClaims
}

// Refresh token claims
type RefreshClaims struct {
	UserID    uuid.UUID `json:"sub"`
	SessionID uuid.UUID `json:"sid"`
	jwt.RegisteredClaims
}

/*
|--------------------------------------------------------------------------
| Token Generators
|--------------------------------------------------------------------------
*/

// GenerateAccessToken creates a short-lived JWT access token
func GenerateAccessToken(userID uuid.UUID) (string, time.Time, error) {
	ttl, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_TTL"))
	if err != nil {
		return "", time.Time{}, err
	}

	expiresAt := time.Now().Add(ttl)

	claims := AccessClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
	signed, err := token.SignedString(secret)
	if err != nil {
		return "", time.Time{}, err
	}

	return signed, expiresAt, nil
}

// GenerateRefreshToken creates a long-lived refresh token
func GenerateRefreshToken(userID, sessionID uuid.UUID) (string, time.Time, error) {
	ttl, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_TTL"))
	if err != nil {
		return "", time.Time{}, err
	}

	expiresAt := time.Now().Add(ttl)

	claims := RefreshClaims{
		UserID:    userID,
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
	signed, err := token.SignedString(secret)
	if err != nil {
		return "", time.Time{}, err
	}

	return signed, expiresAt, nil
}

/*
|--------------------------------------------------------------------------
| Token Parsers / Validators
|--------------------------------------------------------------------------
*/

// ParseAccessToken validates access token and returns claims
func ParseAccessToken(tokenString string) (*AccessClaims, error) {
	claims := &AccessClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid access token")
	}

	return claims, nil
}

// ParseRefreshToken validates refresh token and returns claims
func ParseRefreshToken(tokenString string) (*RefreshClaims, error) {
	claims := &RefreshClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	return claims, nil
}
