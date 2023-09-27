package pengajuan

import (
	"layanan-kependudukan-api/detail_pengajuan"
	"time"
)

type Pengajuan struct {
	ID         int                                `json:"id"`
	Code       string                             `json:"code"`
	Name       string                             `json:"name"`
	Layanan    string                             `json:"layanan"`
	LayananID  int                                `json:"layanan_id"`
	Status     string                             `json:"status"`
	Keterangan string                             `json:"keterangan"`
	CreatedBy  int                                `json:"created_by"`
	CreatedAt  time.Time                          `json:"created_at"`
	UpdatedAt  time.Time                          `json:"updated_at"`
	Detail     []detail_pengajuan.DetailPengajuan `json:"detail" gorm:"foreignKey:PengajuanID"`
}

func (Pengajuan) TableName() string {
	return "tb_pengajuan"
}
