package aparatur_desa

type GetAparaturDesaDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateAparaturDesaInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
