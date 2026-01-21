package app

import (
	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/features/v1/users"
	"gorm.io/gorm"
)

// RegisterRoutes wires all feature modules into the Gin router
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	users.RegisterModule(router, db)

}
