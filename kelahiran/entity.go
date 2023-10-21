package kelahiran

import (
	"time"
)

type Kelahiran struct {
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

func (Kelahiran) TableName() string {
	return "tb_kelahiran"
}
