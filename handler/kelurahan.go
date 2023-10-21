package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/kelurahan"
	"net/http"

	"github.com/gin-gonic/gin"
)

type kelurahanHandler struct {
	kelurahanService kelurahan.Service
	authService      auth.Service
}

func NewKelurahanHandler(kelurahanService kelurahan.Service, authService auth.Service) *kelurahanHandler {
	return &kelurahanHandler{kelurahanService, authService}
}

func (h *kelurahanHandler) CreateKelurahan(c *gin.Context) {
	var input kelurahan.CreateKelurahanInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create Kelurahan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKelurahan, err := h.kelurahanService.CreateKelurahan(input)
	if err != nil {
		response := helper.APIResponse("Failed create Kelurahan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := kelurahan.FormatKelurahan(newKelurahan)
	response := helper.APIResponse("Success create Kelurahan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *kelurahanHandler) UpdateKelurahan(c *gin.Context) {
	var inputID kelurahan.GetKelurahanDetailInput
	var inputData kelurahan.CreateKelurahanInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update Kelurahan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update Kelurahan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKelurahan, err := h.kelurahanService.UpdateKelurahan(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update Kelurahan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := kelurahan.FormatKelurahan(newKelurahan)
	response := helper.APIResponse("Success Update Kelurahan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *kelurahanHandler) DeleteKelurahan(c *gin.Context) {
	var inputID kelurahan.GetKelurahanDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete Kelurahan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.kelurahanService.DeleteKelurahan(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete Kelurahan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete Kelurahan", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *kelurahanHandler) GetKelurahans(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.kelurahanService.GetKelurahans(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get kelurahan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	kelurahans, _ := pagination.Data.([]kelurahan.Kelurahan)
	pagination.Data = kelurahan.FormatKelurahans(kelurahans)

	response := helper.APIResponse("Success get kelurahan", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)

}

func (h *kelurahanHandler) GetKelurahan(c *gin.Context) {
	var inputID kelurahan.GetKelurahanDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get Kelurahan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKelurahan, err := h.kelurahanService.GetKelurahanByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get Kelurahan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := kelurahan.FormatKelurahan(newKelurahan)
	response := helper.APIResponse("Success Get Kelurahan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
