package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mnzsss/go-jobs/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	// Get all openings
	if err := db.Find(&openings).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	SendSuccess(ctx, "list", openings)
}

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	// Bind the request
	ctx.BindJSON(&request)

	// Validate the request
	if err := request.Validate(); err != nil {
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
		SendError(ctx, http.StatusInternalServerError, err.Error())

		return
	}

	SendSuccess(ctx, "create", opening)
}

func GetOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	// Validate the id
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "param").Error())
		return
	}

	// Find the opening
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	SendSuccess(ctx, "get", opening)
}

func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	// Validate the id
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "param").Error())
		return
	}

	request := UpdateOpeningRequest{}

	// Bind the request
	ctx.BindJSON(&request)

	// Validate the request
	if err := request.Validate(); err != nil {
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Find the opening
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	// Update the opening
	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Description != "" {
		opening.Description = request.Description
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	// Save the opening
	if err := db.Save(&opening).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	SendSuccess(ctx, "update", opening)
}

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	// Validate the id
	if id == "" {
		SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "param").Error())
		return
	}

	// Find the opening
	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	// Delete the opening
	if err := db.Delete(&opening).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening: %v", err.Error()))
		return
	}

	// Send success
	SendSuccess(ctx, "delete", opening)
}
