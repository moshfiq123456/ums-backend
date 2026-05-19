package role_permissions

// @Summary      Assign a permission to a role
// @Tags         Role Permissions
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        roleId  path      int                      true  "Role ID"
// @Param        body    body      AssignPermissionRequest  true  "Permission ID"
// @Success      200     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /roles/{roleId}/permissions [post]
func assignPermissionToRoleDoc() {}

// @Summary      Bulk assign permissions to a role
// @Tags         Role Permissions
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        roleId  path      int                          true  "Role ID"
// @Param        body    body      BulkAssignPermissionRequest  true  "Permission IDs"
// @Success      200     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /roles/{roleId}/permissions/bulk [post]
func bulkAssignPermissionsToRoleDoc() {}

// @Summary      Remove a permission from a role
// @Tags         Role Permissions
// @Security     OAuth2Password
// @Produce      json
// @Param        roleId        path      int  true  "Role ID"
// @Param        permissionId  path      int  true  "Permission ID"
// @Success      200           {object}  map[string]string
// @Failure      500           {object}  map[string]string
// @Router       /roles/{roleId}/permissions/{permissionId} [delete]
func removePermissionFromRoleDoc() {}

// @Summary      List permissions of a role
// @Tags         Role Permissions
// @Security     OAuth2Password
// @Produce      json
// @Param        roleId  path      int  true   "Role ID"
// @Param        page    query     int  false  "Page"
// @Param        size    query     int  false  "Size"
// @Success      200     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]string
// @Router       /roles/{roleId}/permissions [get]
func listRolePermissionsDoc() {}
