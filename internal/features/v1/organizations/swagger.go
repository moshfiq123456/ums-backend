package organizations

// @Summary      Create an organization
// @Tags         Organizations
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        body  body      CreateOrgRequest  true  "Org payload"
// @Success      201   {object}  OrgResponse
// @Failure      400   {object}  map[string]string
// @Router       /organizations [post]
func createOrgDoc() {}

// @Summary      List organizations
// @Tags         Organizations
// @Security     OAuth2Password
// @Produce      json
// @Param        search     query     string  false  "Search"
// @Param        is_active  query     bool    false  "Filter by active"
// @Param        plan       query     string  false  "Filter by plan"
// @Param        page       query     int     false  "Page"
// @Param        size       query     int     false  "Size"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]string
// @Router       /organizations [get]
func listOrgsDoc() {}

// @Summary      Get an organization
// @Tags         Organizations
// @Security     OAuth2Password
// @Produce      json
// @Param        id   path      string  true  "Org ID"
// @Success      200  {object}  OrgResponse
// @Failure      404  {object}  map[string]string
// @Router       /organizations/{id} [get]
func getOrgDoc() {}

// @Summary      Update an organization
// @Tags         Organizations
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        id    path      string            true  "Org ID"
// @Param        body  body      UpdateOrgRequest  true  "Update payload"
// @Success      200   {object}  OrgResponse
// @Failure      500   {object}  map[string]string
// @Router       /organizations/{id} [put]
func updateOrgDoc() {}

// @Summary      Update org status
// @Tags         Organizations
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        id    path      string                  true  "Org ID"
// @Param        body  body      UpdateOrgStatusRequest  true  "Status payload"
// @Success      200   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /organizations/{id}/status [patch]
func setOrgStatusDoc() {}

// @Summary      Delete an organization
// @Tags         Organizations
// @Security     OAuth2Password
// @Produce      json
// @Param        id   path      string  true  "Org ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /organizations/{id} [delete]
func deleteOrgDoc() {}
