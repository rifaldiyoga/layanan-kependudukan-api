package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/berpergian"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type berpergianHandler struct {
	berpergianService berpergian.Service
	layananService    layanan.Service
	pengajuanHandler  pengajuanHandler
	authService       auth.Service
}

func NewBerpergianHandler(berpergianService berpergian.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *berpergianHandler {
	return &berpergianHandler{berpergianService, layananService, pengajuanHandler, authService}
}

func (h *berpergianHandler) CreateBerpergian(c *gin.Context) {
	var input berpergian.CreateBerpergianInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create berpergian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentLayanan, err := h.layananService.GetLayananByCode("SKBBK")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create berpergian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, _ := c.FormFile("lampiran")

	if file != nil {
		path := helper.FormatFileName(file.Filename)
		filePath := filepath.Join("documents/berpergian", path)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.Lampiran = filePath
	}

	input.NIK = currentUser.Nik
	newBerpergian, err := h.berpergianService.CreateBerpergian(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create berpergian", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newBerpergian.ID
	inputPengajuan.Keterangan = newBerpergian.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := berpergian.FormatBerpergian(newBerpergian)
	response := helper.APIResponse("Success create berpergian", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *berpergianHandler) UpdateBerpergian(c *gin.Context) {
	var inputID berpergian.GetBerpergianDetailInput
	var inputData berpergian.CreateBerpergianInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update berpergian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update berpergian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newBerpergian, err := h.berpergianService.UpdateBerpergian(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update berpergian", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := berpergian.FormatBerpergian(newBerpergian)
	response := helper.APIResponse("Success Update berpergian", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *berpergianHandler) DeleteBerpergian(c *gin.Context) {
	var inputID berpergian.GetBerpergianDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete berpergian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.berpergianService.DeleteBerpergian(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete berpergian", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete berpergian", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *berpergianHandler) GetBerpergians(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.berpergianService.GetBerpergians(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get berpergian", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	berpergians, _ := pagination.Data.([]berpergian.Berpergian)
	pagination.Data = berpergian.FormatBerpergians(berpergians)

	response := helper.APIResponse("Success get berpergian", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *berpergianHandler) GetBerpergian(c *gin.Context) {
	var inputID berpergian.GetBerpergianDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get berpergian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newBerpergian, err := h.berpergianService.GetBerpergianByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get berpergian", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := berpergian.FormatBerpergian(newBerpergian)
	response := helper.APIResponse("Success Get berpergian", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
