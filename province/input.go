package province

type GetProvinceDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateProvinceInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
