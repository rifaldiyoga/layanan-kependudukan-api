package kepolisian

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type KepolisianFormatter struct {
	Kepolisian
	CreatedAt string                     `json:"created_at"`
	Penduduk  penduduk.PendudukFormatter `json:"penduduk"`
}

func FormatKepolisian(kepolisian Kepolisian) KepolisianFormatter {
	formatter := KepolisianFormatter{
		Kepolisian: kepolisian,
		CreatedAt:  helper.FormatDateToString(kepolisian.CreatedAt),
		Penduduk:   penduduk.FormatPenduduk(kepolisian.Penduduk),
	}

	return formatter
}

func FormatKepolisians(sktms []Kepolisian) []KepolisianFormatter {
	var sktmsFormatter []KepolisianFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatKepolisian(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
