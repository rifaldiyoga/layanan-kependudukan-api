package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/janda"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type jandaHandler struct {
	jandaService     janda.Service
	layananService   layanan.Service
	pengajuanHandler pengajuanHandler
	authService      auth.Service
}

func NewJandaHandler(jandaService janda.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *jandaHandler {
	return &jandaHandler{jandaService, layananService, pengajuanHandler, authService}
}

func (h *jandaHandler) CreateJanda(c *gin.Context) {
	var input janda.CreateJandaInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create janda", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentLayanan, err := h.layananService.GetLayananByCode("SKJD")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create janda", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, _ := c.FormFile("lampiran")

	if file != nil {
		path := helper.FormatFileName(file.Filename)
		filePath := filepath.Join("documents/janda", path)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.Lampiran = filePath
	}

	input.NIK = currentUser.Nik
	newJanda, err := h.jandaService.CreateJanda(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create janda", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newJanda.ID
	inputPengajuan.Keterangan = newJanda.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := janda.FormatJanda(newJanda)
	response := helper.APIResponse("Success create janda", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *jandaHandler) UpdateJanda(c *gin.Context) {
	var inputID janda.GetJandaDetailInput
	var inputData janda.CreateJandaInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update janda", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update janda", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newJanda, err := h.jandaService.UpdateJanda(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update janda", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := janda.FormatJanda(newJanda)
	response := helper.APIResponse("Success Update janda", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *jandaHandler) DeleteJanda(c *gin.Context) {
	var inputID janda.GetJandaDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete janda", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.jandaService.DeleteJanda(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete janda", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete janda", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *jandaHandler) GetJandas(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.jandaService.GetJandas(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get janda", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	jandas, _ := pagination.Data.([]janda.Janda)
	pagination.Data = janda.FormatJandas(jandas)

	response := helper.APIResponse("Success get janda", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *jandaHandler) GetJanda(c *gin.Context) {
	var inputID janda.GetJandaDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get janda", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newJanda, err := h.jandaService.GetJandaByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get janda", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := janda.FormatJanda(newJanda)
	response := helper.APIResponse("Success Get janda", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
