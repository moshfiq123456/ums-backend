package user_hierarchy

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterModule wires repository → service → handler → routes
func RegisterModule(router *gin.Engine, db *gorm.DB) {
	// 1️⃣ Repository
	repo := NewRepository(db)

	// 2️⃣ Service
	service := NewService(repo)

	// 3️⃣ Handler
	handler := NewHandler(service)

	// 4️⃣ Register routes via separate file
	RegisterRoutes(router, handler)
}
