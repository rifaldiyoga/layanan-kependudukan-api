package pernah_menikah

type GetPernahMenikahDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreatePernahMenikahInput struct {
	NIK        string `form:"nik"`
	NIKSuami   string `form:"nik_suami"`
	NIKIstri   string `form:"nik_istri"`
	KodeSurat  string `form:"kode_surat"`
	Lampiran   string `form:"lampiran_path" `
	Keterangan string `form:"keterangan" binding:"required"`
	Type       string `form:"type" `
}
