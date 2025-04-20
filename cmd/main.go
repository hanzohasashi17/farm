package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hanzohasashi17/farm/internal/config"
	database "github.com/hanzohasashi17/farm/internal/db"
	"github.com/hanzohasashi17/farm/internal/user"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := config.GetConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Database connecting error: %v", err)
	}
	defer db.Close()

	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("migrations failed: %v", err)
	}

	userRepo := user.NewUserRepo(db)
    userService := user.NewUserService(userRepo)
    userHandler := user.NewUserHandler(userService)

    router := gin.Default()
    userHandler.RegisterRoutes(router)

    router.Run(fmt.Sprintf(":%s", cfg.AppPort))
}