package religion

type GetReligionDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateReligionInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
