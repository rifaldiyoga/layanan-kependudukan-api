package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/education"
	"layanan-kependudukan-api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type educationHandler struct {
	educationService education.Service
	authService      auth.Service
}

func NewEducationHandler(educationService education.Service, authService auth.Service) *educationHandler {
	return &educationHandler{educationService, authService}
}

func (h *educationHandler) CreateEducation(c *gin.Context) {
	var input education.CreateEducationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create education", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	neweducation, err := h.educationService.CreateEducation(input)
	if err != nil {
		response := helper.APIResponse("Failed create education", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := education.FormatEducation(neweducation)
	response := helper.APIResponse("Success create education", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *educationHandler) UpdateEducation(c *gin.Context) {
	var inputID education.GetEducationDetailInput
	var inputData education.CreateEducationInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update education", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update education", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	neweducation, err := h.educationService.UpdateEducation(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update education", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := education.FormatEducation(neweducation)
	response := helper.APIResponse("Success Update education", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *educationHandler) DeleteEducation(c *gin.Context) {
	var inputID education.GetEducationDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete education", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.educationService.DeleteEducation(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete education", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete education", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *educationHandler) GetEducations(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.educationService.GetEducations(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get education", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	educations, _ := pagination.Data.([]education.Education)
	pagination.Data = education.FormatEducations(educations)

	response := helper.APIResponse("Success get education", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)

}
