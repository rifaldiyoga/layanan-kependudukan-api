package kelahiran

type KelahiranFormatter struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func FormatKelahiran(kelahiran Kelahiran) KelahiranFormatter {
	formatter := KelahiranFormatter{
		ID: kelahiran.ID,
	}

	return formatter
}

func FormatKelahirans(kelahirans []Kelahiran) []KelahiranFormatter {
	var kelahiransFormatter []KelahiranFormatter

	for _, kelahiran := range kelahirans {
		kelahiranFormatter := FormatKelahiran(kelahiran)
		kelahiransFormatter = append(kelahiransFormatter, kelahiranFormatter)
	}

	return kelahiransFormatter
}
