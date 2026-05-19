package user_permissions

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/middleware"
)

func RegisterRoutes(router gin.IRouter, handler *Handler) {
	protected := router.Group("/users-permissions")
	protected.Use(middleware.JWTAuth(os.Getenv("ACCESS_TOKEN_SECRET")), middleware.RequireOrg())
	{
		protected.GET("", handler.ListAll)
		protected.POST("/:id/permissions", handler.AssignPermissions)
		protected.DELETE("/:id/permissions", handler.RemovePermissions)
		protected.GET("/:id/permissions", handler.ListPermissions)
	}
}
