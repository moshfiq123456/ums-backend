package app

import (
	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/auth"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/permissions"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/role_permissions"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/roles"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/user_hierarchy"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/user_permissions"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/user_roles"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/users"
	"gorm.io/gorm"
)

// RegisterRoutes wires all feature modules into the Gin router
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	users.RegisterModule(router, db)
	auth.RegisterModule(router, db)
	user_roles.RegisterModule(router, db)
	user_permissions.RegisterModule(router, db)
	user_hierarchy.RegisterModule(router, db)
	role_permissions.RegisterModule(router, db)
	roles.RegisterModule(router, db)
	permissions.RegisterModule(router, db)
}
