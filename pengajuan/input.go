package pengajuan

type GetPengajuanDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreatePengajuanInput struct {
	Layanan    string `json:"layanan" binding:"required"`
	Code       string `json:"code" `
	Keterangan string `json:"keterangan" binding:"required"`
	LayananID  int    `json:"layanan_id" binding:"required"`
	Status     string `json:"status"`
	RefID      int    `json:"ref_id" binding:"required"`
	Note       string `json:"note" `
}

type UpdateStatusPengajuanInput struct {
	PengajuanID string `json:"pengajuan_id" binding:"required"`
	Status      string `json:"status" binding:"required"`
	Note        string `json:"note" `
}
