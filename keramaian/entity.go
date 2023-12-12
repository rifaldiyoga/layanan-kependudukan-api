package keramaian

import (
	"layanan-kependudukan-api/penduduk"
	"time"
)

type Keramaian struct {
	ID         int               `json:"id"`
	NIK        string            `json:"nik"`
	KodeSurat  string            `json:"kode_surat"`
	NamaAcara  string            `json:"nama_acara"`
	Tanggal    string            `json:"tanggal"`
	Waktu      string            `json:"waktu"`
	Tempat     string            `json:"tempat"`
	Alamat     string            `json:"alamat"`
	Telpon     string            `json:"telpon"`
	Lampiran   string            `json:"lampiran"`
	Keterangan string            `json:"keterangan"`
	Status     bool              `json:"status"`
	CreatedAt  time.Time         `json:"created_at"`
	CreatedBy  int               `json:"created_by"`
	Penduduk   penduduk.Penduduk `json:"penduduk" gorm:"foreignKey:NIK; references:NIK"`
}

func (Keramaian) TableName() string {
	return "tb_keramaian"
}
