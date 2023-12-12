package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/sku"
	"layanan-kependudukan-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type skuHandler struct {
	skuService       sku.Service
	layananService   layanan.Service
	pengajuanHandler pengajuanHandler
	authService      auth.Service
}

func NewSKUHandler(skuService sku.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *skuHandler {
	return &skuHandler{skuService, layananService, pengajuanHandler, authService}
}

func (h *skuHandler) CreateSKU(c *gin.Context) {
	var input sku.CreateSKUInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create sku", http.StatusUnprocessableEntity, "error", errorMessage)
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
	// filePath := filepath.Join("documents/sku", path)
	// if err := c.SaveUploadedFile(file, filePath); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
	// 	return
	// }

	// input.LampiranPath = filePath

	currentLayanan, err := h.layananService.GetLayananByCode("SKU")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create sku", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.NIK = currentUser.Nik
	newSKU, err := h.skuService.CreateSKU(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create sku", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newSKU.ID
	inputPengajuan.Keterangan = newSKU.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := sku.FormatSKU(newSKU)
	response := helper.APIResponse("Success create sku", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *skuHandler) UpdateSKU(c *gin.Context) {
	var inputID sku.GetSKUDetailInput
	var inputData sku.CreateSKUInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update sku", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update sku", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSKU, err := h.skuService.UpdateSKU(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update sku", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := sku.FormatSKU(newSKU)
	response := helper.APIResponse("Success Update sku", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *skuHandler) DeleteSKU(c *gin.Context) {
	var inputID sku.GetSKUDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete sku", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.skuService.DeleteSKU(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete sku", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete sku", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *skuHandler) GetSKUs(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.skuService.GetSKUs(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get sku", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	skus, _ := pagination.Data.([]sku.SKU)
	pagination.Data = sku.FormatSKUs(skus)

	response := helper.APIResponse("Success get sku", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *skuHandler) GetSKU(c *gin.Context) {
	var inputID sku.GetSKUDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get sku", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newSKU, err := h.skuService.GetSKUByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get sku", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := sku.FormatSKU(newSKU)
	response := helper.APIResponse("Success Get sku", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
