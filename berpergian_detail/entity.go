package berpergian_detail

import "time"

type BerpergianDetail struct {
	ID           int       `json:"id"`
	NIK          string    `json:"nik"`
	Nama         string    `json:"nama"`
	StatusFamily string    `json:"status_family"`
	BerpergianID int       `json:"berpergian_id"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    int       `json:"created_by"`
}

func (BerpergianDetail) TableName() string {
	return "tb_berpergian_detail"
}
