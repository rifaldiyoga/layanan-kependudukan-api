package penghasilan

import (
	"layanan-kependudukan-api/penduduk"
	"time"
)

type Penghasilan struct {
	ID          int               `json:"id"`
	NIK         string            `json:"nik"`
	KodeSurat   string            `json:"kode_surat"`
	Lampiran    string            `json:"lampiran"`
	Penghasilan string            `json:"penghasilan"`
	Keterangan  string            `json:"keterangan"`
	Status      bool              `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	CreatedBy   int               `json:"created_by"`
	Penduduk    penduduk.Penduduk `json:"penduduk" gorm:"foreignKey:NIK; references:NIK"`
}

func (Penghasilan) TableName() string {
	return "tb_penghasilan"
}
