package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/penduduk"
	"layanan-kependudukan-api/pengajuan"
	"layanan-kependudukan-api/pengajuan_detail"
	"layanan-kependudukan-api/user"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
)

type pengajuanHandler struct {
	app                    *firebase.App
	pengajuanService       pengajuan.Service
	detailPengajuanService pengajuan_detail.Service
	layananService         layanan.Service
	userService            user.Service
	pendudukService        penduduk.Service
	authService            auth.Service
}

func NewPengajuanHandler(
	app *firebase.App, pengajuanService pengajuan.Service, detailPengajuanService pengajuan_detail.Service, layananService layanan.Service, userService user.Service, pendudukService penduduk.Service, authService auth.Service) *pengajuanHandler {
	return &pengajuanHandler{app, pengajuanService, detailPengajuanService, layananService,

		userService, pendudukService, authService}
}

func (h *pengajuanHandler) CreatePengajuan(c *gin.Context, input pengajuan.CreatePengajuanInput, currentLayanan layanan.Layanan) {

	currentUser, _ := c.Get("currentUser")
	userObject := currentUser.(user.User)

	status := "PENDING"
	if input.Status == "" || input.Status == "PENDING" {
		if currentLayanan.IsConfirm {
			status = "PENDING_RT"
		} else {
			status = "PENDING_ADMIN"
		}
	} else {
		status = input.Status
	}

	currentLayanan, _ = h.layananService.GetLayananByID(input.LayananID)

	input.Code = currentLayanan.Code
	input.Status = status
	newPengajuan, err := h.pengajuanService.CreatePengajuan(input, userObject)
	if err != nil {
		// response := helper.APIResponse("Failed create pengajuan", http.StatusBadRequest, "error", nil)
		// c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.detailPengajuanService.CreateDetailPengajuan(newPengajuan.ID, status, userObject)
	if err != nil {
		// response := helper.APIResponse("Failed create pengajuan", http.StatusBadRequest, "error", nil)
		// c.JSON(http.StatusBadRequest, response)
		return
	}

	newPengajuan, err = h.pengajuanService.GetPengajuanByID(newPengajuan.ID)
	if err != nil {
		// response := helper.APIResponse("Failed get pengajuan", http.StatusBadRequest, "error", nil)
		// c.JSON(http.StatusBadRequest, response)
		return
	}

	h.SendNotification(c, newPengajuan, userObject)

	// formatter := pengajuan.FormatPengajuan(newPengajuan)
	// response := helper.APIResponse("Success create pengajuan", http.StatusOK, "success", formatter)
	// c.JSON(http.StatusOK, response)
}

func (h *pengajuanHandler) UpdatePengajuan(c *gin.Context) {
	var inputID pengajuan.GetPengajuanDetailInput
	var inputData pengajuan.CreatePengajuanInput

	currentUser, _ := c.Get("currentUser")
	userObject := currentUser.(user.User)

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update pengajuan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Update pengajuan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// currentLayanan, err := h.layananService.GetLayananByID(inputData.LayananID)
	// if err != nil {
	// 	errors := helper.FormatValidationError(err)
	// 	errorMessage := gin.H{"errors": errors}

	// 	response := helper.APIResponse("Failed create sktm", http.StatusUnprocessableEntity, "error", errorMessage)
	// 	c.JSON(http.StatusUnprocessableEntity, response)
	// 	return
	// }

	newPengajuan, err := h.pengajuanService.UpdatePengajuan(inputID, inputData, userObject)
	if err != nil {
		response := helper.APIResponse("Failed Update pengajuan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.detailPengajuanService.CreateDetailPengajuan(inputID.ID, inputData.Status, userObject)
	if err != nil {
		response := helper.APIResponse("Failed create pengajuan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newPengajuan, err = h.pengajuanService.GetPengajuanByID(newPengajuan.ID)
	if err != nil {
		response := helper.APIResponse("Failed get pengajuan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// if inputData.Status == "VALID" {
	// 	if currentLayanan.Code == "SKTM" {
	// 		var inputSKTM sktm.CreateSKTMInput
	// 		inputSKTM.NIK = userObject.Nik

	// 		_, err := h.sktmService.UpdateStatus(newPengajuan.RefID)
	// 		if err != nil {
	// 			errors := helper.FormatValidationError(err)
	// 			errorMessage := gin.H{"errors": errors}

	// 			response := helper.APIResponse("Failed create pengajuan", http.StatusUnprocessableEntity, "error", errorMessage)
	// 			c.JSON(http.StatusUnprocessableEntity, response)
	// 			return
	// 		}
	// 	}
	// }

	h.SendNotification(c, newPengajuan, userObject)

	formatter := pengajuan.FormatPengajuan(newPengajuan)
	response := helper.APIResponse("Success Update pengajuan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pengajuanHandler) DeletePengajuan(c *gin.Context) {
	var inputID pengajuan.GetPengajuanDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Delete pengajuan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.pengajuanService.DeletePengajuan(inputID)
	if err != nil {
		response := helper.APIResponse("Failed Delete pengajuan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Delete pengajuan", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *pengajuanHandler) GetPengajuanUser(c *gin.Context) {
	currentUser, _ := c.Get("currentUser")
	userObject := currentUser.(user.User)

	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	pagination, err := h.pengajuanService.GetPengajuanUser(pagination, userObject)
	if err != nil {
		response := helper.APIResponse("Failed get pengajuan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pengajuans, _ := pagination.Data.([]pengajuan.Pengajuan)
	pagination.Data = pengajuan.FormatPengajuans(pengajuans)

	response := helper.APIResponse("Success get pengajuan", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *pengajuanHandler) GetPengajuanAdmin(c *gin.Context) {
	var pagination helper.Pagination

	helper.GetPagingValue(c, &pagination)

	params := c.Request.URL.Query()

	pagination, err := h.pengajuanService.GetPengajuan(pagination, params)
	if err != nil {
		response := helper.APIResponse("Failed get pengajuan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	pengajuans, _ := pagination.Data.([]pengajuan.Pengajuan)
	pagination.Data = pengajuan.FormatPengajuans(pengajuans)

	response := helper.APIResponse("Success get pengajuan", http.StatusOK, "success", pagination)
	c.JSON(http.StatusOK, response)
}

func (h *pengajuanHandler) GetPengajuan(c *gin.Context) {
	var inputID pengajuan.GetPengajuanDetailInput

	errUri := c.ShouldBindUri(&inputID)
	if errUri != nil {
		errors := helper.FormatValidationError(errUri)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed Get Pengajuan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPengajuan, err := h.pengajuanService.GetPengajuanByID(inputID.ID)
	if err != nil {
		response := helper.APIResponse("Failed Get Pengajuan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := pengajuan.FormatPengajuan(newPengajuan)
	response := helper.APIResponse("Success Get Pengajuan", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pengajuanHandler) SendNotification(c *gin.Context, pengajuan pengajuan.Pengajuan, currentUser user.User) {

	pengaju, err := h.pendudukService.GetPendudukByNIK(pengajuan.NIK)
	if err != nil {
		response := helper.APIResponse("Failed Get Pengajuan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	pengajuUser, err := h.userService.GetUserByNIK(pengajuan.NIK)
	if err != nil {
		response := helper.APIResponse("Failed Get Pengajuan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	RT, err := h.pendudukService.GetRTByPengaju(pengaju.RtID, pengaju.RwID)
	if err != nil {
		response := helper.APIResponse("Failed Get Pengajuan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	rtUser, err := h.userService.GetUserByNIK(RT.NIK)
	if err != nil {
		response := helper.APIResponse("Failed Get Pengajuan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	RW, err := h.pendudukService.GetRWByPengaju(pengaju.RwID)
	if err != nil {
		response := helper.APIResponse("Failed Get Pengajuan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	rwUser, err := h.userService.GetUserByNIK(RW.NIK)
	if err != nil {
		response := helper.APIResponse("Failed Get Pengajuan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	adminUser, err := h.userService.GetUserByAdmin()
	if err != nil {
		response := helper.APIResponse("Failed Get Pengajuan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if pengajuan.Status == "PENDING_RT" {
		helper.SendNotification(h.app, rtUser.Email, rtUser.Token, "Pengajuan Surat Baru!", "Pengajuan "+pengajuan.Layanan+" dari "+pengajuan.Name)
	}
	if pengajuan.Status == "PENDING_ADMIN" {
		for _, value := range adminUser {
			helper.SendNotification(h.app, value.Email, value.Token, "Pengajuan Surat Baru!", "Pengajuan "+pengajuan.Layanan+" dari "+pengajuan.Name)
		}
	}
	if pengajuan.Status == "REJECTED_RT" || pengajuan.Status == "REJECTED_RW" || pengajuan.Status == "REJECTED" {
		//kirim notif ke user
		name := ""
		if pengajuan.Status == "REJECTED_RT" {
			name = RT.Fullname
		}
		if pengajuan.Status == "REJECTED_RW" {
			name = RW.Fullname
		}
		helper.SendNotification(h.app, pengajuUser.Email, pengajuUser.Token, "Pengajuan Ditolak "+currentUser.Role+"!", "Pengajuan "+pengajuan.Layanan+" telah ditolak oleh "+name)
	}
	if pengajuan.Status == "APPROVED_RT" {
		// kriim notif ke rw dan user
		helper.SendNotification(h.app, rwUser.Email, rwUser.Token, "Pengajuan Surat Baru!", "Pengajuan "+pengajuan.Layanan+" dari "+pengajuan.Name)
		helper.SendNotification(h.app, pengajuUser.Email, pengajuUser.Token, "Pengajuan Disetujui RT!", "Pengajuan "+pengajuan.Layanan+" telah disetujui oleh "+RT.Fullname)
	}
	if pengajuan.Status == "APPROVED_RW" {
		// kiriim notif ke admin dan user
		helper.SendNotification(h.app, pengajuUser.Email, pengajuUser.Token, "Pengajuan Disetujui RW!", "Pengajuan "+pengajuan.Layanan+" dari "+pengajuan.Name)
		for _, value := range adminUser {
			helper.SendNotification(h.app, value.Email, value.Token, "Pengajuan Surat Baru!", "Pengajuan "+pengajuan.Layanan+" dari "+pengajuan.Name)
		}
	}
	if pengajuan.Status == "VALID" {
		// kiriim notif ke admin dan user
		helper.SendNotification(h.app, pengajuUser.Email, pengajuUser.Token, "Pengajuan Disetujui Kelurahan!", "Pengajuan "+pengajuan.Layanan+" dari "+pengajuan.Name)
	}

}
