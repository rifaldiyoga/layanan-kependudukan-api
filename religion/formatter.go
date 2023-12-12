package religion

import "layanan-kependudukan-api/helper"

type ReligionFormatter struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func FormatReligion(religion Religion) ReligionFormatter {
	formatter := ReligionFormatter{
		ID:        religion.ID,
		Code:      religion.Code,
		Name:      religion.Name,
		CreatedAt: helper.FormatDateToString(religion.CreatedAt),
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
