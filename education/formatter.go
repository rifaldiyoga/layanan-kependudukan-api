package education

import "layanan-kependudukan-api/helper"

type EducationFormatter struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func FormatEducation(education Education) EducationFormatter {
	formatter := EducationFormatter{
		ID:        education.ID,
		Code:      education.Code,
		Name:      education.Name,
		CreatedAt: helper.FormatDateToString(education.CreatedAt),
	}

	return formatter
}

func FormatEducations(educations []Education) []EducationFormatter {
	var educationsFormatter []EducationFormatter

	for _, education := range educations {
		educationFormatter := FormatEducation(education)
		educationsFormatter = append(educationsFormatter, educationFormatter)
	}

	return educationsFormatter
}
