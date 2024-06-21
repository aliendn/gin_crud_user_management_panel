package main

import (
	"gin_crud_user_management_panel/internal/routes"
	"gin_crud_user_management_panel/pkg/database"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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

	url := ginSwagger.URL("http://localhost:8080/swagger/index.html")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run(":8080")

	// Register routes
	routes.RegisterRoutes(r, db)

	// Start the server
	serverPort := viper.GetString("server.port")
	if err := r.Run(serverPort); err != nil {
		log.Fatalf("Error starting server, %s", err)
	}
}
