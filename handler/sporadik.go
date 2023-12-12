package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/sporadik"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type sporadikHandler struct {
	sporadikService  sporadik.Service
	layananService   layanan.Service
	pengajuanHandler pengajuanHandler
	authService      auth.Service
}

func NewSporadikHandler(sporadikService sporadik.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *sporadikHandler {
	return &sporadikHandler{sporadikService, layananService, pengajuanHandler, authService}
}

func (h *sporadikHandler) CreateSporadik(c *gin.Context) {
	var input sporadik.CreateSporadikInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create sporadik", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	filePemohon, _ := c.FormFile("lampiran_pemohon")

	if filePemohon != nil {
		path := helper.FormatFileName(filePemohon.Filename)
		filePath := filepath.Join("documents/sporadik", path)
		if err := c.SaveUploadedFile(filePemohon, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.LampiranPemohon = filePath
	}

	fileSpLama, _ := c.FormFile("lampiran_sporadik_lama")

	if fileSpLama != nil {
		path := helper.FormatFileName(fileSpLama.Filename)
		filePath := filepath.Join("documents/sporadik", path)
		if err := c.SaveUploadedFile(fileSpLama, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.LampiranSporadikLama = filePath
	}

	fileSpBaru, _ := c.FormFile("lampiran_sporadik_baru")

	if fileSpBaru != nil {
		path := helper.FormatFileName(fileSpBaru.Filename)
		filePath := filepath.Join("documents/sporadik", path)
		if err := c.SaveUploadedFile(fileSpBaru, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.LampiranSporadikBaru = filePath
	}

	fileBukti, _ := c.FormFile("lampiran_bukti")

	if fileBukti != nil {
		path := helper.FormatFileName(fileBukti.Filename)
		filePath := filepath.Join("documents/sporadik", path)
		if err := c.SaveUploadedFile(fileBukti, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.LampiranBukti = filePath
	}

	fileLunasPbb, _ := c.FormFile("lampiran_lunas_pbb")

	if fileLunasPbb != nil {
		path := helper.FormatFileName(fileLunasPbb.Filename)
		filePath := filepath.Join("documents/sporadik", path)
		if err := c.SaveUploadedFile(fileLunasPbb, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.LampiranLunasPbb = filePath
	}

	currentLayanan, err := h.layananService.GetLayananByCode("SSP")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create sporadik", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.NIK = currentUser.Nik
	newSporadik, err := h.sporadikService.CreateSporadik(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create sporadik", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newSporadik.ID
	inputPengajuan.Keterangan = newSporadik.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := sporadik.FormatSporadik(newSporadik)
	response := helper.APIResponse("Success create sporadik", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *sporadikHandler) UpdateSporadik(c *gin.Context) {
	var inputID sporadik.GetSporadikDetailInput
	var inputData sporadik.CreateSporadikInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update sporadik", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update sporadik", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSporadik, err := h.sporadikService.UpdateSporadik(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update sporadik", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := sporadik.FormatSporadik(newSporadik)
	response := helper.APIResponse("Success Update sporadik", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *sporadikHandler) DeleteSporadik(c *gin.Context) {
	var inputID sporadik.GetSporadikDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete sporadik", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.sporadikService.DeleteSporadik(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete sporadik", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete sporadik", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *sporadikHandler) GetSporadiks(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.sporadikService.GetSporadiks(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get sporadik", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	sporadiks, _ := pagination.Data.([]sporadik.Sporadik)
	pagination.Data = sporadik.FormatSporadiks(sporadiks)

	response := helper.APIResponse("Success get sporadik", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *sporadikHandler) GetSporadik(c *gin.Context) {
	var inputID sporadik.GetSporadikDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get sporadik", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSporadik, err := h.sporadikService.GetSporadikByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get sporadik", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := sporadik.FormatSporadik(newSporadik)
	response := helper.APIResponse("Success Get sporadik", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
