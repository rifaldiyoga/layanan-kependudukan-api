package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/penduduk"
	"layanan-kependudukan-api/pengajuan"
	penahMenikah "layanan-kependudukan-api/pernah_menikah"
	"layanan-kependudukan-api/user"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type penahMenikahHandler struct {
	penahMenikahService penahMenikah.Service
	pendudukService     penduduk.Service
	layananService      layanan.Service
	pengajuanHandler    pengajuanHandler
	authService         auth.Service
}

func NewPernahMenikahHandler(penahMenikahService penahMenikah.Service, pendudukService penduduk.Service, layananService layanan.Service, pengajuanHandler pengajuanHandler, authService auth.Service) *penahMenikahHandler {
	return &penahMenikahHandler{penahMenikahService, pendudukService, layananService, pengajuanHandler, authService}
}

func (h *penahMenikahHandler) CreatePernahMenikah(c *gin.Context) {
	var input penahMenikah.CreatePernahMenikahInput

	userObject, _ := c.Get("currentUser")
	currentUser := userObject.(user.User)

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create penahMenikah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentLayanan, err := h.layananService.GetLayananByCode("SKPN")
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create penahMenikah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, _ := c.FormFile("lampiran")

	if file != nil {
		path := helper.FormatFileName(file.Filename)
		filePath := filepath.Join("documents/pernah_menikah", path)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		input.Lampiran = filePath
	}

	input.NIK = currentUser.Nik
	currentPenduduk, err := h.pendudukService.GetPendudukByNIK(input.NIK)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create penahMenikah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	currentKK, err := h.pendudukService.GetPendudukByNoKK(currentPenduduk.NoKK)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create penahMenikah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if currentPenduduk.JK == "L" {
		input.NIKSuami = currentPenduduk.NIK
		for _, value := range currentKK {
			if value.StatusFamily == "Istri" {
				input.NIKIstri = value.NIK
			}
		}
	} else {
		input.NIKIstri = currentPenduduk.NIK
		for _, value := range currentKK {
			if value.StatusFamily == "Suami" || value.StatusFamily == "Kepala Keluarga" {
				input.NIKIstri = value.NIK
			}
		}
	}

	newPernahMenikah, err := h.penahMenikahService.CreatePernahMenikah(input, currentLayanan, currentUser)
	if err != nil {
		response := helper.APIResponse("Failed create penahMenikah", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inputPengajuan := pengajuan.CreatePengajuanInput{}
	inputPengajuan.Layanan = currentLayanan.Name
	inputPengajuan.LayananID = currentLayanan.ID
	inputPengajuan.RefID = newPernahMenikah.ID
	inputPengajuan.Keterangan = newPernahMenikah.Keterangan

	h.pengajuanHandler.CreatePengajuan(c, inputPengajuan, currentLayanan)

	formatter := penahMenikah.FormatPernahMenikah(newPernahMenikah)
	response := helper.APIResponse("Success create penahMenikah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *penahMenikahHandler) UpdatePernahMenikah(c *gin.Context) {
	var inputID penahMenikah.GetPernahMenikahDetailInput
	var inputData penahMenikah.CreatePernahMenikahInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update penahMenikah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update penahMenikah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPernahMenikah, err := h.penahMenikahService.UpdatePernahMenikah(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed Update penahMenikah", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := penahMenikah.FormatPernahMenikah(newPernahMenikah)
	response := helper.APIResponse("Success Update penahMenikah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *penahMenikahHandler) DeletePernahMenikah(c *gin.Context) {
	var inputID penahMenikah.GetPernahMenikahDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete penahMenikah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.penahMenikahService.DeletePernahMenikah(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete penahMenikah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete penahMenikah", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *penahMenikahHandler) GetPernahMenikahs(c *gin.Context) {

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()
	pagination, err := h.penahMenikahService.GetPernahMenikahs(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get penahMenikah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	penahMenikahs, _ := pagination.Data.([]penahMenikah.PernahMenikah)
	pagination.Data = penahMenikah.FormatPernahMenikahs(penahMenikahs)

	response := helper.APIResponse("Success get penahMenikah", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *penahMenikahHandler) GetPernahMenikah(c *gin.Context) {
	var inputID penahMenikah.GetPernahMenikahDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get penahMenikah", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPernahMenikah, err := h.penahMenikahService.GetPernahMenikahByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get penahMenikah", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := penahMenikah.FormatPernahMenikah(newPernahMenikah)
	response := helper.APIResponse("Success Get penahMenikah", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
