package users

// @Summary      Create a user
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        body  body      CreateUserRequest  true  "User payload"
// @Success      201   {object}  UserResponse
// @Failure      400   {object}  map[string]string
// @Router       /users [post]
func createUserDoc() {}

// @Summary      List users
// @Tags         Users
// @Security     OAuth2Password
// @Produce      json
// @Param        org_id  query     string  false  "Filter by org ID"
// @Param        search  query     string  false  "Search by name or email"
// @Param        status  query     string  false  "Filter by status"
// @Param        page    query     int     false  "Page number"
// @Param        size    query     int     false  "Page size"
// @Success      200     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]string
// @Router       /users [get]
func listUsersDoc() {}

// @Summary      Get a user
// @Tags         Users
// @Security     OAuth2Password
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  UserResponse
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
func getUserDoc() {}

// @Summary      Update a user
// @Tags         Users
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        id    path      string             true  "User ID"
// @Param        body  body      UpdateUserRequest  true  "Update payload"
// @Success      200   {object}  UserResponse
// @Failure      400   {object}  map[string]string
// @Router       /users/{id} [put]
func updateUserDoc() {}

// @Summary      Delete a user
// @Tags         Users
// @Security     OAuth2Password
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /users/{id} [delete]
func deleteUserDoc() {}

// @Summary      Update user status
// @Tags         Users
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        id    path      string               true  "User ID"
// @Param        body  body      UpdateStatusRequest  true  "Status payload"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /users/{id}/status [patch]
func setUserStatusDoc() {}

// @Summary      Change user password
// @Tags         Users
// @Security     OAuth2Password
// @Accept       json
// @Produce      json
// @Param        id    path      string                 true  "User ID"
// @Param        body  body      ChangePasswordRequest  true  "Password payload"
// @Success      200   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users/{id}/password [patch]
func changePasswordDoc() {}

// @Summary      Upload user avatar
// @Tags         Users
// @Security     OAuth2Password
// @Accept       multipart/form-data
// @Produce      json
// @Param        id      path      string  true  "User ID"
// @Param        avatar  formData  file    true  "Avatar file (JPEG, PNG, WebP — max 2MB)"
// @Success      200     {object}  map[string]string
// @Failure      400     {object}  map[string]string
// @Router       /users/{id}/avatar [patch]
func uploadAvatarDoc() {}

// @Summary      Remove user avatar
// @Tags         Users
// @Security     OAuth2Password
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id}/avatar [delete]
func removeAvatarDoc() {}
