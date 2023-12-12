package kelahiran

import (
	"layanan-kependudukan-api/penduduk"
	"time"
)

type Kelahiran struct {
	ID                int       `json:"id"`
	Nama              string    `json:"nama"`
	NIK               string    `json:"nik"`
	KodeSurat         string    `json:"kode_surat"`
	BirthDate         time.Time `json:"birth_date"`
	BirthPlace        string    `json:"birth_place"`
	AnakKe            int       `json:"anak_ke"`
	Jam               string    `json:"jam"`
	JK                string    `json:"jk"`
	NikAyah           string    `json:"nik_ayah"`
	NikIbu            string    `json:"nik_ibu"`
	LampiranBukuNikah string    `json:"lampiran_buku_nikah"`
	LampiranKetRs     string    `json:"lampiran_ket_rs"`
	KecamatanID       int       `json:"subdistrict_id"`
	// Kecamatan         subdistrict.SubDistrict `json:"kecamatan" gorm:"foreignKey:KecamatanID; preload:true"`
	KotaID int `json:"district_id"`
	// Kota              district.District       `json:"kota" gorm:"foreignKey:KotaID; preload:true"`
	ProvinsiID int `json:"province_id"`
	// Provinsi          province.Province       `json:"provinsi" gorm:"foreignKey:ProvinsiID; preload:true"`
	Keterangan string            `json:"keterangan"`
	Status     bool              `json:"status"`
	CreatedAt  time.Time         `json:"created_at"`
	CreatedBy  int               `json:"created_by"`
	Penduduk   penduduk.Penduduk `json:"penduduk" gorm:"foreignKey:NIK; references:NIK"`
	Ayah       penduduk.Penduduk `json:"ayah" gorm:"foreignKey:NikAyah; references:NIK"`
	Ibu        penduduk.Penduduk `json:"ibu" gorm:"foreignKey:NikIbu; references:NIK"`
}

func (Kelahiran) TableName() string {
	return "tb_kelahiran"
}
