package kelurahan

type KelurahanFormatter struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func FormatKelurahan(kelurahan Kelurahan) KelurahanFormatter {
	formatter := KelurahanFormatter{
		ID:   kelurahan.ID,
		Code: kelurahan.Code,
		Name: kelurahan.Name,
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
