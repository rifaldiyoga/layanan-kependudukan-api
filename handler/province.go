package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/province"
	"net/http"

	"github.com/gin-gonic/gin"
)

type provinceHandler struct {
	provinceService province.Service
	authService     auth.Service
}

func NewProvinceHandler(provinceService province.Service, authService auth.Service) *provinceHandler {
	return &provinceHandler{provinceService, authService}
}

func (h *provinceHandler) CreateProvince(c *gin.Context) {
	var input province.CreateProvinceInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create Province", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProvince, err := h.provinceService.CreateProvince(input)
	if err != nil {
		response := helper.APIResponse("Failed create Province", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := province.FormatProvince(newProvince)
	response := helper.APIResponse("Success create Province", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *provinceHandler) UpdateProvince(c *gin.Context) {
	var inputID province.GetProvinceDetailInput
	var inputData province.CreateProvinceInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update Province", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update Province", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProvince, err := h.provinceService.UpdateProvince(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update Province", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := province.FormatProvince(newProvince)
	response := helper.APIResponse("Success Update Province", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *provinceHandler) DeleteProvince(c *gin.Context) {
	var inputID province.GetProvinceDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete Province", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.provinceService.DeleteProvince(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete Province", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete Province", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *provinceHandler) GetProvinces(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.provinceService.GetProvinces(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get province", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	provinces, _ := pagination.Data.([]province.Province)
	pagination.Data = province.FormatProvinces(provinces)

	response := helper.APIResponse("Success get province", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)

}
