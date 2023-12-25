package aparatur_desa

import (
	"layanan-kependudukan-api/position"
	"time"
)

type AparaturDesa struct {
	ID        int               `json:"id"`
	Nip       string            `json:"nip"`
	Nama      string            `json:"nama"`
	JabatanID int               `json:"jabatan_id"`
	Jabatan   position.Position `json:"jabatan" gorm:"foreignKey:JabatanID; "`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (AparaturDesa) TableName() string {
	return "tb_aparatur_desa"
}
