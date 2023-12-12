package kematian

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type KematianFormatter struct {
	Kematian
	CreatedAt string                     `json:"created_at"`
	Penduduk  penduduk.PendudukFormatter `json:"penduduk"`
	Jenazah   penduduk.PendudukFormatter `json:"jenazah"`
}

func FormatKematian(kematian Kematian) KematianFormatter {
	formatter := KematianFormatter{
		Kematian:  kematian,
		CreatedAt: helper.FormatDateToString(kematian.CreatedAt),
		// SubDistrict: subdistrict.FormatSubDistrict(kematian.Kecamatan),
		// District:    district.FormatDistrict(kematian.Kota),
		// Provinsi:    province.FormatProvince(kematian.Provinsi),
		Penduduk: penduduk.FormatPenduduk(kematian.Penduduk),
		Jenazah:  penduduk.FormatPenduduk(kematian.Jenazah),
	}

	return formatter
}

func FormatKematians(sktms []Kematian) []KematianFormatter {
	var sktmsFormatter []KematianFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatKematian(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
