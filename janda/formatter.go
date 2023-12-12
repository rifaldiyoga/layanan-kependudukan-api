package janda

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type JandaFormatter struct {
	Janda
	CreatedAt string                     `json:"created_at"`
	Penduduk  penduduk.PendudukFormatter `json:"penduduk"`
}

func FormatJanda(janda Janda) JandaFormatter {
	formatter := JandaFormatter{
		Janda:     janda,
		CreatedAt: helper.FormatDateToString(janda.CreatedAt),
		Penduduk:  penduduk.FormatPenduduk(janda.Penduduk),
	}

	return formatter
}

func FormatJandas(sktms []Janda) []JandaFormatter {
	var sktmsFormatter []JandaFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatJanda(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
