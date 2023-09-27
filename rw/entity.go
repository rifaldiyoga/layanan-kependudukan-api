package rw

import "time"

type RW struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RW) TableName() string {
	return "tb_rw"
}
