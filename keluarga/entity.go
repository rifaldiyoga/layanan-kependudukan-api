package keluarga

import (
	"time"
)

type Keluarga struct {
	ID                int       `json:"id"`
	NoKK              string    `json:"no_kk"`
	NIKKepalaKeluarga string    `json:"nik_kepala_keluarga"`
	KepalaKeluarga    string    `json:"kepala_keluarga"`
	Address           string    `json:"alamat"`
	RtID              int       `json:"rt_id"`
	RwID              int       `json:"rw_id"`
	KelurahanID       int       `json:"kelurahan_id"`
	KecamatanID       int       `json:"subdistrict_id"`
	KotaID            int       `json:"district_iid"`
	ProvinsiID        int       `json:"province_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (Keluarga) TableName() string {
	return "tb_kartu_keluarga"
}
