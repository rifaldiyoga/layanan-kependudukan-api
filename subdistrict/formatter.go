package subdistrict

type SubDistrictFormatter struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func FormatSubDistrict(subDistrict SubDistrict) SubDistrictFormatter {
	formatter := SubDistrictFormatter{
		ID:   subDistrict.ID,
		Code: subDistrict.Code,
		Name: subDistrict.Name,
	}

	return formatter
}

func FormatSubDistricts(subDistricts []SubDistrict) []SubDistrictFormatter {
	var subDistrictsFormatter []SubDistrictFormatter

	for _, subDistrict := range subDistricts {
		subDistrictFormatter := FormatSubDistrict(subDistrict)
		subDistrictsFormatter = append(subDistrictsFormatter, subDistrictFormatter)
	}

	return subDistrictsFormatter
}
