package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	pindahDetail "layanan-kependudukan-api/pindah_detail"
	"layanan-kependudukan-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pindahDetailHandler struct {
	pindahDetailService pindahDetail.Service
	authService         auth.Service
}

func NewPindahDetailHandler(pindahDetailService pindahDetail.Service, authService auth.Service) *pindahDetailHandler {
	return &pindahDetailHandler{pindahDetailService, authService}
}

func (h *pindahDetailHandler) CreatePindahDetail(c *gin.Context) {
	var input pindahDetail.CreatePindahDetailInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create PindahDetail", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPindahDetail, err := h.pindahDetailService.CreatePindahDetail(input, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create PindahDetail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := pindahDetail.FormatPindahDetail(newPindahDetail)
	response := helper.APIResponse("Success create PindahDetail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pindahDetailHandler) UpdatePindahDetail(c *gin.Context) {
	var inputID pindahDetail.GetPindahDetailDetailInput
	var inputData pindahDetail.CreatePindahDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update PindahDetail", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update PindahDetail", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPindahDetail, err := h.pindahDetailService.UpdatePindahDetail(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update PindahDetail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := pindahDetail.FormatPindahDetail(newPindahDetail)
	response := helper.APIResponse("Success Update PindahDetail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pindahDetailHandler) DeletePindahDetail(c *gin.Context) {
	var inputID pindahDetail.GetPindahDetailDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete PindahDetail", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.pindahDetailService.DeletePindahDetail(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete PindahDetail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete PindahDetail", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *pindahDetailHandler) GetPindahDetails(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.pindahDetailService.GetPindahDetails(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get pindahDetail", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pindahDetails, _ := pagination.Data.([]pindahDetail.PindahDetail)
	pagination.Data = pindahDetail.FormatPindahDetails(pindahDetails)

	response := helper.APIResponse("Success get pindahDetail", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)

}
