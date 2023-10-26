package status

import "time"

type Status struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Status) TableName() string {
	return "tb_status_keluarga"
}
