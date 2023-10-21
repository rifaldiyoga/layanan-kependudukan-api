package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/rw"
	"net/http"

	"github.com/gin-gonic/gin"
)

type rwHandler struct {
	rwService   rw.Service
	authService auth.Service
}

func NewRWHandler(rwService rw.Service, authService auth.Service) *rwHandler {
	return &rwHandler{rwService, authService}
}

func (h *rwHandler) CreateRW(c *gin.Context) {
	var input rw.CreateRWInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create rw", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newrw, err := h.rwService.CreateRW(input)
	if err != nil {
		response := helper.APIResponse("Failed create rw", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := rw.FormatRW(newrw)
	response := helper.APIResponse("Success create rw", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *rwHandler) UpdateRW(c *gin.Context) {
	var inputID rw.GetRWDetailInput
	var inputData rw.CreateRWInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update rw", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update rw", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newrw, err := h.rwService.UpdateRW(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update rw", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := rw.FormatRW(newrw)
	response := helper.APIResponse("Success Update rw", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *rwHandler) DeleteRW(c *gin.Context) {
	var inputID rw.GetRWDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete rw", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.rwService.DeleteRW(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete rw", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete rw", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *rwHandler) GetRWs(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.rwService.GetRWs(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get rw", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	rws, _ := pagination.Data.([]rw.RW)
	pagination.Data = rw.FormatRWs(rws)

	response := helper.APIResponse("Success get rw", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *rwHandler) GetRW(c *gin.Context) {
	var inputID rw.GetRWDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get RW", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newRW, err := h.rwService.GetRWByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get RW", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := rw.FormatRW(newRW)
	response := helper.APIResponse("Success Get RW", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
