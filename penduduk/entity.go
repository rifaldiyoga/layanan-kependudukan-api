package penduduk

import (
	"time"
)

type Penduduk struct {
	ID           int
	NIK          string
	NoKK         string
	Fullname     string
	BirthPlace   string
	BirthDate    time.Time
	ReligionID   int
	PendidikanID int
	PekerjaanID  int
	Nationality  string
	MariedType   string
	MariedDate   time.Time
	BloodType    string
	RtID         int
	RwID         int
	KelurahanID  int
	KecamatanID  int
	KotaID       int
	ProvinsiID   int
	JK           string
	Address      string
	StatusFamily string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Penduduk) TableName() string {
	return "tb_penduduk"
}
