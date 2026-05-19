package user_hierarchy

// @Summary      List all user hierarchy relationships
// @Tags         User Hierarchy
// @Security     OAuth2Password
// @Produce      json
// @Param        org_id     query     string  false  "Filter by org ID"
// @Param        search     query     string  false  "Search"
// @Param        parent_id  query     string  false  "Filter by parent ID"
// @Param        child_id   query     string  false  "Filter by child ID"
// @Param        page       query     int     false  "Page"
// @Param        size       query     int     false  "Size"
// @Success      200  {object}  map[string]interface{}
// @Router       /users/hierarchy [get]
func listHierarchyDoc() {}

// @Summary      Assign a child user to a parent
// @Tags         User Hierarchy
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        parentId  path      string              true  "Parent User ID"
// @Param        body      body      AssignChildRequest  true  "Child user ID"
// @Success      200       {object}  map[string]string
// @Failure      400       {object}  map[string]string
// @Router       /users/hierarchy/{parentId}/children [post]
func assignChildDoc() {}

// @Summary      Remove a child from a parent
// @Tags         User Hierarchy
// @Security     OAuth2Password
// @Produce      json
// @Param        parentId  path      string  true  "Parent User ID"
// @Param        childId   path      string  true  "Child User ID"
// @Success      200       {object}  map[string]string
// @Failure      500       {object}  map[string]string
// @Router       /users/hierarchy/{parentId}/children/{childId} [delete]
func removeChildDoc() {}

// @Summary      Get children of a user
// @Tags         User Hierarchy
// @Security     OAuth2Password
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  []map[string]interface{}
// @Failure      500  {object}  map[string]string
// @Router       /users/hierarchy/{id}/children [get]
func getChildrenDoc() {}

// @Summary      Get parent of a user
// @Tags         User Hierarchy
// @Security     OAuth2Password
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]string
// @Router       /users/hierarchy/{id}/parent [get]
func getParentDoc() {}

// @Summary      Check if two users are related
// @Tags         User Hierarchy
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        body  body      CheckHierarchyRequest   true  "Parent and child IDs"
// @Success      200   {object}  CheckHierarchyResponse
// @Failure      500   {object}  map[string]string
// @Router       /users/hierarchy/check [post]
func checkHierarchyDoc() {}
