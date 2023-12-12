package district

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/province"
)

type DistrictFormatter struct {
	ID         int                        `json:"id"`
	Code       string                     `json:"code"`
	Name       string                     `json:"name"`
	ProvinceID int                        `json:"province_id"`
	CreatedAt  string                     `json:"created_at"`
	Provinsi   province.ProvinceFormatter `json:"provinsi" `
}

func FormatDistrict(district District) DistrictFormatter {
	formatter := DistrictFormatter{
		ID:         district.ID,
		Code:       district.Code,
		Name:       district.Name,
		ProvinceID: district.ProvinsiID,
		CreatedAt:  helper.FormatDateToString(district.CreatedAt),
		Provinsi:   province.FormatProvince(district.Provinsi),
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
