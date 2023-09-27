package detail_pengajuan

type GetDetailPengajuanDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateDetailPengajuanInput struct {
	Keterangan string `json:"keterangan" binding:"required"`
	LayananID  int    `json:"layanan_id" binding:"required"`
	Status     string `json:"status"`
}
