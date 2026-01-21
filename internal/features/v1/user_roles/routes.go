package user_roles

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/middleware"
)

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	protected := router.Group("/users-roles")
	protected.Use(middleware.JWTAuth(os.Getenv("ACCESS_TOKEN_SECRET")))
	{
		protected.POST("/:userId/roles", handler.AssignRoles)
		protected.DELETE("/:userId/roles", handler.RemoveRoles)
		protected.GET("/:userId/roles", handler.ListRoles)
	}
}
