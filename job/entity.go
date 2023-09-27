package job

import "time"

type Job struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Job) TableName() string {
	return "tb_pekerjaan"
}
