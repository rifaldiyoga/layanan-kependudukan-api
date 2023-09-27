package subdistrict

type GetSubDistrictDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateSubDistrictInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
