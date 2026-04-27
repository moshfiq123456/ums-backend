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
		protected.GET("", handler.ListAll)
		protected.POST("/:id/roles", handler.AssignRoles)
		protected.DELETE("/:id/roles", handler.RemoveRoles)
		protected.GET("/:id/roles", handler.ListRoles)
	}
}
