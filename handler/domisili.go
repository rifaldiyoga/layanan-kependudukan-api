package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/domisili"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type domisiliHandler struct {
	domisiliService  domisili.Service
	layananService   layanan.Service
	pengajuanHandler pengajuanHandler
	authService      auth.Service
}

func NewDomisiliHandler(domisiliService domisili.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *domisiliHandler {
	return &domisiliHandler{domisiliService, layananService, pengajuanHandler, authService}
}

func (h *domisiliHandler) CreateDomisili(c *gin.Context) {
	var input domisili.CreateDomisiliInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create domisili", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, _ := c.FormFile("lampiran")

	if file != nil {
		path := helper.FormatFileName(file.Filename)
		filePath := filepath.Join("documents/domisili", path)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.LampiranPath = filePath
	}

	currentLayanan, err := h.layananService.GetLayananByCode("SKD")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create domisili", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.NIK = currentUser.Nik
	newDomisili, err := h.domisiliService.CreateDomisili(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create domisili", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newDomisili.ID
	inputPengajuan.Keterangan = newDomisili.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := domisili.FormatDomisili(newDomisili)
	response := helper.APIResponse("Success create domisili", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *domisiliHandler) UpdateDomisili(c *gin.Context) {
	var inputID domisili.GetDomisiliDetailInput
	var inputData domisili.CreateDomisiliInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update domisili", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update domisili", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDomisili, err := h.domisiliService.UpdateDomisili(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update domisili", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := domisili.FormatDomisili(newDomisili)
	response := helper.APIResponse("Success Update domisili", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *domisiliHandler) DeleteDomisili(c *gin.Context) {
	var inputID domisili.GetDomisiliDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete domisili", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.domisiliService.DeleteDomisili(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete domisili", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete domisili", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *domisiliHandler) GetDomisilis(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.domisiliService.GetDomisilis(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get domisili", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	domisilis, _ := pagination.Data.([]domisili.Domisili)
	pagination.Data = domisili.FormatDomisilis(domisilis)

	response := helper.APIResponse("Success get domisili", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *domisiliHandler) GetDomisili(c *gin.Context) {
	var inputID domisili.GetDomisiliDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get domisili", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newDomisili, err := h.domisiliService.GetDomisiliByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get domisili", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := domisili.FormatDomisili(newDomisili)
	response := helper.APIResponse("Success Get domisili", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
