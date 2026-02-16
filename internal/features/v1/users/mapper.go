package users

import "github.com/moshfiq123456/ums-backend/internal/models"

// toResponse: convert models.User -> UserResponse (basic, used for list/create/update)
func toResponse(user models.User) UserResponse {
	return UserResponse{
		ID:          user.ID.String(),
		Name:        user.Name,
		Email:       user.Email,
		Phone:       user.Phone,
		Status:      user.Status,
		Roles:       []UserRoleItem{},
		Permissions: []UserPermissionItem{},
		Hierarchy:   UserHierarchyInfo{Children: []HierarchyUserItem{}},
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

// toDetailResponse: convert models.User with preloaded relations + hierarchy into full UserResponse
func toDetailResponse(user models.User, parent *models.User, children []models.User) UserResponse {
	roles := make([]UserRoleItem, 0, len(user.Roles))
	for _, r := range user.Roles {
		roles = append(roles, UserRoleItem{
			ID:   r.ID,
			Name: r.Name,
			Code: r.Code,
		})
	}

	permissions := make([]UserPermissionItem, 0, len(user.Permissions))
	for _, up := range user.Permissions {
		permissions = append(permissions, UserPermissionItem{
			ID:    up.Permission.ID,
			Code:  up.Permission.Code,
			Name:  up.Permission.Name,
			Allow: up.Allow,
		})
	}

	var parentItem *HierarchyUserItem
	if parent != nil {
		parentItem = &HierarchyUserItem{
			ID:   parent.ID.String(),
			Name: parent.Name,
		}
	}

	childItems := make([]HierarchyUserItem, 0, len(children))
	for _, c := range children {
		childItems = append(childItems, HierarchyUserItem{
			ID:   c.ID.String(),
			Name: c.Name,
		})
	}

	return UserResponse{
		ID:          user.ID.String(),
		Name:        user.Name,
		Email:       user.Email,
		Phone:       user.Phone,
		Status:      user.Status,
		Roles:       roles,
		Permissions: permissions,
		Hierarchy: UserHierarchyInfo{
			Parent:   parentItem,
			Children: childItems,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func toResponseList(users []models.User) []UserResponse {
	res := make([]UserResponse, 0, len(users))
	for _, u := range users {
		res = append(res, toResponse(u))
	}
	return res
}
