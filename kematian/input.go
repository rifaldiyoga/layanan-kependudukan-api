package kematian

type GetKematianDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateKematianInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
