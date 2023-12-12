package belum_menikah

type GetBelumMenikahDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateBelumMenikahInput struct {
	NIK        string `form:"nik"`
	KodeSurat  string `form:"kode_surat"`
	Lampiran   string `form:"lampiran_path" `
	Keterangan string `form:"keterangan" binding:"required"`
}
