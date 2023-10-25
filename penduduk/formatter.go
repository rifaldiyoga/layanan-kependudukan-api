package penduduk

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type PendudukFormatter struct {
	ID            int       `json:"id"`
	NIK           string    `json:"nik"`
	NoKK          string    `json:"no_kk"`
	Fullname      string    `json:"fullname"`
	BirthPlace    string    `json:"birth_place"`
	BirthDate     time.Time `json:"birth_date"`
	ReligionID    int       `json:"religion_id"`
	PendidikanID  int       `json:"education_id"`
	PekerjaanID   int       `json:"job_id"`
	Nationality   string    `json:"nationality"`
	MariedType    string    `json:"maried_type"`
	MariedDate    time.Time `json:"maried_date"`
	BloodType     string    `json:"blood_type"`
	Address       string    `json:"address"`
	RtID          int       `json:"rt_id"`
	RwID          int       `json:"rw_id"`
	KelurahanID   int       `json:"kelurahan_id"`
	SubDistrictID int       `json:"subdistrict_id"`
	DistictID     int       `json:"district_id"`
	ProvinceID    int       `json:"province_id"`
	JK            string    `json:"jk"`
	StatusFamily  string    `json:"status_family"`
	CreatedAt     string    `json:"created_at"`
	UpdatedAt     string    `json:"updated_at"`
}

func FormatPenduduk(penduduk Penduduk) PendudukFormatter {
	formatter := PendudukFormatter{
		ID:            penduduk.ID,
		NIK:           penduduk.NIK,
		Fullname:      penduduk.Fullname,
		BirthPlace:    penduduk.BirthPlace,
		BirthDate:     penduduk.BirthDate,
		ReligionID:    penduduk.ReligionID,
		PekerjaanID:   penduduk.PekerjaanID,
		PendidikanID:  penduduk.PendidikanID,
		Nationality:   penduduk.Nationality,
		MariedType:    penduduk.MariedType,
		MariedDate:    penduduk.MariedDate,
		BloodType:     penduduk.BloodType,
		Address:       penduduk.Address,
		RtID:          penduduk.RtID,
		RwID:          penduduk.RwID,
		KelurahanID:   penduduk.KelurahanID,
		SubDistrictID: penduduk.KecamatanID,
		DistictID:     penduduk.KotaID,
		ProvinceID:    penduduk.ProvinsiID,
		JK:            penduduk.JK,
		NoKK:          penduduk.NoKK,
		StatusFamily:  penduduk.StatusFamily,
		CreatedAt:     helper.FormatDateToString(penduduk.CreatedAt),
		UpdatedAt:     helper.FormatDateToString(penduduk.UpdatedAt),
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
