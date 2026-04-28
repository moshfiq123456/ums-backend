package organizations

import "time"

type OrgFilter struct {
	Search   string `form:"search"`
	IsActive *bool  `form:"is_active"`
	Plan     string `form:"plan"`
}

type CreateOrgRequest struct {
	Name   string `json:"name"   binding:"required,min=2,max=200"`
	Slug   string `json:"slug"   binding:"required,min=2,max=100"`
	Domain string `json:"domain" binding:"omitempty"`
	Plan   string `json:"plan"   binding:"omitempty,oneof=free pro enterprise"`
}

type UpdateOrgRequest struct {
	Name   string `json:"name"   binding:"omitempty,min=2,max=200"`
	Domain string `json:"domain" binding:"omitempty"`
	Plan   string `json:"plan"   binding:"omitempty,oneof=free pro enterprise"`
}

type UpdateOrgStatusRequest struct {
	IsActive bool `json:"is_active"`
}

type OrgResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Domain    *string   `json:"domain,omitempty"`
	LogoURL   *string   `json:"logo_url,omitempty"`
	IsActive  bool      `json:"is_active"`
	Plan      string    `json:"plan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Stats
	UserCount int64 `json:"user_count"`
	RoleCount int64 `json:"role_count"`
}
