package sessions

// @Summary      List active sessions
// @Tags         Sessions
// @Security     OAuth2Password
// @Produce      json
// @Param        org_id  query     string  false  "Filter by org ID"
// @Param        search  query     string  false  "Search by name or email"
// @Param        page    query     int     false  "Page"
// @Param        size    query     int     false  "Size"
// @Success      200     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]string
// @Router       /sessions/active [get]
func listActiveDoc() {}

// @Summary      Force logout a session
// @Tags         Sessions
// @Security     OAuth2Password
// @Produce      json
// @Param        id   path      string  true  "Session ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /sessions/{id} [delete]
func forceLogoutDoc() {}
