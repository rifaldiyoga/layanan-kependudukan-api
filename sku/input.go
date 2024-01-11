package sku

type GetSKUDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateSKUInput struct {
	KodeSurat  string `form:"kode_surat"`
	Usaha      string `form:"usaha" binding:"required"`
	Keterangan string `form:"keterangan" binding:"required"`
	NIK        string `form:"nik" `
	Status     string `form:"status"`
	Lampiran   string `form:"lampiran_path"`
}
