package rumah

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type RumahFormatter struct {
	Rumah
	CreatedAt string                     `json:"created_at"`
	Penduduk  penduduk.PendudukFormatter `json:"penduduk"`
}

func FormatRumah(rumah Rumah) RumahFormatter {
	formatter := RumahFormatter{
		Rumah:     rumah,
		CreatedAt: helper.FormatDateToString(rumah.CreatedAt),
		Penduduk:  penduduk.FormatPenduduk(rumah.Penduduk),
	}

	return formatter
}

func FormatRumahs(sktms []Rumah) []RumahFormatter {
	var sktmsFormatter []RumahFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatRumah(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
