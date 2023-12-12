package tanah

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type TanahFormatter struct {
	Tanah
	CreatedAt string                     `json:"created_at"`
	Penduduk  penduduk.PendudukFormatter `json:"penduduk"`
}

func FormatTanah(tanah Tanah) TanahFormatter {
	formatter := TanahFormatter{
		Tanah:     tanah,
		CreatedAt: helper.FormatDateToString(tanah.CreatedAt),
		Penduduk:  penduduk.FormatPenduduk(tanah.Penduduk),
	}

	return formatter
}

func FormatTanahs(sktms []Tanah) []TanahFormatter {
	var sktmsFormatter []TanahFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatTanah(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
