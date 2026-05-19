package organizations

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterModule(router gin.IRouter, db *gorm.DB) {
	repo    := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)
	RegisterRoutes(router, handler)
}
