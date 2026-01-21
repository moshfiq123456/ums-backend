package roles

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/middleware"
)

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	protected := router.Group("/roles")
	protected.Use(middleware.JWTAuth(os.Getenv("ACCESS_TOKEN_SECRET")))
	{
		// CRUD
		protected.POST("/", handler.Create)
		protected.GET("/", handler.List)
		protected.GET("/:roleId", handler.Get)
		protected.PUT("/:roleId", handler.Update)
		protected.PATCH("/:roleId/status", handler.SetStatus)
		protected.DELETE("/:roleId", handler.Delete)
	}
}
