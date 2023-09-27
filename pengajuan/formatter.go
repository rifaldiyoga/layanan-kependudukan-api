package pengajuan

import (
	"layanan-kependudukan-api/helper"
)

type PengajuanFormatter struct {
	Pengajuan
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FormatPengajuan(Pengajuan Pengajuan) PengajuanFormatter {
	formatter := PengajuanFormatter{
		Pengajuan: Pengajuan,
		CreatedAt: helper.FormatDateToString(Pengajuan.CreatedAt),
		UpdatedAt: helper.FormatDateToString(Pengajuan.UpdatedAt),
		Status:    Pengajuan.Detail[len(Pengajuan.Detail)-1].Status,
	}

	return formatter
}

func FormatPengajuans(Pengajuans []Pengajuan) []PengajuanFormatter {
	var PengajuansFormatter []PengajuanFormatter

	for _, Pengajuan := range Pengajuans {
		PengajuanFormatter := FormatPengajuan(Pengajuan)
		PengajuansFormatter = append(PengajuansFormatter, PengajuanFormatter)
	}

	return PengajuansFormatter
}