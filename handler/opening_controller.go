package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mnzsss/go-jobs/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List of openings",
	})
}

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	// Bind the request
	ctx.BindJSON(&request)

	// Validate the request
	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())

		return
	}

	// Create the opening
	opening := schemas.Opening{
		Role:        request.Role,
		Company:     request.Company,
		Location:    request.Location,
		Description: request.Description,
		Salary:      request.Salary,
		Remote:      *request.Remote,
		Link:        request.Link,
	}

	// Create the opening
	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error creating opening: %v", err.Error())

		SendError(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	SendSuccess(ctx, "create", opening)
}

func GetOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get a one opening",
	})
}

func UpdateOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Update a one opening",
	})
}

func DeleteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete a one opening",
	})
}
