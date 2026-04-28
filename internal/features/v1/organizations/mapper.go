package organizations

import "github.com/moshfiq123456/ums-backend/internal/models"

func toResponse(org models.Organization, userCount, roleCount int64) OrgResponse {
	return OrgResponse{
		ID:        org.ID.String(),
		Name:      org.Name,
		Slug:      org.Slug,
		Domain:    org.Domain,
		LogoURL:   org.LogoURL,
		IsActive:  org.IsActive,
		Plan:      org.Plan,
		CreatedAt: org.CreatedAt,
		UpdatedAt: org.UpdatedAt,
		UserCount: userCount,
		RoleCount: roleCount,
	}
}

func toResponseList(orgs []models.Organization) []OrgResponse {
	res := make([]OrgResponse, 0, len(orgs))
	for _, o := range orgs {
		res = append(res, toResponse(o, 0, 0))
	}
	return res
}
