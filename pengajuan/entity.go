package pengajuan

import (
	"layanan-kependudukan-api/penduduk"
	"layanan-kependudukan-api/pengajuan_detail"
	"time"
)

type Pengajuan struct {
	ID         int                                `json:"id"`
	Code       string                             `json:"code"`
	Name       string                             `json:"name"`
	Layanan    string                             `json:"layanan"`
	LayananID  int                                `json:"layanan_id"`
	Status     string                             `json:"status"`
	NIK        string                             `json:"nik"`
	Keterangan string                             `json:"keterangan"`
	Note       string                             `json:"note"`
	RefID      int                                `json:"ref_id"`
	CreatedBy  int                                `json:"created_by"`
	CreatedAt  time.Time                          `json:"created_at"`
	UpdatedAt  time.Time                          `json:"updated_at"`
	Detail     []pengajuan_detail.DetailPengajuan `json:"detail" gorm:"foreignKey:PengajuanID"`
	Penduduk   penduduk.Penduduk                  `json:"penduduk" gorm:"foreignKey:NIK; references:NIK"`
}

func (Pengajuan) TableName() string {
	return "tb_pengajuan"
}
