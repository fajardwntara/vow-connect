package main

import (
	"log"
	"net/http"

	"github.com/fajardwntara/vow-connect/api/config"
	"github.com/fajardwntara/vow-connect/api/domain/user"
	"github.com/fajardwntara/vow-connect/api/routes"
	"github.com/fajardwntara/vow-connect/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, _ := config.LoadConfig()
	database.ConnectDB(cfg)

	router := gin.Default()

	// Tes Ping
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong!!",
		})
	})

	db := database.GetDB()
	// Init Repo
	userRepo := user.NewUserRepository(db)

	// Routers of Repo
	routes.UserRouteRegistry(router, userRepo)

	// Starting server
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
