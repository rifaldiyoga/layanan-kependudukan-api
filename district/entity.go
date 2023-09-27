package district

import "time"

type District struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (District) TableName() string {
	return "tb_kota"
}
