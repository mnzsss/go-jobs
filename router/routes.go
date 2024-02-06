package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mnzsss/go-jobs/handler"
)

func initializeRoutes(router *gin.Engine) {
	// Create a group for the API
	v1 := router.Group("/api/v1")

	// Openings routes - /api/v1/openings
	v1.GET("/openings", handler.ListOpeningsHandler)
	v1.POST("/openings", handler.CreateOpeningHandler)
	v1.DELETE("/openings/:id", handler.DeleteOpeningHandler)
	v1.PUT("/openings/:id", handler.UpdateOpeningHandler)
	v1.GET("/openings/:id", handler.GetOpeningHandler)
}
