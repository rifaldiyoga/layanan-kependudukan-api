package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/rumah"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type rumahHandler struct {
	rumahService     rumah.Service
	layananService   layanan.Service
	pengajuanHandler pengajuanHandler
	authService      auth.Service
}

func NewRumahHandler(rumahService rumah.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *rumahHandler {
	return &rumahHandler{rumahService, layananService, pengajuanHandler, authService}
}

func (h *rumahHandler) CreateRumah(c *gin.Context) {
	var input rumah.CreateRumahInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create rumah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentLayanan, err := h.layananService.GetLayananByCode("SKTMR")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create rumah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, _ := c.FormFile("lampiran")

	if file != nil {
		path := helper.FormatFileName(file.Filename)
		filePath := filepath.Join("documents/rumah", path)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.Lampiran = filePath
	}

	input.NIK = currentUser.Nik
	newRumah, err := h.rumahService.CreateRumah(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create rumah", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newRumah.ID
	inputPengajuan.Keterangan = newRumah.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := rumah.FormatRumah(newRumah)
	response := helper.APIResponse("Success create rumah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *rumahHandler) UpdateRumah(c *gin.Context) {
	var inputID rumah.GetRumahDetailInput
	var inputData rumah.CreateRumahInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update rumah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update rumah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newRumah, err := h.rumahService.UpdateRumah(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update rumah", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := rumah.FormatRumah(newRumah)
	response := helper.APIResponse("Success Update rumah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *rumahHandler) DeleteRumah(c *gin.Context) {
	var inputID rumah.GetRumahDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete rumah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.rumahService.DeleteRumah(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete rumah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete rumah", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *rumahHandler) GetRumahs(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.rumahService.GetRumahs(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get rumah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	rumahs, _ := pagination.Data.([]rumah.Rumah)
	pagination.Data = rumah.FormatRumahs(rumahs)

	response := helper.APIResponse("Success get rumah", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *rumahHandler) GetRumah(c *gin.Context) {
	var inputID rumah.GetRumahDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get rumah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newRumah, err := h.rumahService.GetRumahByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get rumah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := rumah.FormatRumah(newRumah)
	response := helper.APIResponse("Success Get rumah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
