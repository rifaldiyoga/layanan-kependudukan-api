package kelurahan

type GetKelurahanDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateKelurahanInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
