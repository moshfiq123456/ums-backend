package user_permissions

// @Summary      List all user-permission assignments
// @Tags         User Permissions
// @Security     OAuth2Password
// @Produce      json
// @Param        org_id         query     string  false  "Filter by org ID"
// @Param        search         query     string  false  "Search"
// @Param        user_id        query     string  false  "Filter by user ID"
// @Param        permission_id  query     string  false  "Filter by permission ID"
// @Param        page           query     int     false  "Page"
// @Param        size           query     int     false  "Size"
// @Success      200  {object}  map[string]interface{}
// @Router       /users-permissions [get]
func listAllUserPermissionsDoc() {}

// @Summary      Assign permissions to a user
// @Tags         User Permissions
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        id    path      string                    true  "User ID"
// @Param        body  body      AssignPermissionsRequest  true  "Permission IDs"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /users-permissions/{id}/permissions [post]
func assignPermissionsDoc() {}

// @Summary      Remove permissions from a user
// @Tags         User Permissions
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        id    path      string                    true  "User ID"
// @Param        body  body      RemovePermissionsRequest  true  "Permission IDs"
// @Success      200   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users-permissions/{id}/permissions [delete]
func removePermissionsDoc() {}

// @Summary      List permissions of a user
// @Tags         User Permissions
// @Security     OAuth2Password
// @Produce      json
// @Param        id    path      string  true   "User ID"
// @Param        page  query     int     false  "Page"
// @Param        size  query     int     false  "Size"
// @Success      200   {object}  map[string]interface{}
// @Router       /users-permissions/{id}/permissions [get]
func listUserPermissionsDoc() {}
