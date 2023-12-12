package pindah_detail

import "time"

type PindahDetail struct {
	ID           int       `json:"id"`
	NIK          string    `json:"nik"`
	Nama         string    `json:"nama"`
	StatusFamily string    `json:"status_family"`
	PindahID     int       `json:"pindah_id"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    int       `json:"created_by"`
}

func (PindahDetail) TableName() string {
	return "tb_pindah_detail"
}
