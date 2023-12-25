package position

import "time"

type Position struct {
	ID        int
	Code      string
	Jabatan   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Position) TableName() string {
	return "tb_jabatan"
}
