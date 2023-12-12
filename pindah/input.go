package pindah

type GetPindahDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreatePindahInput struct {
	NIK               string `form:"nik"`
	Type              string `form:"type" binding:"required"`
	NikKepalaKeluarga string `form:"nik_kepala_keluarga"`
	KodeSurat         string `form:"kode_surat"`
	AlasanPindah      string `form:"alasan_pindah" binding:"required"`
	AlamatTujuan      string `form:"alamat_tujuan" binding:"required"`
	Rt                string `form:"rt" binding:"required"`
	Rw                string `form:"rw" binding:"required"`
	Kelurahan         string `form:"kelurahan" binding:"required"`
	KecamatanID       int    `form:"subdistrict_id" binding:"required"`
	KotaID            int    `form:"district_id" binding:"required"`
	ProvinsiID        int    `form:"province_id" binding:"required"`
	KodePos           string `form:"kode_pos" binding:"required"`
	Telepon           string `form:"telepon" binding:"required"`
	Lampiran          string `form:"lampiran_path"`
	JenisKepindahan   string `form:"jenis_kepindahan" binding:"required"`
	StatusTidakPindah string `form:"status_tidak_pindah" binding:"required"`
	StatusPindah      string `form:"status_pindah" binding:"required"`
}
