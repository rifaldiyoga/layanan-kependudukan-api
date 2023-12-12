package pengajuan_detail

import (
	"layanan-kependudukan-api/helper"
)

type DetailPengajuanFormatter struct {
	ID          int    `json:"id"`
	PengajuanID int    `json:"pengajuan_id"`
	Status      string `json:"status"`
	Name        string `json:"name"`
	Note        string `json:"note"`
	CreatedBy   int    `json:"created_by"`
	CreatedAt   string `json:"created_at"`
}

func FormatDetailPengajuan(Pengajuan DetailPengajuan) DetailPengajuanFormatter {
	formatter := DetailPengajuanFormatter{
		ID:          Pengajuan.ID,
		Status:      Pengajuan.Status,
		PengajuanID: Pengajuan.PengajuanID,
		Name:        Pengajuan.Name,
		Note:        getNote(Pengajuan.Status),
		CreatedBy:   Pengajuan.CreatedBy,
		CreatedAt:   helper.FormatDateToString(Pengajuan.CreatedAt),
	}

	return formatter
}

func FormatDetailPengajuans(Pengajuans []DetailPengajuan) []DetailPengajuanFormatter {
	var PengajuansFormatter []DetailPengajuanFormatter

	for _, Pengajuan := range Pengajuans {
		DetailPengajuanFormatter := FormatDetailPengajuan(Pengajuan)
		PengajuansFormatter = append(PengajuansFormatter, DetailPengajuanFormatter)
	}

	return PengajuansFormatter
}

func getNote(status string) string {
	if status == "PENDING_RT" {
		return "Pengajuan Dibuat. Menunggu persetujuan RT"
	}
	if status == "APPROVED_RT" {
		return "Pengajuan telah disetujui RT. Menunggu persetujuan RW"
	}
	if status == "REJECTED_RT" {
		return "Pengajuan telah ditolak oleh RT"
	}
	if status == "APPROVED_RW" {
		return "Pengajuan telah disetujui RW. Menunggu persetujuan Kelurahan"
	}
	if status == "REJECTED_RW" {
		return "Pengajuan telah ditolak oleh RW"
	}
	if status == "PENDING_ADMIN" {
		return "Pengajuan Dibuat. Menunggu persetujuan Kelurahan"
	}

	if status == "VALID" {
		return "Pengajuan telah disetujui oleh Kelurahan"
	}
	if status == "REJECTED" {
		return "Pengajuan telah ditolak oleh Kelurahan"
	}
	return ""
}
