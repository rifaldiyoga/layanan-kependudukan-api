package sktm

type GetSKTMDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateSKTMInput struct {
	KodeSurat  string `json:"kode_surat"`
	Keterangan string `json:"keterangan" binding:"required"`
	NIK        string `json:"nik" `
	Status     string `json:"status"`
}
