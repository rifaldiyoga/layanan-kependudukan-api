package sktm

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type SKTMFormatter struct {
	SKTM
	CreatedAt string                     `json:"created_at"`
	Penduduk  penduduk.PendudukFormatter `json:"penduduk"`
}

func FormatSKTM(sktm SKTM) SKTMFormatter {
	formatter := SKTMFormatter{
		SKTM:      sktm,
		CreatedAt: helper.FormatDateToString(sktm.CreatedAt),
		Penduduk:  penduduk.FormatPenduduk(sktm.Penduduk),
	}

	return formatter
}

func FormatSKTMs(sktms []SKTM) []SKTMFormatter {
	var sktmsFormatter []SKTMFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatSKTM(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
