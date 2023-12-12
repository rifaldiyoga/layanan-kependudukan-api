package kematian

import (
	"layanan-kependudukan-api/penduduk"
	"time"
)

type Kematian struct {
	ID            int               `json:"id"`
	NIK           string            `json:"nik"`
	NikJenazah    string            `json:"nik_jenazah"`
	KodeSurat     string            `json:"kode_surat"`
	Keterangan    string            `json:"keterangan"`
	TglKematian   time.Time         `json:"tgl_kematian"`
	Jam           string            `json:"jam"`
	Sebab         string            `json:"sebab"`
	Tempat        string            `json:"tempat"`
	Status        bool              `json:"status"`
	Saksi1        string            `json:"saksi1"`
	Saksi2        string            `json:"saksi2"`
	LampiranKetRs string            `json:"lampiran_ket_rs"`
	CreatedAt     time.Time         `json:"created_at"`
	CreatedBy     int               `json:"created_by"`
	Penduduk      penduduk.Penduduk `json:"penduduk" gorm:"foreignKey:NIK; references:NIK"`
	Jenazah       penduduk.Penduduk `json:"jenazah" gorm:"foreignKey:NikJenazah; references:NIK"`
}

func (Kematian) TableName() string {
	return "tb_kematian"
}
