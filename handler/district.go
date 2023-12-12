package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/district"
	"layanan-kependudukan-api/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type districtHandler struct {
	districtService district.Service
	authService     auth.Service
}

func NewDistrictHandler(districtService district.Service, authService auth.Service) *districtHandler {
	return &districtHandler{districtService, authService}
}

func (h *districtHandler) CreateDistrict(c *gin.Context) {
	var input district.CreateDistrictInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create district", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newdistrict, err := h.districtService.CreateDistrict(input)
	if err != nil {
		response := helper.APIResponse("Failed create district", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := district.FormatDistrict(newdistrict)
	response := helper.APIResponse("Success create district", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *districtHandler) UpdateDistrict(c *gin.Context) {
	var inputID district.GetDistrictDetailInput
	var inputData district.CreateDistrictInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update district", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update district", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newdistrict, err := h.districtService.UpdateDistrict(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update district", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := district.FormatDistrict(newdistrict)
	response := helper.APIResponse("Success Update district", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *districtHandler) DeleteDistrict(c *gin.Context) {
	var inputID district.GetDistrictDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete district", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.districtService.DeleteDistrict(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete district", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete district", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *districtHandler) GetDistricts(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	provinceId := 0
	intVar, errConv := strconv.Atoi(c.Query("province_id"))
	if errConv == nil {
		provinceId = intVar
	}

	pagination, err := h.districtService.GetDistricts(pagination, provinceId)
	if err != nil {
		response := helper.APIResponse("Failed get district", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	districts, _ := pagination.Data.([]district.District)
	pagination.Data = district.FormatDistricts(districts)

	response := helper.APIResponse("Success get district", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *districtHandler) GetDistrict(c *gin.Context) {
	var inputID district.GetDistrictDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get District", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDistrict, err := h.districtService.GetDistrictByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get District", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := district.FormatDistrict(newDistrict)
	response := helper.APIResponse("Success Get District", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
