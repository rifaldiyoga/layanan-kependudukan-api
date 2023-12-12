package pindah

import (
	"layanan-kependudukan-api/district"
	"layanan-kependudukan-api/penduduk"
	"layanan-kependudukan-api/pindah_detail"
	"layanan-kependudukan-api/province"
	"layanan-kependudukan-api/subdistrict"
	"time"
)

type Pindah struct {
	ID                int                          `json:"id"`
	NIK               string                       `json:"nik"`
	NikKepalaKeluarga string                       `json:"nik_kepala_keluarga"`
	KodeSurat         string                       `json:"kode_surat"`
	Type              string                       `json:"type"`
	AlasanPindah      string                       `json:"alasan_pindah"`
	AlamatTujuan      string                       `json:"alamat_tujuan"`
	Rt                string                       `json:"rt"`
	Rw                string                       `json:"rw"`
	Kelurahan         string                       `json:"kelurahan"`
	KecamatanID       int                          `json:"subdistrict_id"`
	Kecamatan         subdistrict.SubDistrict      `json:"kecamatan" gorm:"foreignKey:KecamatanID; preload:true"`
	KotaID            int                          `json:"district_id"`
	Kota              district.District            `json:"kota" gorm:"foreignKey:KotaID; preload:true"`
	ProvinsiID        int                          `json:"province_id"`
	Provinsi          province.Province            `json:"provinsi" gorm:"foreignKey:ProvinsiID; preload:true"`
	KodePos           string                       `json:"kode_pos"`
	Telepon           string                       `json:"telepon"`
	JenisKepindahan   string                       `json:"jenis_kepindahan"`
	StatusTidakPindah string                       `json:"status_tidak_pindah"`
	StatusPindah      string                       `json:"status_pindah"`
	Status            bool                         `json:"status"`
	Lampiran          string                       `json:"lampiran"`
	CreatedAt         time.Time                    `json:"created_at"`
	CreatedBy         int                          `json:"created_by"`
	Penduduk          penduduk.Penduduk            `json:"penduduk" gorm:"foreignKey:NIK; references:NIK"`
	PindahDetail      []pindah_detail.PindahDetail `json:"detail" gorm:"foreignKey:PindahID"`
}

func (Pindah) TableName() string {
	return "tb_pindah"
}
