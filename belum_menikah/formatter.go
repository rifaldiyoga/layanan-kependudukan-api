package belum_menikah

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type BelumMenikahFormatter struct {
	BelumMenikah
	CreatedAt string                     `json:"created_at"`
	Penduduk  penduduk.PendudukFormatter `json:"penduduk"`
}

func FormatBelumMenikah(janda BelumMenikah) BelumMenikahFormatter {
	formatter := BelumMenikahFormatter{
		BelumMenikah: janda,
		Penduduk:     penduduk.FormatPenduduk(janda.Penduduk),
		CreatedAt:    helper.FormatDateToString(janda.CreatedAt),
	}

	return formatter
}

func FormatBelumMenikahs(sktms []BelumMenikah) []BelumMenikahFormatter {
	var sktmsFormatter []BelumMenikahFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatBelumMenikah(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
