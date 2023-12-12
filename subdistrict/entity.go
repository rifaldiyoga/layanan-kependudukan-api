package subdistrict

import (
	"layanan-kependudukan-api/district"
	"time"
)

type SubDistrict struct {
	ID        int
	Code      string
	Name      string
	KotaID    int
	Kota      district.District `json:"kota" gorm:"foreignKey:KotaID;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (SubDistrict) TableName() string {
	return "tb_kecamatan"
}
