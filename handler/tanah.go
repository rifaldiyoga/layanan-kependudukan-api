package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/tanah"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type tanahHandler struct {
	tanahService     tanah.Service
	layananService   layanan.Service
	pengajuanHandler pengajuanHandler
	authService      auth.Service
}

func NewTanahHandler(tanahService tanah.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *tanahHandler {
	return &tanahHandler{tanahService, layananService, pengajuanHandler, authService}
}

func (h *tanahHandler) CreateTanah(c *gin.Context) {
	var input tanah.CreateTanahInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create tanah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, _ := c.FormFile("lampiran")

	if file != nil {
		path := helper.FormatFileName(file.Filename)
		filePath := filepath.Join("documents/tanah", path)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.Lampiran = filePath
	}

	code := "SKKT"
	if input.Type == "Sporadik" {
		code = "SPORADIK"
	}

	currentLayanan, err := h.layananService.GetLayananByCode(code)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create tanah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.NIK = currentUser.Nik
	newTanah, err := h.tanahService.CreateTanah(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create tanah", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newTanah.ID
	inputPengajuan.Keterangan = newTanah.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := tanah.FormatTanah(newTanah)
	response := helper.APIResponse("Success create tanah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *tanahHandler) UpdateTanah(c *gin.Context) {
	var inputID tanah.GetTanahDetailInput
	var inputData tanah.CreateTanahInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update tanah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update tanah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newTanah, err := h.tanahService.UpdateTanah(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update tanah", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := tanah.FormatTanah(newTanah)
	response := helper.APIResponse("Success Update tanah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *tanahHandler) DeleteTanah(c *gin.Context) {
	var inputID tanah.GetTanahDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete tanah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.tanahService.DeleteTanah(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete tanah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete tanah", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *tanahHandler) GetTanahs(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.tanahService.GetTanahs(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get tanah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	tanahs, _ := pagination.Data.([]tanah.Tanah)
	pagination.Data = tanah.FormatTanahs(tanahs)

	response := helper.APIResponse("Success get tanah", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *tanahHandler) GetTanah(c *gin.Context) {
	var inputID tanah.GetTanahDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get tanah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newTanah, err := h.tanahService.GetTanahByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get tanah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := tanah.FormatTanah(newTanah)
	response := helper.APIResponse("Success Get tanah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
