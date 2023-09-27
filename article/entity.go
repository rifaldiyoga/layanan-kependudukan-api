package article

import "time"

type Article struct {
	ID        int
	Title     string
	Author    string
	Content   string
	ImageUrl  string
	Tag       string
	CreatedAt time.Time
}

func (Article) TableName() string {
	return "tb_article"
}
