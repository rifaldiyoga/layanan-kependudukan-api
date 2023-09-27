package province

import "time"

type Province struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Province) TableName() string {
	return "tb_provinsi"
}
