package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/rt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type rtHandler struct {
	rtService   rt.Service
	authService auth.Service
}

func NewRTHandler(rtService rt.Service, authService auth.Service) *rtHandler {
	return &rtHandler{rtService, authService}
}

func (h *rtHandler) CreateRT(c *gin.Context) {
	var input rt.CreateRTInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create rt", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newrt, err := h.rtService.CreateRT(input)
	if err != nil {
		response := helper.APIResponse("Failed create rt", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := rt.FormatRT(newrt)
	response := helper.APIResponse("Success create rt", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *rtHandler) UpdateRT(c *gin.Context) {
	var inputID rt.GetRTDetailInput
	var inputData rt.CreateRTInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update rt", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update rt", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newrt, err := h.rtService.UpdateRT(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update rt", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := rt.FormatRT(newrt)
	response := helper.APIResponse("Success Update rt", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *rtHandler) DeleteRT(c *gin.Context) {
	var inputID rt.GetRTDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete rt", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.rtService.DeleteRT(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete rt", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete rt", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *rtHandler) GetRTs(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.rtService.GetRTs(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get rt", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	rts, _ := pagination.Data.([]rt.RT)
	pagination.Data = rt.FormatRTs(rts)

	response := helper.APIResponse("Success get rt", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)

}

func (h *rtHandler) GetRT(c *gin.Context) {
	var inputID rt.GetRTDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get RT", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newRT, err := h.rtService.GetRTByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get RT", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := rt.FormatRT(newRT)
	response := helper.APIResponse("Success Get RT", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
