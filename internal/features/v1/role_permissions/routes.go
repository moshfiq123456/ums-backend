package role_permissions

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/middleware"
)

func RegisterRoutes(router gin.IRouter, handler *Handler) {
	protected := router.Group("/roles")
	protected.Use(middleware.JWTAuth(os.Getenv("ACCESS_TOKEN_SECRET")), middleware.RequireOrg())
	{
		permissions := protected.Group("/:roleId/permissions")
		{
			permissions.POST("", handler.Assign)
			permissions.POST("/bulk", handler.BulkAssign)
			permissions.GET("", handler.List)
			permissions.DELETE("/:permissionId", handler.Remove)
		}
	}
}
