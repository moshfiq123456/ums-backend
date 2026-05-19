package auth

// @Summary      Login
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      LoginRequest  true  "Login credentials"
// @Success      200   {object}  LoginResponse
// @Failure      400   {object}  map[string]string
// @Failure      401   {object}  map[string]string
// @Router       /auth/login [post]
func loginDoc() {}

// @Summary      Issue a guest token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body  body      GuestRequest   true  "Guest request — org_id required, email + name optional"
// @Success      200   {object}  GuestResponse
// @Failure      400   {object}  map[string]string
// @Router       /auth/guest [post]
func guestDoc() {}

// @Summary      Refresh access token
// @Tags         Auth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /auth/refresh [post]
func refreshDoc() {}

// @Summary      Logout
// @Tags         Auth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /auth/logout [post]
func logoutDoc() {}
