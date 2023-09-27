package district

type GetDistrictDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateDistrictInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
