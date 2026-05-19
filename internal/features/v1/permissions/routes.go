package permissions

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/middleware"
)

func RegisterRoutes(router gin.IRouter, handler *Handler) {
	protected := router.Group("/permissions")
	protected.Use(middleware.JWTAuth(os.Getenv("ACCESS_TOKEN_SECRET")), middleware.RequireOrg())
	{
		protected.POST("", handler.Create)
		protected.GET("", handler.List)
		protected.GET("/:id", handler.Get)
		protected.PUT("/:id", handler.Update)
		protected.DELETE("/:id", handler.Delete)
	}
}
