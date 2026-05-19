package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router gin.IRouter, handler *Handler) {
	app := router.Group("/auth")
	{
		app.POST("/login", handler.Login)
		app.POST("/guest", handler.Guest)
		app.POST("/token", handler.Token)
		app.POST("/refresh", handler.Refresh)
		app.POST("/logout", handler.Logout)
	}
}
