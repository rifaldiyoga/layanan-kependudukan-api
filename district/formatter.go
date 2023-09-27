package district

type DistrictFormatter struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func FormatDistrict(district District) DistrictFormatter {
	formatter := DistrictFormatter{
		ID:   district.ID,
		Code: district.Code,
		Name: district.Name,
	}

	return formatter
}

func FormatDistricts(districts []District) []DistrictFormatter {
	var districtsFormatter []DistrictFormatter

	for _, district := range districts {
		districtFormatter := FormatDistrict(district)
		districtsFormatter = append(districtsFormatter, districtFormatter)
	}

	return districtsFormatter
}
