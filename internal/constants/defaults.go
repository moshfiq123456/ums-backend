package constants

import "github.com/google/uuid"

// DefaultOrgID is the fixed UUID for the built-in "Default" organization.
// Users and roles created without an explicit org context are placed here.
var DefaultOrgID = uuid.MustParse("00000000-0000-0000-0000-000000000001")
