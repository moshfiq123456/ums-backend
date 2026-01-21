package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	app := router.Group("/auth")
	{
		app.POST("/login", handler.Login)
		app.POST("/refresh", handler.Refresh)
		app.POST("/logout", handler.Logout)
	}
}
