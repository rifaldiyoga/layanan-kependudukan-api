package subdistrict

import (
	"layanan-kependudukan-api/district"
	"layanan-kependudukan-api/helper"
)

type SubDistrictFormatter struct {
	ID        int                        `json:"id"`
	Code      string                     `json:"code"`
	Name      string                     `json:"name"`
	CreatedAt string                     `json:"created_at"`
	KotaID    int                        `json:"kota_id"`
	Kota      district.DistrictFormatter `json:"kota"`
}

func FormatSubDistrict(subDistrict SubDistrict) SubDistrictFormatter {
	formatter := SubDistrictFormatter{
		ID:        subDistrict.ID,
		Code:      subDistrict.Code,
		Name:      helper.CapitalizeEachWord(subDistrict.Name),
		CreatedAt: helper.FormatDateToString(subDistrict.CreatedAt),
		KotaID:    subDistrict.KotaID,
		Kota:      district.FormatDistrict(subDistrict.Kota),
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
