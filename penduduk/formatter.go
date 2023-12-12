package penduduk

import (
	"layanan-kependudukan-api/district"
	"layanan-kependudukan-api/education"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/job"
	"layanan-kependudukan-api/kelurahan"
	"layanan-kependudukan-api/province"
	"layanan-kependudukan-api/religion"
	"layanan-kependudukan-api/rt"
	"layanan-kependudukan-api/rw"
	"layanan-kependudukan-api/subdistrict"
)

type PendudukFormatter struct {
	Penduduk
	BirthDate   string                           `json:"birth_date"`
	CreatedAt   string                           `json:"created_at"`
	UpdatedAt   string                           `json:"updated_at"`
	Religion    religion.ReligionFormatter       `json:"religion"`
	Education   education.EducationFormatter     `json:"education"`
	Job         job.JobFormatter                 `json:"job"`
	RT          rt.RTFormatter                   `json:"rt"`
	RW          rw.RWFormatter                   `json:"rw"`
	Keluarhan   kelurahan.KelurahanFormatter     `json:"kelurahan"`
	SubDistrict subdistrict.SubDistrictFormatter `json:"kecamatan"`
	District    district.DistrictFormatter       `json:"kota"`
	Provinsi    province.ProvinceFormatter       `json:"provinsi"`
}

func FormatPenduduk(penduduk Penduduk) PendudukFormatter {
	formatter := PendudukFormatter{
		Penduduk:    penduduk,
		Religion:    religion.FormatReligion(penduduk.Religion),
		Education:   education.FormatEducation(penduduk.Education),
		Job:         job.FormatJob(penduduk.Job),
		RT:          rt.FormatRT(penduduk.RT),
		RW:          rw.FormatRW(penduduk.RW),
		Keluarhan:   kelurahan.FormatKelurahan(penduduk.Kelurahan),
		SubDistrict: subdistrict.FormatSubDistrict(penduduk.Kecamatan),
		District:    district.FormatDistrict(penduduk.Kota),
		Provinsi:    province.FormatProvince(penduduk.Provinsi),
		BirthDate:   helper.FormatDateToString(penduduk.BirthDate),
		CreatedAt:   helper.FormatDateToString(penduduk.CreatedAt),
		UpdatedAt:   helper.FormatDateToString(penduduk.UpdatedAt),
	}

	return formatter
}

func FormatPenduduks(penduduks []Penduduk) []PendudukFormatter {
	var penduduksFormatter []PendudukFormatter

	for _, penduduk := range penduduks {
		pendudukFormatter := FormatPenduduk(penduduk)
		penduduksFormatter = append(penduduksFormatter, pendudukFormatter)
	}

	return penduduksFormatter
}
