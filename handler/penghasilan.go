package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/penghasilan"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type penghasilanHandler struct {
	penghasilanService penghasilan.Service
	layananService     layanan.Service
	pengajuanHandler   pengajuanHandler
	authService        auth.Service
}

func NewPenghasilanHandler(penghasilanService penghasilan.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *penghasilanHandler {
	return &penghasilanHandler{penghasilanService, layananService, pengajuanHandler, authService}
}

func (h *penghasilanHandler) CreatePenghasilan(c *gin.Context) {
	var input penghasilan.CreatePenghasilanInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create penghasilan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentLayanan, err := h.layananService.GetLayananByCode("SKPOT")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create penghasilan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, _ := c.FormFile("lampiran")

	if file != nil {
		path := helper.FormatFileName(file.Filename)
		filePath := filepath.Join("documents/penghasilan", path)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.Lampiran = filePath
	}

	input.NIK = currentUser.Nik
	newPenghasilan, err := h.penghasilanService.CreatePenghasilan(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create penghasilan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newPenghasilan.ID
	inputPengajuan.Keterangan = newPenghasilan.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := penghasilan.FormatPenghasilan(newPenghasilan)
	response := helper.APIResponse("Success create penghasilan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *penghasilanHandler) UpdatePenghasilan(c *gin.Context) {
	var inputID penghasilan.GetPenghasilanDetailInput
	var inputData penghasilan.CreatePenghasilanInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update penghasilan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update penghasilan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPenghasilan, err := h.penghasilanService.UpdatePenghasilan(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update penghasilan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := penghasilan.FormatPenghasilan(newPenghasilan)
	response := helper.APIResponse("Success Update penghasilan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *penghasilanHandler) DeletePenghasilan(c *gin.Context) {
	var inputID penghasilan.GetPenghasilanDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete penghasilan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.penghasilanService.DeletePenghasilan(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete penghasilan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete penghasilan", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *penghasilanHandler) GetPenghasilans(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.penghasilanService.GetPenghasilans(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get penghasilan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	penghasilans, _ := pagination.Data.([]penghasilan.Penghasilan)
	pagination.Data = penghasilan.FormatPenghasilans(penghasilans)

	response := helper.APIResponse("Success get penghasilan", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *penghasilanHandler) GetPenghasilan(c *gin.Context) {
	var inputID penghasilan.GetPenghasilanDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get penghasilan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPenghasilan, err := h.penghasilanService.GetPenghasilanByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get penghasilan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := penghasilan.FormatPenghasilan(newPenghasilan)
	response := helper.APIResponse("Success Get penghasilan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
