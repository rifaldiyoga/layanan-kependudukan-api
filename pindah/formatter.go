package pindah

import (
	"layanan-kependudukan-api/district"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
	"layanan-kependudukan-api/pindah_detail"
	"layanan-kependudukan-api/province"
	"layanan-kependudukan-api/subdistrict"
)

type PindahFormatter struct {
	Pindah
	CreatedAt    string                                `json:"created_at"`
	Penduduk     penduduk.PendudukFormatter            `json:"penduduk"`
	PindahDetail []pindah_detail.PindahDetailFormatter `json:"pindah_detail"`
	SubDistrict  subdistrict.SubDistrictFormatter      `json:"kecamatan"`
	District     district.DistrictFormatter            `json:"kota"`
	Provinsi     province.ProvinceFormatter            `json:"provinsi"`
}

func FormatPindah(berpergian Pindah) PindahFormatter {
	formatter := PindahFormatter{
		Pindah:       berpergian,
		CreatedAt:    helper.FormatDateToString(berpergian.CreatedAt),
		Penduduk:     penduduk.FormatPenduduk(berpergian.Penduduk),
		PindahDetail: pindah_detail.FormatPindahDetails(berpergian.PindahDetail),
		SubDistrict:  subdistrict.FormatSubDistrict(berpergian.Kecamatan),
		District:     district.FormatDistrict(berpergian.Kota),
		Provinsi:     province.FormatProvince(berpergian.Provinsi),
	}

	return formatter
}

func FormatPindahs(sktms []Pindah) []PindahFormatter {
	var sktmsFormatter []PindahFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatPindah(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
