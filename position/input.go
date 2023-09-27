package position

type GetPositionDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreatePositionInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
