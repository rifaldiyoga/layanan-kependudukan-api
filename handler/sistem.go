package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/sistem"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sistemHandler struct {
	sistemService sistem.Service
	authService   auth.Service
}

func NewSistemHandler(sistemService sistem.Service, authService auth.Service) *sistemHandler {
	return &sistemHandler{sistemService, authService}
}

func (h *sistemHandler) CreateSistem(c *gin.Context) {
	var input sistem.CreateSistemInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create sistem", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSistem, err := h.sistemService.CreateSistem(input)
	if err != nil {
		response := helper.APIResponse("Failed create sistem", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := sistem.FormatSistem(newSistem)
	response := helper.APIResponse("Success create sistem", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *sistemHandler) UpdateSistem(c *gin.Context) {
	var inputID sistem.GetSistemDetailInput
	var inputData sistem.CreateSistemInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update sistem", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update sistem", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSistem, err := h.sistemService.UpdateSistem(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update sistem", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := sistem.FormatSistem(newSistem)
	response := helper.APIResponse("Success Update sistem", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *sistemHandler) DeleteSistem(c *gin.Context) {
	var inputID sistem.GetSistemDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete sistem", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.sistemService.DeleteSistem(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete sistem", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete sistem", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *sistemHandler) GetSistems(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.sistemService.GetSistems(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get sistem", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	sistems, _ := pagination.Data.([]sistem.Sistem)
	pagination.Data = sistem.FormatSistems(sistems)

	response := helper.APIResponse("Success get sistem", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *sistemHandler) GetSistem(c *gin.Context) {
	var inputID sistem.GetSistemDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get sistem", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSistem, err := h.sistemService.GetSistemByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get sistem", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := sistem.FormatSistem(newSistem)
	response := helper.APIResponse("Success Get sistem", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
