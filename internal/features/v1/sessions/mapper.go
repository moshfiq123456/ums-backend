package sessions

import (
	"time"

	"github.com/moshfiq123456/ums-backend/internal/models"
)

func toResponse(s models.LoginSession) ActiveSessionResponse {
	return ActiveSessionResponse{
		SessionID:  s.ID.String(),
		OrgID:      s.User.OrganizationID.String(),
		OrgName:    s.User.Organization.Name,
		UserID:     s.UserID.String(),
		UserName:   s.User.Name,
		UserEmail:  s.User.Email,
		UserType:   s.User.UserType,
		IPAddress:  s.IPAddress,
		UserAgent:  s.UserAgent,
		LoggedInAt: s.LoggedInAt.Format(time.RFC3339),
		ExpiresAt:  s.RefreshExpiresAt.Format(time.RFC3339),
	}
}
