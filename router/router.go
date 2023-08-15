package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Get the port from the environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Run the server
	router.Run(":" + port)
}
