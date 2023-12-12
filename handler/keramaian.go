package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/keramaian"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type keramaianHandler struct {
	keramaianService keramaian.Service
	layananService   layanan.Service
	pengajuanHandler pengajuanHandler
	authService      auth.Service
}

func NewKeramaianHandler(keramaianService keramaian.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *keramaianHandler {
	return &keramaianHandler{keramaianService, layananService, pengajuanHandler, authService}
}

func (h *keramaianHandler) CreateKeramaian(c *gin.Context) {
	var input keramaian.CreateKeramaianInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create keramaian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentLayanan, err := h.layananService.GetLayananByCode("SIK")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create keramaian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, _ := c.FormFile("lampiran")

	if file != nil {
		path := helper.FormatFileName(file.Filename)
		filePath := filepath.Join("documents/keramaian", path)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.Lampiran = filePath
	}

	input.NIK = currentUser.Nik
	newKeramaian, err := h.keramaianService.CreateKeramaian(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create keramaian", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newKeramaian.ID
	inputPengajuan.Keterangan = newKeramaian.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := keramaian.FormatKeramaian(newKeramaian)
	response := helper.APIResponse("Success create keramaian", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *keramaianHandler) UpdateKeramaian(c *gin.Context) {
	var inputID keramaian.GetKeramaianDetailInput
	var inputData keramaian.CreateKeramaianInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update keramaian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update keramaian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKeramaian, err := h.keramaianService.UpdateKeramaian(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update keramaian", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := keramaian.FormatKeramaian(newKeramaian)
	response := helper.APIResponse("Success Update keramaian", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *keramaianHandler) DeleteKeramaian(c *gin.Context) {
	var inputID keramaian.GetKeramaianDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete keramaian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.keramaianService.DeleteKeramaian(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete keramaian", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete keramaian", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *keramaianHandler) GetKeramaians(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.keramaianService.GetKeramaians(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get keramaian", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	keramaians, _ := pagination.Data.([]keramaian.Keramaian)
	pagination.Data = keramaian.FormatKeramaians(keramaians)

	response := helper.APIResponse("Success get keramaian", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *keramaianHandler) GetKeramaian(c *gin.Context) {
	var inputID keramaian.GetKeramaianDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get keramaian", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKeramaian, err := h.keramaianService.GetKeramaianByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get keramaian", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := keramaian.FormatKeramaian(newKeramaian)
	response := helper.APIResponse("Success Get keramaian", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
