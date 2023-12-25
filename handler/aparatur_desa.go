package handler

import (
	aparaturDesa "layanan-kependudukan-api/aparatur_desa"
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type aparaturDesaHandler struct {
	aparaturDesaService aparaturDesa.Service
	authService         auth.Service
}

func NewAparaturDesaHandler(aparaturDesaService aparaturDesa.Service, authService auth.Service) *aparaturDesaHandler {
	return &aparaturDesaHandler{aparaturDesaService, authService}
}

func (h *aparaturDesaHandler) CreateAparaturDesa(c *gin.Context) {
	var input aparaturDesa.CreateAparaturDesaInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create aparaturDesa", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newaparaturDesa, err := h.aparaturDesaService.CreateAparaturDesa(input)
	if err != nil {
		response := helper.APIResponse("Failed create aparaturDesa", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := aparaturDesa.FormatAparaturDesa(newaparaturDesa)
	response := helper.APIResponse("Success create aparaturDesa", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *aparaturDesaHandler) UpdateAparaturDesa(c *gin.Context) {
	var inputID aparaturDesa.GetAparaturDesaDetailInput
	var inputData aparaturDesa.CreateAparaturDesaInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update aparaturDesa", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update aparaturDesa", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newaparaturDesa, err := h.aparaturDesaService.UpdateAparaturDesa(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update aparaturDesa", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := aparaturDesa.FormatAparaturDesa(newaparaturDesa)
	response := helper.APIResponse("Success Update aparaturDesa", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *aparaturDesaHandler) DeleteAparaturDesa(c *gin.Context) {
	var inputID aparaturDesa.GetAparaturDesaDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete aparaturDesa", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.aparaturDesaService.DeleteAparaturDesa(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete aparaturDesa", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete aparaturDesa", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *aparaturDesaHandler) GetAparaturDesas(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.aparaturDesaService.GetAparaturDesas(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get aparaturDesa", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	aparaturDesas, _ := pagination.Data.([]aparaturDesa.AparaturDesa)
	pagination.Data = aparaturDesa.FormatAparaturDesas(aparaturDesas)

	response := helper.APIResponse("Success get aparaturDesa", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *aparaturDesaHandler) GetAparaturDesa(c *gin.Context) {
	var inputID aparaturDesa.GetAparaturDesaDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get AparaturDesa", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newAparaturDesa, err := h.aparaturDesaService.GetAparaturDesaByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get AparaturDesa", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := aparaturDesa.FormatAparaturDesa(newAparaturDesa)
	response := helper.APIResponse("Success Get AparaturDesa", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
