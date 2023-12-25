package sistem

import (
	"layanan-kependudukan-api/district"
	"layanan-kependudukan-api/province"
	"layanan-kependudukan-api/subdistrict"
)

type SistemFormatter struct {
	Sistem
	SubDistrict subdistrict.SubDistrictFormatter `json:"kecamatan"`
	District    district.DistrictFormatter       `json:"kota"`
	Provinsi    province.ProvinceFormatter       `json:"provinsi"`
}

func FormatSistem(sistem Sistem) SistemFormatter {
	formatter := SistemFormatter{
		Sistem:      sistem,
		SubDistrict: subdistrict.FormatSubDistrict(sistem.Kecamatan),
		District:    district.FormatDistrict(sistem.Kota),
		Provinsi:    province.FormatProvince(sistem.Provinsi),
	}

	return formatter
}

func FormatSistems(sistems []Sistem) []SistemFormatter {
	var sistemsFormatter []SistemFormatter

	for _, sistem := range sistems {
		sistemFormatter := FormatSistem(sistem)
		sistemsFormatter = append(sistemsFormatter, sistemFormatter)
	}

	return sistemsFormatter
}
