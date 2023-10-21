package kelahiran

type GetKelahiranDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateKelahiranInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
