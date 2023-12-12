package sktm

import (
	"layanan-kependudukan-api/penduduk"
	"time"
)

type SKTM struct {
	ID         int               `json:"id"`
	NIK        string            `json:"nik"`
	KodeSurat  string            `json:"kode_surat"`
	Keterangan string            `json:"keterangan"`
	Status     bool              `json:"status"`
	CreatedAt  time.Time         `json:"created_at"`
	CreatedBy  int               `json:"created_by"`
	Penduduk   penduduk.Penduduk `json:"penduduk" gorm:"foreignKey:NIK; references:NIK"`
}

func (SKTM) TableName() string {
	return "tb_sktm"
}
