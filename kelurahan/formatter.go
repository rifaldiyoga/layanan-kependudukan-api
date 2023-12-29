package kelurahan

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/subdistrict"
)

type KelurahanFormatter struct {
	ID        int                              `json:"id"`
	Code      string                           `json:"code"`
	Name      string                           `json:"name"`
	CreatedAt string                           `json:"created_at"`
	Kecamatan subdistrict.SubDistrictFormatter `json:"kecamatan" `
}

func FormatKelurahan(kelurahan Kelurahan) KelurahanFormatter {
	formatter := KelurahanFormatter{
		ID:        kelurahan.ID,
		Code:      kelurahan.Code,
		Name:      helper.CapitalizeEachWord(kelurahan.Name),
		CreatedAt: helper.FormatDateToString(kelurahan.CreatedAt),
		Kecamatan: subdistrict.FormatSubDistrict(kelurahan.Kecamatan),
	}

	return formatter
}

func FormatKelurahans(kelurahans []Kelurahan) []KelurahanFormatter {
	var kelurahansFormatter []KelurahanFormatter

	for _, kelurahan := range kelurahans {
		kelurahanFormatter := FormatKelurahan(kelurahan)
		kelurahansFormatter = append(kelurahansFormatter, kelurahanFormatter)
	}

	return kelurahansFormatter
}
