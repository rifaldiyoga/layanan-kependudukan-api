package penghasilan

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type PenghasilanFormatter struct {
	Penghasilan
	CreatedAt string                     `json:"created_at"`
	Penduduk  penduduk.PendudukFormatter `json:"penduduk"`
}

func FormatPenghasilan(penghasilan Penghasilan) PenghasilanFormatter {
	formatter := PenghasilanFormatter{
		Penghasilan: penghasilan,
		CreatedAt:   helper.FormatDateToString(penghasilan.CreatedAt),
		Penduduk:    penduduk.FormatPenduduk(penghasilan.Penduduk),
	}

	return formatter
}

func FormatPenghasilans(sktms []Penghasilan) []PenghasilanFormatter {
	var sktmsFormatter []PenghasilanFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatPenghasilan(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
