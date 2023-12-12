package kelurahan

import (
	"layanan-kependudukan-api/subdistrict"
	"time"
)

type Kelurahan struct {
	ID          int
	Code        string
	Name        string
	KecamatanID int
	Kecamatan   subdistrict.SubDistrict `json:"kecamatan" gorm:"foreignKey:KecamatanID; "`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Kelurahan) TableName() string {
	return "tb_kelurahan"
}
