package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mnzsss/go-jobs/docs"
	"github.com/mnzsss/go-jobs/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()

	basePath := "/api/v1"

	// Set the base path for the API
	docs.SwaggerInfo.BasePath = basePath

	// Create a group for the API
	v1 := router.Group(basePath)

	{ // Openings routes - /api/v1/openings
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.DELETE("/openings/:id", handler.DeleteOpeningHandler)
		v1.PUT("/openings/:id", handler.UpdateOpeningHandler)
		v1.GET("/openings/:id", handler.GetOpeningHandler)
	}

	// Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
