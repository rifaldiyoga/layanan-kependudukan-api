package rumah

type GetRumahDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateRumahInput struct {
	NIK        string `form:"nik"`
	KodeSurat  string `form:"kode_surat"`
	Lampiran   string `form:"lampiran_path" `
	Keterangan string `form:"keterangan" binding:"required"`
}
