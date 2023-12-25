package aparatur_desa

import (
	"layanan-kependudukan-api/position"
	"time"
)

type AparaturDesa struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	JabatanID int
	Jabatan   position.Position `json:"jabatan" gorm:"foreignKey:JabatanID; "`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (AparaturDesa) TableName() string {
	return "tb_aparatur_desa"
}
