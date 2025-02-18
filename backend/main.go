package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tobratland/wishlist/backend/config"
	"github.com/tobratland/wishlist/backend/models"
	"github.com/tobratland/wishlist/backend/routes"
)

func main() {
	// Initialize configuration
	config.Init()

	// Connect to the database
	models.ConnectDatabase()

	// Migrate the schema
	models.Migrate()

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":8080")
}
