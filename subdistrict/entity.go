package subdistrict

import "time"

type SubDistrict struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SubDistrict) TableName() string {
	return "tb_kecamatan"
}
