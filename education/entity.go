package education

import "time"

type Education struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Education) TableName() string {
	return "tb_pendidikan"
}
