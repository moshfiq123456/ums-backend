package app

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moshfiq123456/ums-backend/internal/config"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	cfg    *config.Config
	db     *gorm.DB
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:5173",
			"http://localhost:4200",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true, // 🔥 REQUIRED FOR COOKIES
		MaxAge:           12 * time.Hour,
	}))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return &Server{
		router: router,
		cfg:    cfg,
		db:     db,
	}
}

func (s *Server) Start(registerRoutes func(*gin.Engine, *gorm.DB)) {
	registerRoutes(s.router, s.db)

	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"port":   s.cfg.Port,
		})
	})

	log.Printf("🚀 Server starting on port %d...\n", s.cfg.Port)
	if err := s.router.Run(fmt.Sprintf(":%d", s.cfg.Port)); err != nil {
		log.Fatal(err)
	}
}
