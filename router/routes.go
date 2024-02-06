package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Create a group for the API
	v1 := router.Group("/api/v1")

	v1.GET("/openings", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "List of openings",
		})
	})

	v1.POST("/openings", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Create a new opening",
		})
	})

	v1.DELETE("/openings/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Delete a one opening",
		})
	})

	v1.PUT("/openings/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Update a one opening",
		})
	})

	v1.GET("/openings/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Get a one opening",
		})
	})
}
