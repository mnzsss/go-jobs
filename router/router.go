package router

import (
	"github.com/gin-gonic/gin"
)

func Init() {
	// Creates a gin router with default middleware:
	router := gin.Default()

	// Initialize the routes
	initializeRoutes(router)

	// Run the server on :8080
	router.Run(":8080")
}
