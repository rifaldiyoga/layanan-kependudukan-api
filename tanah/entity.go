package tanah

import (
	"layanan-kependudukan-api/penduduk"
	"time"
)

type Tanah struct {
	ID           int               `json:"id"`
	NIK          string            `json:"nik"`
	KodeSurat    string            `json:"kode_surat"`
	Lampiran     string            `json:"lampiran"`
	Type         string            `json:"type"`
	Keterangan   string            `json:"keterangan"`
	Lokasi       string            `json:"lokasi"`
	LuasTanah    string            `json:"luas_tanah"`
	Panjang      int               `json:"panjang"`
	Lebar        int               `json:"lebar"`
	BatasBarat   string            `json:"batas_barat"`
	BatasTimur   string            `json:"batas_timur"`
	BatasUtara   string            `json:"batas_utara"`
	BatasSelatan string            `json:"batas_selatan"`
	Saksi1       string            `json:"saksi1"`
	Saksi2       string            `json:"saksi2"`
	Status       bool              `json:"status"`
	CreatedAt    time.Time         `json:"created_at"`
	CreatedBy    int               `json:"created_by"`
	Penduduk     penduduk.Penduduk `json:"penduduk" gorm:"foreignKey:NIK; references:NIK"`
}

func (Tanah) TableName() string {
	return "tb_tanah"
}
