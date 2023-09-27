package rt

type GetRTDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateRTInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
