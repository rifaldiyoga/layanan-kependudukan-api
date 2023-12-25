package sistem

import (
	"layanan-kependudukan-api/district"
	"layanan-kependudukan-api/province"
	"layanan-kependudukan-api/subdistrict"
)

type Sistem struct {
	ID          int                     `json:"id"`
	Code        string                  `json:"code"`
	Nama        string                  `json:"nama"`
	Alamat      string                  `json:"alamat"`
	Telp        string                  `json:"telp"`
	KodePos     string                  `json:"kode_pos"`
	KecamatanID int                     `json:"subdistrict_id"`
	Kecamatan   subdistrict.SubDistrict `json:"kecamatan" gorm:"foreignKey:KecamatanID; preload:true"`
	KotaID      int                     `json:"district_id"`
	Kota        district.District       `json:"kota" gorm:"foreignKey:KotaID; preload:true"`
	ProvinsiID  int                     `json:"province_id"`
	Provinsi    province.Province       `json:"provinsi" gorm:"foreignKey:ProvinsiID; preload:true"`
}

func (Sistem) TableName() string {
	return "tb_sistem"
}
