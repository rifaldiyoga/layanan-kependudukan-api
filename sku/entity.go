package sku

import (
	"layanan-kependudukan-api/penduduk"
	"time"
)

type SKU struct {
	ID         int               `json:"id"`
	NIK        string            `json:"nik"`
	KodeSurat  string            `json:"kode_surat"`
	Usaha      string            `json:"usaha"`
	Keterangan string            `json:"keterangan"`
	Status     bool              `json:"status"`
	Lampiran   string            `json:"lampiran"`
	CreatedAt  time.Time         `json:"created_at"`
	CreatedBy  int               `json:"created_by"`
	Penduduk   penduduk.Penduduk `json:"penduduk" gorm:"foreignKey:NIK; references:NIK"`
}

func (SKU) TableName() string {
	return "tb_sku"
}
