package article

type GetArticleDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateArticleInput struct {
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Content  string `json:"content" binding:"required"`
	ImageUrl string `json:"image_url"`
	Tag      string `json:"tag" binding:"required"`
}
