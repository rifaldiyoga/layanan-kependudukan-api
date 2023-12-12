package kepolisian

type GetKepolisianDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateKepolisianInput struct {
	NIK        string `form:"nik"`
	KodeSurat  string `form:"kode_surat"`
	Lampiran   string `form:"lampiran_path" `
	Keterangan string `form:"keterangan" binding:"required"`
}
