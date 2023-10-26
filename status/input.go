package status

type GetStatusDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateStatusInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
