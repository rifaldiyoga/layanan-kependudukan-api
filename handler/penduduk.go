package handler

import (
	"fmt"
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pendudukHandler struct {
	pendudukService penduduk.Service
	authService     auth.Service
}

func NewPendudukHandler(pendudukService penduduk.Service, authService auth.Service) *pendudukHandler {
	return &pendudukHandler{pendudukService, authService}
}

func (h *pendudukHandler) CreatePenduduk(c *gin.Context) {
	var input penduduk.CreatePendudukInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		fmt.Print(err.Error())
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create penduduk", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newpenduduk, err := h.pendudukService.CreatePenduduk(input)
	if err != nil {
		response := helper.APIResponse("Failed create penduduk", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := penduduk.FormatPenduduk(newpenduduk)
	response := helper.APIResponse("Success create penduduk", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pendudukHandler) UpdatePenduduk(c *gin.Context) {
	var inputID penduduk.GetPendudukDetailInput
	var inputData penduduk.CreatePendudukInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update penduduk", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update penduduk", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newpenduduk, err := h.pendudukService.UpdatePenduduk(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update penduduk", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := penduduk.FormatPenduduk(newpenduduk)
	response := helper.APIResponse("Success Update penduduk", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pendudukHandler) DeletePenduduk(c *gin.Context) {
	var inputID penduduk.GetPendudukDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete penduduk", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.pendudukService.DeletePenduduk(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete penduduk", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete penduduk", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *pendudukHandler) GetPenduduks(c *gin.Context) {
	nik := c.Query("no_kk")

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.pendudukService.GetPenduduks(pagination, nik)
	if err != nil {
		response := helper.APIResponse("Failed get Penduduk", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	Penduduks, _ := pagination.Data.([]penduduk.Penduduk)
	pagination.Data = penduduk.FormatPenduduks(Penduduks)

	response := helper.APIResponse("Success get Penduduk", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)

}

func (h *pendudukHandler) GetPenduduk(c *gin.Context) {
	var inputID penduduk.GetPendudukDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get Penduduk", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPenduduk, err := h.pendudukService.GetPendudukByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get Penduduk", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := penduduk.FormatPenduduk(newPenduduk)
	response := helper.APIResponse("Success Get Penduduk", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
