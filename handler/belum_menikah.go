package handler

import (
	"layanan-kependudukan-api/auth"
	belumMenikah "layanan-kependudukan-api/belum_menikah"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type belumMenikahHandler struct {
	belumMenikahService belumMenikah.Service
	layananService      layanan.Service
	pengajuanHandler    pengajuanHandler
	authService         auth.Service
}

func NewBelumMenikahHandler(belumMenikahService belumMenikah.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *belumMenikahHandler {
	return &belumMenikahHandler{belumMenikahService, layananService, pengajuanHandler, authService}
}

func (h *belumMenikahHandler) CreateBelumMenikah(c *gin.Context) {
	var input belumMenikah.CreateBelumMenikahInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create belumMenikah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentLayanan, err := h.layananService.GetLayananByCode("SKBPN")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create belumMenikah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, _ := c.FormFile("lampiran")

	if file != nil {
		path := helper.FormatFileName(file.Filename)
		filePath := filepath.Join("documents/belum_menikah", path)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.Lampiran = filePath
	}

	input.NIK = currentUser.Nik
	newBelumMenikah, err := h.belumMenikahService.CreateBelumMenikah(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create belumMenikah", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newBelumMenikah.ID
	inputPengajuan.Keterangan = newBelumMenikah.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := belumMenikah.FormatBelumMenikah(newBelumMenikah)
	response := helper.APIResponse("Success create belumMenikah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *belumMenikahHandler) UpdateBelumMenikah(c *gin.Context) {
	var inputID belumMenikah.GetBelumMenikahDetailInput
	var inputData belumMenikah.CreateBelumMenikahInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update belumMenikah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update belumMenikah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newBelumMenikah, err := h.belumMenikahService.UpdateBelumMenikah(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update belumMenikah", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := belumMenikah.FormatBelumMenikah(newBelumMenikah)
	response := helper.APIResponse("Success Update belumMenikah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *belumMenikahHandler) DeleteBelumMenikah(c *gin.Context) {
	var inputID belumMenikah.GetBelumMenikahDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete belumMenikah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.belumMenikahService.DeleteBelumMenikah(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete belumMenikah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete belumMenikah", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *belumMenikahHandler) GetBelumMenikahs(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()

	pagination, err := h.belumMenikahService.GetBelumMenikahs(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get belumMenikah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	belumMenikahs, _ := pagination.Data.([]belumMenikah.BelumMenikah)
	pagination.Data = belumMenikah.FormatBelumMenikahs(belumMenikahs)

	response := helper.APIResponse("Success get belumMenikah", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *belumMenikahHandler) GetBelumMenikah(c *gin.Context) {
	var inputID belumMenikah.GetBelumMenikahDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get belumMenikah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newBelumMenikah, err := h.belumMenikahService.GetBelumMenikahByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get belumMenikah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := belumMenikah.FormatBelumMenikah(newBelumMenikah)
	response := helper.APIResponse("Success Get belumMenikah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
