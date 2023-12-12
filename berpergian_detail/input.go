package berpergian_detail

type GetBerpergianDetailDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateBerpergianDetailInput struct {
	NIK          string `json:"nik"`
	Nama         string `json:"fullname"`
	StatusFamily string `json:"status_family"`
	BerpergianID int    `json:"berpergian_id"`
}
