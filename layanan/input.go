package layanan

type GetLayananDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateLayananInput struct {
	Code      string `json:"code" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Type      string `json:"type" binding:"required"`
	IsConfirm bool   `json:"is_confirm" binding:"required"`
}
