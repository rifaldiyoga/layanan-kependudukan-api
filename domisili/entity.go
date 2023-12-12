package domisili

import (
	"layanan-kependudukan-api/penduduk"
	"time"
)

type Domisili struct {
	ID              int               `json:"id"`
	NIK             string            `json:"nik"`
	KodeSurat       string            `json:"kode_surat"`
	Type            string            `json:"type"`
	NamaPerusahaan  string            `json:"nama_perusahaan"`
	JenisPerusahaan string            `json:"jenis_perusahaan"`
	TelpPerusahaan  string            `json:"telp_perusahaan"`
	Alamat          string            `json:"alamat_perusahaan"`
	StatusBangunan  string            `json:"status_bangunan"`
	AktaPerusahaan  string            `json:"akta_perusahaan"`
	SKPengesahan    string            `json:"sk_pengesahan"`
	PenanggungJawab string            `json:"penanggung_jawab"`
	Lampiran        string            `json:"lampiran"`
	Keterangan      string            `json:"keterangan"`
	Status          bool              `json:"status"`
	CreatedAt       time.Time         `json:"created_at"`
	CreatedBy       int               `json:"created_by"`
	Penduduk        penduduk.Penduduk `json:"penduduk" gorm:"foreignKey:NIK; references:NIK"`
}

func (Domisili) TableName() string {
	return "tb_domisili"
}
