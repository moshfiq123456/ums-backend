package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/config"
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
