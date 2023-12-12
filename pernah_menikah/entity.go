package pernah_menikah

import (
	"layanan-kependudukan-api/penduduk"
	"time"
)

type PernahMenikah struct {
	ID         int               `json:"id"`
	NIK        string            `json:"nik"`
	NikSuami   string            `json:"nik_suami"`
	NikIstri   string            `json:"nik_istri"`
	KodeSurat  string            `json:"kode_surat"`
	Lampiran   string            `json:"lampiran"`
	Keterangan string            `json:"keterangan"`
	Type       string            `json:"type"`
	Status     bool              `json:"status"`
	CreatedAt  time.Time         `json:"created_at"`
	CreatedBy  int               `json:"created_by"`
	Penduduk   penduduk.Penduduk `json:"penduduk" gorm:"foreignKey:NIK; references:NIK"`
	Suami      penduduk.Penduduk `json:"suami" gorm:"foreignKey:NikSuami; references:NIK"`
	Istri      penduduk.Penduduk `json:"istri" gorm:"foreignKey:NikIstri; references:NIK"`
}

func (PernahMenikah) TableName() string {
	return "tb_pernah_menikah"
}
