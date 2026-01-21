package users

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all user routes
func RegisterRoutes(router *gin.Engine, handler *Handler) {
	api := router.Group("/users")
	{
		api.POST("", handler.CreateUser)

		api.GET("", handler.ListUsers)
		api.GET("/:id", handler.GetUser)
		api.PUT("/:id", handler.UpdateUser)
		api.DELETE("/:id", handler.DeleteUser)
		api.PATCH("/:id/status", handler.SetStatus)
	}
}
