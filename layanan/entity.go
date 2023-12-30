package layanan

import "time"

type Layanan struct {
	ID        int
	Code      string
	Name      string
	Type      string
	IsConfirm bool
	IsSign    bool
	Info      string
	KodeSurat string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Layanan) TableName() string {
	return "tb_layanan"
}
