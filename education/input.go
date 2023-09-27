package education

type GetEducationDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateEducationInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
