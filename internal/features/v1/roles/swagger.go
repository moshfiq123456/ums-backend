package roles

// @Summary      Create a role
// @Tags         Roles
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        body  body      CreateRoleRequest  true  "Role payload"
// @Success      201   {object}  RoleResponse
// @Failure      500   {object}  map[string]string
// @Router       /roles [post]
func createRoleDoc() {}

// @Summary      List roles
// @Tags         Roles
// @Security     OAuth2Password
// @Produce      json
// @Param        org_id     query     string  false  "Filter by org ID"
// @Param        search     query     string  false  "Search"
// @Param        is_active  query     bool    false  "Filter by active status"
// @Param        page       query     int     false  "Page"
// @Param        size       query     int     false  "Size"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]string
// @Router       /roles [get]
func listRolesDoc() {}

// @Summary      Get a role
// @Tags         Roles
// @Security     OAuth2Password
// @Produce      json
// @Param        roleId  path      int  true  "Role ID"
// @Success      200     {object}  RoleResponse
// @Failure      404     {object}  map[string]string
// @Router       /roles/{roleId} [get]
func getRoleDoc() {}

// @Summary      Update a role
// @Tags         Roles
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        roleId  path      int                true  "Role ID"
// @Param        body    body      UpdateRoleRequest  true  "Update payload"
// @Success      200     {object}  RoleResponse
// @Failure      500     {object}  map[string]string
// @Router       /roles/{roleId} [put]
func updateRoleDoc() {}

// @Summary      Update role status
// @Tags         Roles
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        roleId  path      int                      true  "Role ID"
// @Param        body    body      UpdateRoleStatusRequest  true  "Status payload"
// @Success      200     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /roles/{roleId}/status [patch]
func setRoleStatusDoc() {}

// @Summary      Delete a role
// @Tags         Roles
// @Security     OAuth2Password
// @Produce      json
// @Param        roleId  path      int  true  "Role ID"
// @Success      200     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /roles/{roleId} [delete]
func deleteRoleDoc() {}
