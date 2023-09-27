package detail_pengajuan

import "time"

type DetailPengajuan struct {
	ID          int       `json:"id"`
	PengajuanID int       `json:"pengajuan_id"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

func (DetailPengajuan) TableName() string {
	return "tb_detail_pengajuan"
}
