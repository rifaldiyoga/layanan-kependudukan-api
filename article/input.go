package article

type GetArticleDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateArticleInput struct {
	Title     string `form:"title" binding:"required"`
	Author    string `form:"author"`
	Content   string `form:"content" binding:"required"`
	ImagePath string `form:"image"`
	Tag       string `form:"tag" binding:"required"`
}
