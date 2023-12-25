package position

type GetPositionDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreatePositionInput struct {
	Code    string `json:"code" binding:"required"`
	Jabatan string `json:"jabatan" binding:"required"`
}
