package district

import (
	"layanan-kependudukan-api/province"
	"time"
)

type District struct {
	ID         int
	Code       string
	Name       string
	ProvinsiID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Provinsi   province.Province `json:"provinsi" gorm:"foreignKey:ProvinsiID;"`
}

func (District) TableName() string {
	return "tb_kota"
}
