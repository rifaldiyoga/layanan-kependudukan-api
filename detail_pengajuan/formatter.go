package detail_pengajuan

import "time"

type DetailPengajuanFormatter struct {
	ID          int       `json:"id"`
	PengajuanID int       `json:"pengajuan_id"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

func FormatDetailPengajuan(Pengajuan DetailPengajuan) DetailPengajuanFormatter {
	formatter := DetailPengajuanFormatter{
		ID:          Pengajuan.ID,
		Status:      Pengajuan.Status,
		PengajuanID: Pengajuan.PengajuanID,
		Name:        Pengajuan.Name,
		CreatedBy:   Pengajuan.CreatedBy,
		CreatedAt:   Pengajuan.CreatedAt,
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
