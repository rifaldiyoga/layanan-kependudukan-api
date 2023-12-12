package sku

type GetSKUDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateSKUInput struct {
	KodeSurat  string `json:"kode_surat"`
	Usaha      string `json:"usaha" binding:"required"`
	Keterangan string `json:"keterangan" binding:"required"`
	NIK        string `json:"nik" `
	Status     string `json:"status"`
}
