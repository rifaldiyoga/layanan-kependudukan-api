package keluarga

import "layanan-kependudukan-api/penduduk"

type GetKeluargaDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateKeluargaInput struct {
	NoKK              string                         `json:"no_kk"`
	NIKKepalaKeluarga string                         `json:"nik_kepala_keluarga"`
	KepalaKeluarga    string                         `json:"kepala_keluarga"`
	Address           string                         `json:"alamat"`
	RtID              int                            `json:"rt_id"`
	RwID              int                            `json:"rw_id"`
	KelurahanID       int                            `json:"kelurahan_id"`
	KecamatanID       int                            `json:"subdistrict_id"`
	KotaID            int                            `json:"district_id"`
	ProvinsiID        int                            `json:"province_id"`
	Penduduk          []penduduk.CreatePendudukInput `json:"penduduk"`
}
