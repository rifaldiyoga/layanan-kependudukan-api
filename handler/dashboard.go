package handler

import (
	"layanan-kependudukan-api/auth"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/keluarga"
	"layanan-kependudukan-api/penduduk"
	"layanan-kependudukan-api/pengajuan"
	"net/http"

	"github.com/gin-gonic/gin"
)

type dashboardHandler struct {
	pengajuanService pengajuan.Service
	pendudukService  penduduk.Service
	keluargaService  keluarga.Service
	authService      auth.Service
}

func NewDashboardHandler(pengajuanService pengajuan.Service, pendudukService penduduk.Service, keluargaService keluarga.Service, authService auth.Service) *dashboardHandler {
	return &dashboardHandler{pengajuanService, pendudukService, keluargaService, authService}
}

type Dashboard struct {
	Pengajuan int64 `json:"pengajuan"`
	Penduduk  int64 `json:"penduduk"`
	Keluarga  int64 `json:"keluarga"`
}

func (h *dashboardHandler) GetDashboard(c *gin.Context) {
	countPengajuan, err := h.pengajuanService.GetCountPengajuan()
	if err != nil {
		response := helper.APIResponse("Failed get pengajuan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	countPenduduk, err := h.pendudukService.GetCountPenduduk()
	if err != nil {
		response := helper.APIResponse("Failed get pengajuan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	countKeluarga, err := h.keluargaService.GetCountKeluarga()
	if err != nil {
		response := helper.APIResponse("Failed get pengajuan", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := Dashboard{
		Pengajuan: countPengajuan,
		Penduduk:  countPenduduk,
		Keluarga:  countKeluarga,
	}
	response := helper.APIResponse("Success Get Dashboard", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
