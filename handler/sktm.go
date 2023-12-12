package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/sktm"
	"layanan-kependudukan-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sktmHandler struct {
	sktmService      sktm.Service
	layananService   layanan.Service
	pengajuanHandler pengajuanHandler
	authService      auth.Service
}

func NewSKTMHandler(sktmService sktm.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *sktmHandler {
	return &sktmHandler{sktmService, layananService, pengajuanHandler, authService}
}

func (h *sktmHandler) CreateSKTM(c *gin.Context) {
	var input sktm.CreateSKTMInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create sktm", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// file, err := c.FormFile("lampiran")
	// if err != nil {
	// 	fmt.Print(err.Error())
	// 	errors := helper.FormatValidationError(err)
	// 	errorMessage := gin.H{"errors": errors}

	// 	response := helper.APIResponse("Failed create user", http.StatusUnprocessableEntity, "error", errorMessage)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	// path := helper.FormatFileName(file.Filename)
	// filePath := filepath.Join("documents/sktm", path)
	// if err := c.SaveUploadedFile(file, filePath); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
	// 	return
	// }

	// input.LampiranPath = filePath

	currentLayanan, err := h.layananService.GetLayananByCode("SKTM")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create sktm", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.NIK = currentUser.Nik
	newSKTM, err := h.sktmService.CreateSKTM(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create sktm", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newSKTM.ID
	inputPengajuan.Keterangan = newSKTM.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := sktm.FormatSKTM(newSKTM)
	response := helper.APIResponse("Success create sktm", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *sktmHandler) UpdateSKTM(c *gin.Context) {
	var inputID sktm.GetSKTMDetailInput
	var inputData sktm.CreateSKTMInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update sktm", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update sktm", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSKTM, err := h.sktmService.UpdateSKTM(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update sktm", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := sktm.FormatSKTM(newSKTM)
	response := helper.APIResponse("Success Update sktm", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *sktmHandler) DeleteSKTM(c *gin.Context) {
	var inputID sktm.GetSKTMDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete sktm", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.sktmService.DeleteSKTM(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete sktm", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete sktm", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *sktmHandler) GetSKTMs(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.sktmService.GetSKTMs(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get sktm", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	sktms, _ := pagination.Data.([]sktm.SKTM)
	pagination.Data = sktm.FormatSKTMs(sktms)

	response := helper.APIResponse("Success get sktm", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *sktmHandler) GetSKTM(c *gin.Context) {
	var inputID sktm.GetSKTMDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get sktm", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSKTM, err := h.sktmService.GetSKTMByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get sktm", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := sktm.FormatSKTM(newSKTM)
	response := helper.APIResponse("Success Get sktm", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
