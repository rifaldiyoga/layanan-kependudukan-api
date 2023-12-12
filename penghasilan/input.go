package penghasilan

type GetPenghasilanDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreatePenghasilanInput struct {
	NIK         string `form:"nik"`
	KodeSurat   string `form:"kode_surat"`
	Lampiran    string `form:"lampiran_path" `
	Keterangan  string `form:"keterangan" binding:"required"`
	Penghasilan string `form:"penghasilan" binding:"required"`
}
