package pernah_menikah

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type PernahMenikahFormatter struct {
	PernahMenikah
	CreatedAt string                     `json:"created_at"`
	Penduduk  penduduk.PendudukFormatter `json:"penduduk"`
	Suami     penduduk.PendudukFormatter `json:"suami"`
	Istri     penduduk.PendudukFormatter `json:"istri"`
}

func FormatPernahMenikah(janda PernahMenikah) PernahMenikahFormatter {
	formatter := PernahMenikahFormatter{
		PernahMenikah: janda,
		Penduduk:      penduduk.FormatPenduduk(janda.Penduduk),
		Suami:         penduduk.FormatPenduduk(janda.Suami),
		Istri:         penduduk.FormatPenduduk(janda.Istri),
		CreatedAt:     helper.FormatDateToString(janda.CreatedAt),
	}

	return formatter
}

func FormatPernahMenikahs(sktms []PernahMenikah) []PernahMenikahFormatter {
	var sktmsFormatter []PernahMenikahFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatPernahMenikah(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
