package permissions

// @Summary      Create a permission
// @Tags         Permissions
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        body  body      CreatePermissionRequest  true  "Permission payload"
// @Success      201   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /permissions [post]
func createPermissionDoc() {}

// @Summary      List permissions
// @Tags         Permissions
// @Security     OAuth2Password
// @Produce      json
// @Param        org_id  query     string  false  "Filter by org ID"
// @Param        search  query     string  false  "Search"
// @Param        page    query     int     false  "Page"
// @Param        size    query     int     false  "Size"
// @Success      200     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]string
// @Router       /permissions [get]
func listPermissionsDoc() {}

// @Summary      Get a permission
// @Tags         Permissions
// @Security     OAuth2Password
// @Produce      json
// @Param        id   path      int  true  "Permission ID"
// @Success      200  {object}  PermissionResponse
// @Failure      404  {object}  map[string]string
// @Router       /permissions/{id} [get]
func getPermissionDoc() {}

// @Summary      Update a permission
// @Tags         Permissions
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        id    path      int                      true  "Permission ID"
// @Param        body  body      UpdatePermissionRequest  true  "Update payload"
// @Success      200   {object}  PermissionResponse
// @Failure      500   {object}  map[string]string
// @Router       /permissions/{id} [put]
func updatePermissionDoc() {}

// @Summary      Delete a permission
// @Tags         Permissions
// @Security     OAuth2Password
// @Produce      json
// @Param        id   path      int  true  "Permission ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /permissions/{id} [delete]
func deletePermissionDoc() {}
