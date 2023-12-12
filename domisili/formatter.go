package domisili

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type DomisiliFormatter struct {
	Domisili
	CreatedAt string                     `json:"created_at"`
	Penduduk  penduduk.PendudukFormatter `json:"penduduk"`
}

func FormatDomisili(domisili Domisili) DomisiliFormatter {
	formatter := DomisiliFormatter{
		Domisili:  domisili,
		CreatedAt: helper.FormatDateToString(domisili.CreatedAt),
		Penduduk:  penduduk.FormatPenduduk(domisili.Penduduk),
	}

	return formatter
}

func FormatDomisilis(sktms []Domisili) []DomisiliFormatter {
	var sktmsFormatter []DomisiliFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatDomisili(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
