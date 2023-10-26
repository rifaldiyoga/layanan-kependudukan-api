package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/status"
	"net/http"

	"github.com/gin-gonic/gin"
)

type statusHandler struct {
	statusService status.Service
	authService     auth.Service
}

func NewStatusHandler(statusService status.Service, authService auth.Service) *statusHandler {
	return &statusHandler{statusService, authService}
}

func (h *statusHandler) CreateStatus(c *gin.Context) {
	var input status.CreateStatusInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create status", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newStatus, err := h.statusService.CreateStatus(input)
	if err != nil {
		response := helper.APIResponse("Failed create status", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := status.FormatStatus(newStatus)
	response := helper.APIResponse("Success create status", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *statusHandler) UpdateStatus(c *gin.Context) {
	var inputID status.GetStatusDetailInput
	var inputData status.CreateStatusInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update status", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update status", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newStatus, err := h.statusService.UpdateStatus(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update status", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := status.FormatStatus(newStatus)
	response := helper.APIResponse("Success Update status", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *statusHandler) DeleteStatus(c *gin.Context) {
	var inputID status.GetStatusDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete status", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.statusService.DeleteStatus(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete status", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete status", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *statusHandler) GetStatuss(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.statusService.GetStatuss(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get status", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	statuss, _ := pagination.Data.([]status.Status)
	pagination.Data = status.FormatStatuss(statuss)

	response := helper.APIResponse("Success get status", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *statusHandler) GetStatus(c *gin.Context) {
	var inputID status.GetStatusDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get status", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newStatus, err := h.statusService.GetStatusByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get status", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := status.FormatStatus(newStatus)
	response := helper.APIResponse("Success Get status", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
