package pengajuan

type GetPengajuanDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreatePengajuanInput struct {
	Layanan    string `json:"layanan" binding:"required"`
	Keterangan string `json:"keterangan" binding:"required"`
	LayananID  int    `json:"layanan_id" binding:"required"`
	Status     string `json:"status"`
}
