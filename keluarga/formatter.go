package keluarga

type KeluargaFormatter struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func FormatKeluarga(keluarga Keluarga) KeluargaFormatter {
	formatter := KeluargaFormatter{
		ID: keluarga.ID,
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
