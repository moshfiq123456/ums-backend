// @title           UMS Backend API
// @version         1.0
// @description     User Management System — centralized identity and access management.
// @termsOfService  http://swagger.io/terms/

// @contact.name   UMS Support
// @contact.email  support@ums.com

// @host      localhost:8080
// @BasePath  /api/v1/ums

// @securityDefinitions.oauth2.password OAuth2Password
// @tokenUrl /api/v1/ums/auth/token

package main

import (
	"fmt"
	"log"

	_ "github.com/moshfiq123456/ums-backend/docs"
	"github.com/moshfiq123456/ums-backend/internal/app"
	"github.com/moshfiq123456/ums-backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
	)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	log.Println("✅ Database connected")

	server := app.NewServer(cfg, db)
	server.Start(app.RegisterRoutes) // <-- pass the Gin route registrar
}
