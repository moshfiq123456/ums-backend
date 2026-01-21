package user_hierarchy

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/middleware"
)

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	protected := router.Group("/users/hierarchy")
	protected.Use(middleware.JWTAuth(os.Getenv("ACCESS_TOKEN_SECRET")))
	{
		protected.POST("/:parentId/children", handler.AssignChild)
		protected.DELETE("/:parentId/children/:childId", handler.RemoveChild)

		protected.GET("/:userId/children", handler.GetChildren)
		protected.GET("/:userId/parent", handler.GetParent)

		protected.POST("/check", handler.CheckHierarchy)
	}
}
