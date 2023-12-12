package sporadik

import (
	"layanan-kependudukan-api/penduduk"
	"time"
)

type Sporadik struct {
	ID                   int               `json:"id"`
	NIK                  string            `json:"nik"`
	KodeSurat            string            `json:"kode_surat"`
	Keterangan           string            `json:"keterangan"`
	LampiranPemohon      string            `json:"lampiran_pemohon"`
	LampiranSporadikLama string            `json:"lampiran_sporadik_lama"`
	LampiranSporadikBaru string            `json:"lampiran_sporadik_baru"`
	LampiranBukti        string            `json:"lampiran_bukti"`
	LampiranLunasPbb     string            `json:"lampiran_lunas_pbb"`
	Status               bool              `json:"status"`
	CreatedAt            time.Time         `json:"created_at"`
	CreatedBy            int               `json:"created_by"`
	Penduduk             penduduk.Penduduk `json:"penduduk" gorm:"foreignKey:NIK; references:NIK"`
}

func (Sporadik) TableName() string {
	return "tb_sporadik"
}
