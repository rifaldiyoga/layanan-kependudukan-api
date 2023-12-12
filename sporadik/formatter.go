package sporadik

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type SporadikFormatter struct {
	Sporadik
	CreatedAt string                     `json:"created_at"`
	Penduduk  penduduk.PendudukFormatter `json:"penduduk"`
}

func FormatSporadik(sporadik Sporadik) SporadikFormatter {
	formatter := SporadikFormatter{
		Sporadik:  sporadik,
		CreatedAt: helper.FormatDateToString(sporadik.CreatedAt),
		Penduduk:  penduduk.FormatPenduduk(sporadik.Penduduk),
	}

	return formatter
}

func FormatSporadiks(sktms []Sporadik) []SporadikFormatter {
	var sktmsFormatter []SporadikFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatSporadik(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
