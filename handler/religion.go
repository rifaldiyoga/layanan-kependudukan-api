package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/religion"
	"net/http"

	"github.com/gin-gonic/gin"
)

type religionHandler struct {
	religionService religion.Service
	authService     auth.Service
}

func NewReligionHandler(religionService religion.Service, authService auth.Service) *religionHandler {
	return &religionHandler{religionService, authService}
}

func (h *religionHandler) CreateReligion(c *gin.Context) {
	var input religion.CreateReligionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create religion", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newReligion, err := h.religionService.CreateReligion(input)
	if err != nil {
		response := helper.APIResponse("Failed create religion", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := religion.FormatReligion(newReligion)
	response := helper.APIResponse("Success create religion", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *religionHandler) UpdateReligion(c *gin.Context) {
	var inputID religion.GetReligionDetailInput
	var inputData religion.CreateReligionInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update religion", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update religion", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newReligion, err := h.religionService.UpdateReligion(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update religion", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := religion.FormatReligion(newReligion)
	response := helper.APIResponse("Success Update religion", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *religionHandler) DeleteReligion(c *gin.Context) {
	var inputID religion.GetReligionDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete religion", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.religionService.DeleteReligion(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete religion", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete religion", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *religionHandler) GetReligions(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.religionService.GetReligions(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get religion", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	religions, _ := pagination.Data.([]religion.Religion)
	pagination.Data = religion.FormatReligions(religions)

	response := helper.APIResponse("Success get religion", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *religionHandler) GetReligion(c *gin.Context) {
	var inputID religion.GetReligionDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get religion", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newReligion, err := h.religionService.GetReligionByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get religion", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := religion.FormatReligion(newReligion)
	response := helper.APIResponse("Success Get religion", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
