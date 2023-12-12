package tanah

type GetTanahDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateTanahInput struct {
	NIK          string `form:"nik"`
	KodeSurat    string `form:"kode_surat"`
	Lampiran     string `form:"lampiran_path"`
	Keterangan   string `form:"keterangan" binding:"required"`
	Type         string `form:"type" binding:"required"`
	Lokasi       string `form:"lokasi" binding:"required"`
	LuasTanah    string `form:"luas_tanah" binding:"required"`
	Panjang      string `form:"panjang" `
	Lebar        string `form:"lebar" `
	BatasBarat   string `form:"batas_barat" binding:"required"`
	BatasTimur   string `form:"batas_timur" binding:"required"`
	BatasUtara   string `form:"batas_utara" binding:"required"`
	BatasSelatan string `form:"batas_selatan" binding:"required"`
	Saksi1       string `form:"saksi1" `
	Saksi2       string `form:"saksi2" `
}
