package religion

import "time"

type Religion struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Religion) TableName() string {
	return "tb_agama"
}
