package keluarga

type GetKeluargaDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateKeluargaInput struct {
	NoKK              string `json:"no_kk"`
	NIKKepalaKeluarga string `json:"nik_kepala_keluarga"`
	KepalaKeluarga    string `json:"kepala_keluarga"`
	Address           string `json:"alamat"`
	RtID              int    `json:"rt_id"`
	RwID              int    `json:"rw_id"`
	KelurahanID       int    `json:"kelurahan_id"`
	KecamatanID       int    `json:"subdistrict_id"`
}
