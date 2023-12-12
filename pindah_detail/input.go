package pindah_detail

type GetPindahDetailDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreatePindahDetailInput struct {
	NIK          string `json:"nik"`
	Nama         string `json:"fullname"`
	StatusFamily string `json:"status_family"`
	PindahID     int    `json:"pindah_id"`
}
