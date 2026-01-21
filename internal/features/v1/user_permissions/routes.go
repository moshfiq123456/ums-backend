package user_permissions

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/middleware"
)

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	protected := router.Group("/users-permissions")
	protected.Use(middleware.JWTAuth(os.Getenv("ACCESS_TOKEN_SECRET")))
	{
		protected.POST("/:userId/permissions", handler.AssignPermissions)
		protected.DELETE("/:userId/permissions", handler.RemovePermissions)
		protected.GET("/:userId/permissions", handler.ListPermissions)
	}
}
