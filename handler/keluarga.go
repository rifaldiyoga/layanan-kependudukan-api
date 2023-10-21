package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/keluarga"
	"net/http"

	"github.com/gin-gonic/gin"
)

type keluargaHandler struct {
	keluargaService keluarga.Service
	authService     auth.Service
}

func NewKeluargaHandler(keluargaService keluarga.Service, authService auth.Service) *keluargaHandler {
	return &keluargaHandler{keluargaService, authService}
}

func (h *keluargaHandler) CreateKeluarga(c *gin.Context) {
	var input keluarga.CreateKeluargaInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create keluarga", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newkeluarga, err := h.keluargaService.CreateKeluarga(input)
	if err != nil {
		response := helper.APIResponse("Failed create keluarga", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := keluarga.FormatKeluarga(newkeluarga)
	response := helper.APIResponse("Success create keluarga", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *keluargaHandler) UpdateKeluarga(c *gin.Context) {
	var inputID keluarga.GetKeluargaDetailInput
	var inputData keluarga.CreateKeluargaInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update keluarga", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update keluarga", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newkeluarga, err := h.keluargaService.UpdateKeluarga(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update keluarga", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := keluarga.FormatKeluarga(newkeluarga)
	response := helper.APIResponse("Success Update keluarga", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *keluargaHandler) DeleteKeluarga(c *gin.Context) {
	var inputID keluarga.GetKeluargaDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete keluarga", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.keluargaService.DeleteKeluarga(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete keluarga", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete keluarga", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *keluargaHandler) GetKeluargas(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.keluargaService.GetKeluargas(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get Keluarga", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	Keluargas, _ := pagination.Data.([]keluarga.Keluarga)
	pagination.Data = keluarga.FormatKeluargas(Keluargas)

	response := helper.APIResponse("Success get Keluarga", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)

}

func (h *keluargaHandler) GetKeluarga(c *gin.Context) {
	var inputID keluarga.GetKeluargaDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get Keluarga", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKeluarga, err := h.keluargaService.GetKeluargaByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get Keluarga", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := keluarga.FormatKeluarga(newKeluarga)
	response := helper.APIResponse("Success Get Keluarga", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
