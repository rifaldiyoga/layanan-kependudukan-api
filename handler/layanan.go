package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"net/http"

	"github.com/gin-gonic/gin"
)

type layananHandler struct {
	layananService layanan.Service
	authService    auth.Service
}

func NewLayananHandler(layananService layanan.Service, authService auth.Service) *layananHandler {
	return &layananHandler{layananService, authService}
}

func (h *layananHandler) CreateLayanan(c *gin.Context) {
	var input layanan.CreateLayananInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create layanan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newlayanan, err := h.layananService.CreateLayanan(input)
	if err != nil {
		response := helper.APIResponse("Failed create layanan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := layanan.FormatLayanan(newlayanan)
	response := helper.APIResponse("Success create layanan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *layananHandler) UpdateLayanan(c *gin.Context) {
	var inputID layanan.GetLayananDetailInput
	var inputData layanan.CreateLayananInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update layanan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update layanan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newlayanan, err := h.layananService.UpdateLayanan(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update layanan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := layanan.FormatLayanan(newlayanan)
	response := helper.APIResponse("Success Update layanan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *layananHandler) DeleteLayanan(c *gin.Context) {
	var inputID layanan.GetLayananDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete layanan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.layananService.DeleteLayanan(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete layanan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete layanan", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *layananHandler) GetLayanans(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.layananService.GetLayanansPaging(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get Layanan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	Layanans, _ := pagination.Data.([]layanan.Layanan)
	pagination.Data = layanan.FormatLayanans(Layanans)

	response := helper.APIResponse("Success get Layanan", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)

}

func (h *layananHandler) GetLayanansGrouped(c *gin.Context) {

	types, err := h.layananService.GetTypes()
	if err != nil {
		response := helper.APIResponse("Failed get types", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	layanans, err := h.layananService.GetLayanans()
	if err != nil {
		response := helper.APIResponse("Failed get layanan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	layananFormatter := layanan.FormatLayanans(layanans)
	typeFormatter := layanan.FormatTypes(types, layananFormatter)

	response := helper.APIResponse("Success get layanan", http.StatusOK, "success", typeFormatter)
	c.JSON(http.StatusOK, response)

}

func (h *layananHandler) GetRekomLayanans(c *gin.Context) {

	layanans, err := h.layananService.GetRekomLayanans()
	if err != nil {
		response := helper.APIResponse("Failed get layanan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	layananFormatter := layanan.FormatLayanans(layanans)

	response := helper.APIResponse("Success get layanan", http.StatusOK, "success", layananFormatter)
	c.JSON(http.StatusOK, response)

}

func (h *layananHandler) GetLayanan(c *gin.Context) {
	var inputID layanan.GetLayananDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get Layanan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newLayanan, err := h.layananService.GetLayananByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get Layanan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := layanan.FormatLayanan(newLayanan)
	response := helper.APIResponse("Success Get Layanan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
