package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type UserAuthResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	UserType string `json:"user_type"`
	OrgID    string `json:"org_id"`
	OrgSlug  string `json:"org_slug"`
}

type LoginResponse struct {
	User         UserAuthResponse `json:"user"`
	AccessToken  string           `json:"access_token"`
	RefreshToken string           `json:"refresh_token"`
	ExpiresAt    string           `json:"expires_at"`
}

type RefreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    string `json:"expires_at"`
}

type GuestRequest struct {
	OrgID string  `json:"org_id" binding:"required,uuid"`
	Email *string `json:"email"  binding:"omitempty,email"`
	Name  *string `json:"name"`
}

type GuestResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
	UserID      string `json:"user_id"`
	IsNew       bool   `json:"is_new"`
}
