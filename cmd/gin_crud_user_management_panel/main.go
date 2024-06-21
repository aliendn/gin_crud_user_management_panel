package main

import (
	"gin_crud_user_management_panel/internal/routes"
	"gin_crud_user_management_panel/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Load configuration
	viper.SetConfigFile("configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Initialize the database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database, %s", err)
	}

	// Set up Gin
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r, db)

	// Start the server
	serverPort := viper.GetString("server.port")
	if err := r.Run(serverPort); err != nil {
		log.Fatalf("Error starting server, %s", err)
	}
}
