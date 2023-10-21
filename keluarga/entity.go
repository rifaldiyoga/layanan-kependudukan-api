package keluarga

import (
	"time"
)

type Keluarga struct {
	ID           int
	Nik          string
	FullName     string
	BirthPlace   string
	BirthDate    time.Time
	ReligionID   int
	PendidikanID int
	PekerjaanID  int
	Nationality  string
	MariedType   string
	RtID         int
	RwID         int
	KelurahanID  int
	KecamatanID  int
	KotaID       int
	ProvinsiID   int
	jk           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Keluarga) TableName() string {
	return "tb_kartu_keluarga"
}
