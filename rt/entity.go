package rt

import (
	"layanan-kependudukan-api/rw"
	"time"
)

type RT struct {
	ID        int
	Code      string
	Name      string
	RwID      int
	RW        rw.RW `json:"rw" gorm:"foreignKey:RwID; "`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RT) TableName() string {
	return "tb_rt"
}
