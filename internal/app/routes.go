package app

import (
	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/auth"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/organizations"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/permissions"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/role_permissions"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/roles"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/sessions"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/user_hierarchy"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/user_permissions"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/user_roles"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/users"
	"gorm.io/gorm"
)

// RegisterRoutes wires all feature modules into the Gin router
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	v1 := router.Group("/api/v1/ums")

	users.RegisterModule(v1, db)
	auth.RegisterModule(v1, db)
	sessions.RegisterModule(v1, db)
	organizations.RegisterModule(v1, db)
	user_roles.RegisterModule(v1, db)
	user_permissions.RegisterModule(v1, db)
	user_hierarchy.RegisterModule(v1, db)
	role_permissions.RegisterModule(v1, db)
	roles.RegisterModule(v1, db)
	permissions.RegisterModule(v1, db)
}
