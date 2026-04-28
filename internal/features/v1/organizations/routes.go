package organizations

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/middleware"
)

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	protected := router.Group("/organizations")
	protected.Use(middleware.JWTAuth(os.Getenv("ACCESS_TOKEN_SECRET")))
	{
		protected.POST("", handler.Create)
		protected.GET("", handler.List)
		protected.GET("/:id", handler.Get)
		protected.PUT("/:id", handler.Update)
		protected.PATCH("/:id/status", handler.SetStatus)
		protected.DELETE("/:id", handler.Delete)
	}
}
