package keluarga

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type KeluargaFormatter struct {
	Keluarga
	CreatedAt string                       `json:"created_at"`
	UpdatedAt string                       `json:"updated_at"`
	Penduduk  []penduduk.PendudukFormatter `json:"penduduk"`
}

func FormatKeluarga(Keluarga Keluarga) KeluargaFormatter {
	formatter := KeluargaFormatter{
		Keluarga:  Keluarga,
		CreatedAt: helper.FormatDateToString(Keluarga.CreatedAt),
		UpdatedAt: helper.FormatDateToString(Keluarga.UpdatedAt),
		Penduduk:  penduduk.FormatPenduduks(Keluarga.Penduduk),
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
