package rw

type GetRWDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateRWInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
