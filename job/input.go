package job

type GetJobDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateJobInput struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}
