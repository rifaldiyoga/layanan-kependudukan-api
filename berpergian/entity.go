package berpergian

import (
	"layanan-kependudukan-api/berpergian_detail"
	"layanan-kependudukan-api/penduduk"
	"time"
)

type Berpergian struct {
	ID               int                                  `json:"id"`
	NIK              string                               `json:"nik"`
	KodeSurat        string                               `json:"kode_surat"`
	Lampiran         string                               `json:"lampiran"`
	Keterangan       string                               `json:"keterangan"`
	Tujuan           string                               `json:"tujuan"`
	TglBerangkat     time.Time                            `json:"tgl_berangkat"`
	TglKembali       time.Time                            `json:"tgl_kembali"`
	Status           bool                                 `json:"status"`
	CreatedAt        time.Time                            `json:"created_at"`
	CreatedBy        int                                  `json:"created_by"`
	Penduduk         penduduk.Penduduk                    `json:"penduduk" gorm:"foreignKey:NIK; references:NIK"`
	BerpergianDetail []berpergian_detail.BerpergianDetail `json:"berpergian_detail" gorm:"foreignKey:BerpergianID;"`
}

func (Berpergian) TableName() string {
	return "tb_berpergian"
}
