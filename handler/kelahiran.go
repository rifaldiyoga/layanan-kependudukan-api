package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/kelahiran"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type kelahiranHandler struct {
	kelahiranService kelahiran.Service
	layananService   layanan.Service
	pengajuanHandler pengajuanHandler
	authService      auth.Service
}

func NewKelahiranHandler(kelahiranService kelahiran.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *kelahiranHandler {
	return &kelahiranHandler{kelahiranService, layananService, pengajuanHandler, authService}
}

func (h *kelahiranHandler) CreateKelahiran(c *gin.Context) {
	var input kelahiran.CreateKelahiranInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create kelahiran", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentLayanan, err := h.layananService.GetLayananByCode("SKKH")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create kelahiran", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, _ := c.FormFile("lampiran_buku_nikah")

	if file != nil {
		path := helper.FormatFileName(file.Filename)
		filePath := filepath.Join("documents/kelahiran", path)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.LampiranBukuNikah = filePath
	}

	fileKetRs, _ := c.FormFile("lampiran_ket_rs")

	if fileKetRs != nil {
		path := helper.FormatFileName(fileKetRs.Filename)
		filePath := filepath.Join("documents/kelahiran", path)
		if err := c.SaveUploadedFile(fileKetRs, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.LampiranKetRs = filePath
	}

	input.NIK = currentUser.Nik
	newKelahiran, err := h.kelahiranService.CreateKelahiran(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create kelahiran", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newKelahiran.ID
	inputPengajuan.Keterangan = newKelahiran.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := kelahiran.FormatKelahiran(newKelahiran)
	response := helper.APIResponse("Success create kelahiran", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *kelahiranHandler) UpdateKelahiran(c *gin.Context) {
	var inputID kelahiran.GetKelahiranDetailInput
	var inputData kelahiran.CreateKelahiranInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update kelahiran", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update kelahiran", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKelahiran, err := h.kelahiranService.UpdateKelahiran(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update kelahiran", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := kelahiran.FormatKelahiran(newKelahiran)
	response := helper.APIResponse("Success Update kelahiran", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *kelahiranHandler) DeleteKelahiran(c *gin.Context) {
	var inputID kelahiran.GetKelahiranDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete kelahiran", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.kelahiranService.DeleteKelahiran(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete kelahiran", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete kelahiran", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *kelahiranHandler) GetKelahirans(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.kelahiranService.GetKelahirans(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get kelahiran", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	kelahirans, _ := pagination.Data.([]kelahiran.Kelahiran)
	pagination.Data = kelahiran.FormatKelahirans(kelahirans)

	response := helper.APIResponse("Success get kelahiran", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *kelahiranHandler) GetKelahiran(c *gin.Context) {
	var inputID kelahiran.GetKelahiranDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get kelahiran", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newKelahiran, err := h.kelahiranService.GetKelahiranByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get kelahiran", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := kelahiran.FormatKelahiran(newKelahiran)
	response := helper.APIResponse("Success Get kelahiran", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
