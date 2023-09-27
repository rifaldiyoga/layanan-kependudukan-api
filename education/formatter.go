package education

type EducationFormatter struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func FormatEducation(education Education) EducationFormatter {
	formatter := EducationFormatter{
		ID:   education.ID,
		Code: education.Code,
		Name: education.Name,
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
