package rt

import "time"

type RT struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RT) TableName() string {
	return "tb_rt"
}
