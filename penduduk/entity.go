package penduduk

import (
	"layanan-kependudukan-api/district"
	"layanan-kependudukan-api/education"
	"layanan-kependudukan-api/job"
	"layanan-kependudukan-api/kelurahan"
	"layanan-kependudukan-api/province"
	"layanan-kependudukan-api/religion"
	"layanan-kependudukan-api/rt"
	"layanan-kependudukan-api/rw"
	"layanan-kependudukan-api/subdistrict"
	"time"
)

type Penduduk struct {
	ID           int                     `json:"id"`
	NIK          string                  `json:"nik"`
	NoKK         string                  `json:"no_kk"`
	Fullname     string                  `json:"fullname"`
	BirthPlace   string                  `json:"birth_place"`
	BirthDate    time.Time               `json:"birth_date"`
	ReligionID   int                     `json:"religion_id"`
	Religion     religion.Religion       `json:"religion" gorm:"foreignKey:ReligionID; preload:true"`
	PendidikanID int                     `json:"education_id"`
	Education    education.Education     `json:"education" gorm:"foreignKey:PendidikanID; preload:true"`
	PekerjaanID  int                     `json:"job_id"`
	Job          job.Job                 `json:"job" gorm:"foreignKey:PekerjaanID; preload:true"`
	Nationality  string                  `json:"nationality"`
	MariedType   string                  `json:"maried_type"`
	MariedDate   time.Time               `json:"maried_date"`
	BloodType    string                  `json:"blood_type"`
	Address      string                  `json:"address"`
	RtID         int                     `json:"rt_id"`
	RT           rt.RT                   `json:"rt" gorm:"foreignKey:RtID; preload:true"`
	RwID         int                     `json:"rw_id"`
	RW           rw.RW                   `json:"rw" gorm:"foreignKey:RwID; preload:true"`
	KelurahanID  int                     `json:"kelurahan_id"`
	Kelurahan    kelurahan.Kelurahan     `json:"kelurahan" gorm:"foreignKey:KelurahanID; preload:true"`
	KecamatanID  int                     `json:"subdistrict_id"`
	Kecamatan    subdistrict.SubDistrict `json:"kecamatan" gorm:"foreignKey:KecamatanID; preload:true"`
	KotaID       int                     `json:"district_id"`
	Kota         district.District       `json:"kota" gorm:"foreignKey:KotaID; preload:true"`
	ProvinsiID   int                     `json:"province_id"`
	Provinsi     province.Province       `json:"provinsi" gorm:"foreignKey:ProvinsiID; preload:true"`
	JK           string                  `json:"jk"`
	Status       bool                    `json:"status"`
	StatusFamily string                  `json:"status_family"`
	CreatedAt    time.Time               `json:"created_at"`
	UpdatedAt    time.Time               `json:"updated_at"`
}

func (Penduduk) TableName() string {
	return "tb_penduduk"
}
