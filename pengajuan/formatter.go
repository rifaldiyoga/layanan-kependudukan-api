package pengajuan

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
	"layanan-kependudukan-api/pengajuan_detail"
)

type PengajuanFormatter struct {
	Pengajuan
	Detail    []pengajuan_detail.DetailPengajuanFormatter `json:"detail"`
	Penduduk  penduduk.PendudukFormatter                  `json:"penduduk"`
	CreatedAt string                                      `json:"created_at"`
	UpdatedAt string                                      `json:"updated_at"`
}

func FormatPengajuan(Pengajuan Pengajuan) PengajuanFormatter {
	formatter := PengajuanFormatter{
		Pengajuan: Pengajuan,
		Penduduk:  penduduk.FormatPenduduk(Pengajuan.Penduduk),
		Detail:    pengajuan_detail.FormatDetailPengajuans(Pengajuan.Detail),
		CreatedAt: helper.FormatDateToString(Pengajuan.CreatedAt),
		UpdatedAt: helper.FormatDateToString(Pengajuan.UpdatedAt),
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
