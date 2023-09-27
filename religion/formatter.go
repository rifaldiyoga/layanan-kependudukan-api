package religion

type ReligionFormatter struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func FormatReligion(religion Religion) ReligionFormatter {
	formatter := ReligionFormatter{
		ID:   religion.ID,
		Code: religion.Code,
		Name: religion.Name,
	}

	return formatter
}

func FormatReligions(religions []Religion) []ReligionFormatter {
	var religionsFormatter []ReligionFormatter

	for _, religion := range religions {
		religionFormatter := FormatReligion(religion)
		religionsFormatter = append(religionsFormatter, religionFormatter)
	}

	return religionsFormatter
}
