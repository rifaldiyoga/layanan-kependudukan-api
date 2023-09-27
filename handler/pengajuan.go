package handler

import (
	"fmt"
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/detail_pengajuan"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pengajuanHandler struct {
	pengajuanService       pengajuan.Service
	detailPengajuanService detail_pengajuan.Service
	authService            auth.Service
}

func NewpengajuanHandler(pengajuanService pengajuan.Service, detailPengajuanService detail_pengajuan.Service, authService auth.Service) *pengajuanHandler {
	return &pengajuanHandler{pengajuanService, detailPengajuanService, authService}
}

func (h *pengajuanHandler) CreatePengajuan(c *gin.Context) {
	var input pengajuan.CreatePengajuanInput

	fmt.Println(input)
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create pengajuan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser, _ := c.Get("currentUser")
	userObject := currentUser.(user.User)

	newPengajuan, err := h.pengajuanService.CreatePengajuan(input, userObject)
	if err != nil {
		response := helper.APIResponse("Failed create pengajuan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	status := "PENDING"
	if input.Status != "" {
		status = input.Status
	}

	_, err = h.detailPengajuanService.CreateDetailPengajuan(newPengajuan.ID, status, userObject)
	if err != nil {
		response := helper.APIResponse("Failed create pengajuan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newPengajuan, err = h.pengajuanService.GetPengajuanByID(newPengajuan.ID)
	if err != nil {
		response := helper.APIResponse("Failed get pengajuan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := pengajuan.FormatPengajuan(newPengajuan)
	response := helper.APIResponse("Success create pengajuan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pengajuanHandler) UpdatePengajuan(c *gin.Context) {
	var inputID pengajuan.GetPengajuanDetailInput
	var inputData pengajuan.CreatePengajuanInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update pengajuan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update pengajuan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newpengajuan, err := h.pengajuanService.UpdatePengajuan(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update pengajuan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// detailPengajuans, err := h.detailPengajuanService.GetDetailByPengajuan(newpengajuan.ID)
	// if err != nil {
	// 	response := helper.APIResponse("Failed create pengajuan", http.StatusBadRequest, "error", nil)
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	formatter := pengajuan.FormatPengajuan(newpengajuan)
	response := helper.APIResponse("Success Update pengajuan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pengajuanHandler) DeletePengajuan(c *gin.Context) {
	var inputID pengajuan.GetPengajuanDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete pengajuan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.pengajuanService.DeletePengajuan(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete pengajuan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete pengajuan", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *pengajuanHandler) GetPengajuans(c *gin.Context) {
	currentUser, _ := c.Get("currentUser")
	userObject := currentUser.(user.User)

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.pengajuanService.GetPengajuans(pagination, userObject)
	if err != nil {
		response := helper.APIResponse("Failed get pengajuan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pengajuans, _ := pagination.Data.([]pengajuan.Pengajuan)
	pagination.Data = pengajuan.FormatPengajuans(pengajuans)

	response := helper.APIResponse("Success get pengajuan", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)

}
