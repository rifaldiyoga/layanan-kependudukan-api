package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/subdistrict"
	"net/http"

	"github.com/gin-gonic/gin"
)

type subdistrictHandler struct {
	subdistrictService subdistrict.Service
	authService        auth.Service
}

func NewSubDistrictHandler(subdistrictService subdistrict.Service, authService auth.Service) *subdistrictHandler {
	return &subdistrictHandler{subdistrictService, authService}
}

func (h *subdistrictHandler) CreateSubDistrict(c *gin.Context) {
	var input subdistrict.CreateSubDistrictInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create subdistrict", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newsubdistrict, err := h.subdistrictService.CreateSubDistrict(input)
	if err != nil {
		response := helper.APIResponse("Failed create subdistrict", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := subdistrict.FormatSubDistrict(newsubdistrict)
	response := helper.APIResponse("Success create subdistrict", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *subdistrictHandler) UpdateSubDistrict(c *gin.Context) {
	var inputID subdistrict.GetSubDistrictDetailInput
	var inputData subdistrict.CreateSubDistrictInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update subdistrict", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update subdistrict", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newsubdistrict, err := h.subdistrictService.UpdateSubDistrict(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update subdistrict", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := subdistrict.FormatSubDistrict(newsubdistrict)
	response := helper.APIResponse("Success Update subdistrict", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *subdistrictHandler) DeleteSubDistrict(c *gin.Context) {
	var inputID subdistrict.GetSubDistrictDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete subdistrict", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.subdistrictService.DeleteSubDistrict(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete subdistrict", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete subdistrict", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *subdistrictHandler) GetSubDistricts(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.subdistrictService.GetSubDistricts(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get subdistrict", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	subDistricts, _ := pagination.Data.([]subdistrict.SubDistrict)
	pagination.Data = subdistrict.FormatSubDistricts(subDistricts)

	response := helper.APIResponse("Success get subdistrict", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *subdistrictHandler) GetSubDistrict(c *gin.Context) {
	var inputID subdistrict.GetSubDistrictDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get SubDistrict", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSubDistrict, err := h.subdistrictService.GetSubDistrictByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get SubDistrict", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := subdistrict.FormatSubDistrict(newSubDistrict)
	response := helper.APIResponse("Success Get SubDistrict", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
