package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/pindah"
	pindahDetail "layanan-kependudukan-api/pindah_detail"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type pindahHandler struct {
	pindahService       pindah.Service
	layananService      layanan.Service
	pengajuanHandler    pengajuanHandler
	pindahDetailSerivce pindahDetail.Service
	authService         auth.Service
}

func NewPindahHandler(pindahService pindah.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, pindahDetailSerivce pindahDetail.Service, authService auth.Service) *pindahHandler {
	return &pindahHandler{pindahService, layananService, pengajuanHandler, pindahDetailSerivce, authService}
}

func (h *pindahHandler) CreatePindah(c *gin.Context) {
	var input pindah.CreatePindahInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create pindah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	code := "SKPK"
	if input.Type == "Pindah Datang" {
		code = "SKPD"
	}

	currentLayanan, err := h.layananService.GetLayananByCode(code)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create pindah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, _ := c.FormFile("lampiran")

	if file != nil {
		path := helper.FormatFileName(file.Filename)
		filePath := filepath.Join("documents/pindah", path)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.Lampiran = filePath
	}

	input.NIK = currentUser.Nik
	newPindah, err := h.pindahService.CreatePindah(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create pindah", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// for _, element := range input.Penduduk {
	// 	inputDetai := pindahDetail.CreatePindahDetailInput{
	// 		NIK:          element.NIK,
	// 		Nama:         element.Fullname,
	// 		PindahID:     newPindah.ID,
	// 		StatusFamily: element.StatusFamily,
	// 	}
	// 	h.pindahDetailSerivce.CreatePindahDetail(inputDetai, currentUser)
	// }

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newPindah.ID
	inputPengajuan.Keterangan = newPindah.AlasanPindah

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := pindah.FormatPindah(newPindah)
	response := helper.APIResponse("Success create pindah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pindahHandler) UpdatePindah(c *gin.Context) {
	var inputID pindah.GetPindahDetailInput
	var inputData pindah.CreatePindahInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update pindah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update pindah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPindah, err := h.pindahService.UpdatePindah(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update pindah", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := pindah.FormatPindah(newPindah)
	response := helper.APIResponse("Success Update pindah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pindahHandler) DeletePindah(c *gin.Context) {
	var inputID pindah.GetPindahDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete pindah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.pindahService.DeletePindah(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete pindah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete pindah", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *pindahHandler) GetPindahs(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.pindahService.GetPindahs(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get pindah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pindahs, _ := pagination.Data.([]pindah.Pindah)
	pagination.Data = pindah.FormatPindahs(pindahs)

	response := helper.APIResponse("Success get pindah", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *pindahHandler) GetPindah(c *gin.Context) {
	var inputID pindah.GetPindahDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get pindah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPindah, err := h.pindahService.GetPindahByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get pindah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := pindah.FormatPindah(newPindah)
	response := helper.APIResponse("Success Get pindah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
