package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/moshfiq123456/ums-backend/internal/config"
	"github.com/moshfiq123456/ums-backend/internal/middleware"
	"gorm.io/gorm"
)

// Server holds the Gin engine, config, and DB
type Server struct {
	router *gin.Engine
	cfg    *config.Config
	db     *gorm.DB
}

// NewServer initializes the server with middlewares
func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	router := gin.New()
    allowedOrigins := map[string]bool{
		"http://localhost:3000": true,
		"http://localhost:3001": true,
		"http://localhost:5173": true,
		"http://localhost:4200": true,
	}
	router.Use(middleware.OrgContext())
	router.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if allowedOrigins[origin] {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, Cache-Control, X-Requested-With")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Set-Cookie")
			c.Header("Access-Control-Max-Age", "43200")
		}
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	// Standard middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return &Server{
		router: router,
		cfg:    cfg,
		db:     db,
	}
}

// Start runs the server
func (s *Server) Start(registerRoutes func(*gin.Engine, *gorm.DB)) {
    
    // Serve uploaded files
    s.router.Static("/uploads", "./uploads")

    // Swagger UI
    s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Register all routes
    registerRoutes(s.router, s.db)

    // Health checks
    s.router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status":   "healthy",
            "port":     s.cfg.Port,
            "database": s.cfg.Database.DBName,
        })
    })

    s.router.GET("/health/db", func(c *gin.Context) {
        sqlDB, err := s.db.DB()
        if err != nil {
            c.JSON(500, gin.H{"status": "unhealthy", "error": err.Error()})
            return
        }
        if err := sqlDB.Ping(); err != nil {
            c.JSON(500, gin.H{"status": "unhealthy", "error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"status": "healthy"})
    })

    log.Printf("🚀 Server starting on port %d...\n", s.cfg.Port)
    if err := s.router.Run(fmt.Sprintf(":%d", s.cfg.Port)); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
