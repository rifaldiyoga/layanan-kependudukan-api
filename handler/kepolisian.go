package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/kepolisian"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type kepolisianHandler struct {
	kepolisianService kepolisian.Service
	layananService    layanan.Service
	pengajuanHandler  pengajuanHandler
	authService       auth.Service
}

func NewKepolisianHandler(kepolisianService kepolisian.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *kepolisianHandler {
	return &kepolisianHandler{kepolisianService, layananService, pengajuanHandler, authService}
}

func (h *kepolisianHandler) CreateKepolisian(c *gin.Context) {
	var input kepolisian.CreateKepolisianInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create kepolisian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentLayanan, err := h.layananService.GetLayananByCode("SPCK")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create kepolisian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, _ := c.FormFile("lampiran")

	if file != nil {
		path := helper.FormatFileName(file.Filename)
		filePath := filepath.Join("documents/kepolisian", path)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.Lampiran = filePath
	}

	input.NIK = currentUser.Nik
	newKepolisian, err := h.kepolisianService.CreateKepolisian(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create kepolisian", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newKepolisian.ID
	inputPengajuan.Keterangan = newKepolisian.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := kepolisian.FormatKepolisian(newKepolisian)
	response := helper.APIResponse("Success create kepolisian", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *kepolisianHandler) UpdateKepolisian(c *gin.Context) {
	var inputID kepolisian.GetKepolisianDetailInput
	var inputData kepolisian.CreateKepolisianInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update kepolisian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update kepolisian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKepolisian, err := h.kepolisianService.UpdateKepolisian(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update kepolisian", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := kepolisian.FormatKepolisian(newKepolisian)
	response := helper.APIResponse("Success Update kepolisian", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *kepolisianHandler) DeleteKepolisian(c *gin.Context) {
	var inputID kepolisian.GetKepolisianDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete kepolisian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.kepolisianService.DeleteKepolisian(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete kepolisian", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete kepolisian", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *kepolisianHandler) GetKepolisians(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.kepolisianService.GetKepolisians(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get kepolisian", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	kepolisians, _ := pagination.Data.([]kepolisian.Kepolisian)
	pagination.Data = kepolisian.FormatKepolisians(kepolisians)

	response := helper.APIResponse("Success get kepolisian", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *kepolisianHandler) GetKepolisian(c *gin.Context) {
	var inputID kepolisian.GetKepolisianDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get kepolisian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKepolisian, err := h.kepolisianService.GetKepolisianByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get kepolisian", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := kepolisian.FormatKepolisian(newKepolisian)
	response := helper.APIResponse("Success Get kepolisian", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
