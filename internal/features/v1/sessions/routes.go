package sessions

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/middleware"
)

func RegisterRoutes(router gin.IRouter, handler *Handler) {
	protected := router.Group("/sessions")
	protected.Use(middleware.JWTAuth(os.Getenv("ACCESS_TOKEN_SECRET")), middleware.RequireOrg())
	{
		protected.GET("/active", handler.ListActive)
		protected.DELETE("/:id", handler.ForceLogout)
	}
}
