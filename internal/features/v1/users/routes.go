package users

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/middleware"
)

// RegisterRoutes registers all user routes
func RegisterRoutes(router *gin.Engine, handler *Handler) {

	// 🔓 Public route
	router.POST("/users", handler.CreateUser)
	
	// 🔐 Protected routes
	protected := router.Group("/users")
	protected.Use(middleware.JWTAuth(os.Getenv("ACCESS_TOKEN_SECRET")))
	{
		protected.GET("", handler.ListUsers)
		protected.GET("/:id", handler.GetUser)
		protected.PUT("/:id", handler.UpdateUser)
		protected.DELETE("/:id", handler.DeleteUser)
		protected.PATCH("/:id/status", handler.SetStatus)
		protected.PATCH("/:id/password", handler.ChangePassword)
		protected.PATCH("/:id/avatar", handler.UploadAvatar)
		protected.DELETE("/:id/avatar", handler.RemoveAvatar)
	}
}
