package province

type ProvinceFormatter struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func FormatProvince(province Province) ProvinceFormatter {
	formatter := ProvinceFormatter{
		ID:   province.ID,
		Code: province.Code,
		Name: province.Name,
	}

	return formatter
}

func FormatProvinces(provinces []Province) []ProvinceFormatter {
	var provincesFormatter []ProvinceFormatter

	for _, province := range provinces {
		provinceFormatter := FormatProvince(province)
		provincesFormatter = append(provincesFormatter, provinceFormatter)
	}

	return provincesFormatter
}
