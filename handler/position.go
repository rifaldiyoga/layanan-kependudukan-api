package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/position"
	"net/http"

	"github.com/gin-gonic/gin"
)

type positionHandler struct {
	positionService position.Service
	authService     auth.Service
}

func NewPositionHandler(positionService position.Service, authService auth.Service) *positionHandler {
	return &positionHandler{positionService, authService}
}

func (h *positionHandler) CreatePosition(c *gin.Context) {
	var input position.CreatePositionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create Position", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPosition, err := h.positionService.CreatePosition(input)
	if err != nil {
		response := helper.APIResponse("Failed create Position", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := position.FormatPosition(newPosition)
	response := helper.APIResponse("Success create Position", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *positionHandler) UpdatePosition(c *gin.Context) {
	var inputID position.GetPositionDetailInput
	var inputData position.CreatePositionInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update Position", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update Position", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPosition, err := h.positionService.UpdatePosition(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update Position", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := position.FormatPosition(newPosition)
	response := helper.APIResponse("Success Update Position", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *positionHandler) DeletePosition(c *gin.Context) {
	var inputID position.GetPositionDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete Position", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.positionService.DeletePosition(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete Position", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete Position", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *positionHandler) GetPositions(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.positionService.GetPositions(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get position", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	positions, _ := pagination.Data.([]position.Position)
	pagination.Data = position.FormatPositions(positions)

	response := helper.APIResponse("Success get position", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)

}
