package aparatur_desa

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/position"
)

type AparaturDesaFormatter struct {
	ID        int                        `json:"id"`
	Code      string                     `json:"code"`
	Name      string                     `json:"name"`
	CreatedAt string                     `json:"created_at"`
	Jabatan   position.PositionFormatter `json:"jabatan" `
}

func FormatAparaturDesa(kelurahan AparaturDesa) AparaturDesaFormatter {
	formatter := AparaturDesaFormatter{
		ID:        kelurahan.ID,
		Code:      kelurahan.Code,
		Name:      kelurahan.Name,
		CreatedAt: helper.FormatDateToString(kelurahan.CreatedAt),
		Jabatan:   position.FormatPosition(kelurahan.Jabatan),
	}

	return formatter
}

func FormatAparaturDesas(kelurahans []AparaturDesa) []AparaturDesaFormatter {
	var kelurahansFormatter []AparaturDesaFormatter

	for _, kelurahan := range kelurahans {
		kelurahanFormatter := FormatAparaturDesa(kelurahan)
		kelurahansFormatter = append(kelurahansFormatter, kelurahanFormatter)
	}

	return kelurahansFormatter
}
