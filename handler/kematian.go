package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/kematian"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type kematianHandler struct {
	kematianService  kematian.Service
	layananService   layanan.Service
	pengajuanHandler pengajuanHandler
	authService      auth.Service
}

func NewKematianHandler(kematianService kematian.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *kematianHandler {
	return &kematianHandler{kematianService, layananService, pengajuanHandler, authService}
}

func (h *kematianHandler) CreateKematian(c *gin.Context) {
	var input kematian.CreateKematianInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create kematian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentLayanan, err := h.layananService.GetLayananByCode("SKKM")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create kematian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	fileKetRs, _ := c.FormFile("lampiran_ket_rs")

	if fileKetRs != nil {
		path := helper.FormatFileName(fileKetRs.Filename)
		filePath := filepath.Join("documents/kematian", path)
		if err := c.SaveUploadedFile(fileKetRs, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.LampiranKetRs = filePath
	}

	input.NIK = currentUser.Nik
	newKematian, err := h.kematianService.CreateKematian(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create kematian", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newKematian.ID
	inputPengajuan.Keterangan = newKematian.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := kematian.FormatKematian(newKematian)
	response := helper.APIResponse("Success create kematian", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *kematianHandler) UpdateKematian(c *gin.Context) {
	var inputID kematian.GetKematianDetailInput
	var inputData kematian.CreateKematianInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update kematian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update kematian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKematian, err := h.kematianService.UpdateKematian(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update kematian", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := kematian.FormatKematian(newKematian)
	response := helper.APIResponse("Success Update kematian", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *kematianHandler) DeleteKematian(c *gin.Context) {
	var inputID kematian.GetKematianDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete kematian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.kematianService.DeleteKematian(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete kematian", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete kematian", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *kematianHandler) GetKematians(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.kematianService.GetKematians(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get kematian", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	kematians, _ := pagination.Data.([]kematian.Kematian)
	pagination.Data = kematian.FormatKematians(kematians)

	response := helper.APIResponse("Success get kematian", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *kematianHandler) GetKematian(c *gin.Context) {
	var inputID kematian.GetKematianDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get kematian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKematian, err := h.kematianService.GetKematianByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get kematian", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := kematian.FormatKematian(newKematian)
	response := helper.APIResponse("Success Get kematian", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
