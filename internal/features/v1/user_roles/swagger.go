package user_roles

// @Summary      List all user-role assignments
// @Tags         User Roles
// @Security     OAuth2Password
// @Produce      json
// @Param        org_id   query     string  false  "Filter by org ID"
// @Param        search   query     string  false  "Search"
// @Param        user_id  query     string  false  "Filter by user ID"
// @Param        role_id  query     string  false  "Filter by role ID"
// @Param        page     query     int     false  "Page"
// @Param        size     query     int     false  "Size"
// @Success      200  {object}  map[string]interface{}
// @Router       /users-roles [get]
func listAllUserRolesDoc() {}

// @Summary      Assign roles to a user
// @Tags         User Roles
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        id    path      string              true  "User ID"
// @Param        body  body      AssignRolesRequest  true  "Role IDs"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /users-roles/{id}/roles [post]
func assignRolesDoc() {}

// @Summary      Remove roles from a user
// @Tags         User Roles
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        id    path      string              true  "User ID"
// @Param        body  body      RemoveRolesRequest  true  "Role IDs"
// @Success      200   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users-roles/{id}/roles [delete]
func removeRolesDoc() {}

// @Summary      List roles of a user
// @Tags         User Roles
// @Security     OAuth2Password
// @Produce      json
// @Param        id    path      string  true   "User ID"
// @Param        page  query     int     false  "Page"
// @Param        size  query     int     false  "Size"
// @Success      200   {object}  map[string]interface{}
// @Router       /users-roles/{id}/roles [get]
func listUserRolesDoc() {}
