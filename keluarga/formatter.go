package keluarga

import (
	"layanan-kependudukan-api/helper"
)

type KeluargaFormatter struct {
	Keluarga
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func FormatKeluarga(Keluarga Keluarga) KeluargaFormatter {
	formatter := KeluargaFormatter{
		Keluarga:  Keluarga,
		CreatedAt: helper.FormatDateToString(Keluarga.CreatedAt),
		UpdatedAt: helper.FormatDateToString(Keluarga.UpdatedAt),
	}

	return formatter
}

func FormatKeluargas(keluargas []Keluarga) []KeluargaFormatter {
	var keluargasFormatter []KeluargaFormatter

	for _, keluarga := range keluargas {
		keluargaFormatter := FormatKeluarga(keluarga)
		keluargasFormatter = append(keluargasFormatter, keluargaFormatter)
	}

	return keluargasFormatter
}
