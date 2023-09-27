package kelurahan

import "time"

type Kelurahan struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Kelurahan) TableName() string {
	return "tb_kelurahan"
}
