package sessions

type ActiveSessionResponse struct {
	SessionID   string  `json:"session_id"`
	OrgID       string  `json:"org_id"`
	OrgName     string  `json:"org_name"`
	UserID      string  `json:"user_id"`
	UserName    string  `json:"user_name"`
	UserEmail   string  `json:"user_email"`
	UserType    string  `json:"user_type"`
	IPAddress   *string `json:"ip_address,omitempty"`
	UserAgent   *string `json:"user_agent,omitempty"`
	LoggedInAt  string  `json:"logged_in_at"`
	ExpiresAt   string  `json:"expires_at"`
}

type ActiveSessionFilter struct {
	OrgID  string `form:"org_id"`
	Search string `form:"search"`
}
