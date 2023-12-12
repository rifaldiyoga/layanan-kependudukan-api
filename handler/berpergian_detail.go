package handler

import (
	"layanan-kependudukan-api/auth"
	berpergianDetail "layanan-kependudukan-api/berpergian_detail"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type berpergianDetailHandler struct {
	berpergianDetailService berpergianDetail.Service
	authService             auth.Service
}

func NewBerpergianDetailHandler(berpergianDetailService berpergianDetail.Service, authService auth.Service) *berpergianDetailHandler {
	return &berpergianDetailHandler{berpergianDetailService, authService}
}

func (h *berpergianDetailHandler) CreateBerpergianDetail(c *gin.Context) {
	var input berpergianDetail.CreateBerpergianDetailInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create BerpergianDetail", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newBerpergianDetail, err := h.berpergianDetailService.CreateBerpergianDetail(input, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create BerpergianDetail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := berpergianDetail.FormatBerpergianDetail(newBerpergianDetail)
	response := helper.APIResponse("Success create BerpergianDetail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *berpergianDetailHandler) UpdateBerpergianDetail(c *gin.Context) {
	var inputID berpergianDetail.GetBerpergianDetailDetailInput
	var inputData berpergianDetail.CreateBerpergianDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update BerpergianDetail", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update BerpergianDetail", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newBerpergianDetail, err := h.berpergianDetailService.UpdateBerpergianDetail(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update BerpergianDetail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := berpergianDetail.FormatBerpergianDetail(newBerpergianDetail)
	response := helper.APIResponse("Success Update BerpergianDetail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *berpergianDetailHandler) DeleteBerpergianDetail(c *gin.Context) {
	var inputID berpergianDetail.GetBerpergianDetailDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete BerpergianDetail", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.berpergianDetailService.DeleteBerpergianDetail(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete BerpergianDetail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete BerpergianDetail", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *berpergianDetailHandler) GetBerpergianDetails(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.berpergianDetailService.GetBerpergianDetails(pagination)
	if err != nil {
		response := helper.APIResponse("Failed get berpergianDetail", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	berpergianDetails, _ := pagination.Data.([]berpergianDetail.BerpergianDetail)
	pagination.Data = berpergianDetail.FormatBerpergianDetails(berpergianDetails)

	response := helper.APIResponse("Success get berpergianDetail", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)

}
