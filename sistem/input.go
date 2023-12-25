package sistem

type GetSistemDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateSistemInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
