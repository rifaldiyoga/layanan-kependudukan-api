package rw

import (
	"layanan-kependudukan-api/kelurahan"
	"time"
)

type RW struct {
	ID          int
	Code        string
	Name        string
	KelurahanID int
	Kelurahan   kelurahan.Kelurahan `json:"kelurahan" gorm:"foreignKey:KelurahanID; "`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (RW) TableName() string {
	return "tb_rw"
}
