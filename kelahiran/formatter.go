package kelahiran

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type KelahiranFormatter struct {
	Kelahiran
	CreatedAt string `json:"created_at"`
	BirthDate string `json:"birth_date"`
	// SubDistrict subdistrict.SubDistrictFormatter `json:"kecamatan"`
	// District    district.DistrictFormatter       `json:"kota"`
	// Provinsi    province.ProvinceFormatter       `json:"provinsi"`
	Penduduk penduduk.PendudukFormatter `json:"penduduk"`
	Ayah     penduduk.PendudukFormatter `json:"ayah"`
	Ibu      penduduk.PendudukFormatter `json:"ibu"`
}

func FormatKelahiran(kelahiran Kelahiran) KelahiranFormatter {
	formatter := KelahiranFormatter{
		Kelahiran: kelahiran,
		CreatedAt: helper.FormatDateToString(kelahiran.CreatedAt),
		BirthDate: helper.FormatDateToString(kelahiran.BirthDate),
		// SubDistrict: subdistrict.FormatSubDistrict(kelahiran.Kecamatan),
		// District:    district.FormatDistrict(kelahiran.Kota),
		// Provinsi:    province.FormatProvince(kelahiran.Provinsi),
		Penduduk: penduduk.FormatPenduduk(kelahiran.Penduduk),
		Ayah:     penduduk.FormatPenduduk(kelahiran.Ayah),
		Ibu:      penduduk.FormatPenduduk(kelahiran.Ibu),
	}

	return formatter
}

func FormatKelahirans(sktms []Kelahiran) []KelahiranFormatter {
	var sktmsFormatter []KelahiranFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatKelahiran(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
