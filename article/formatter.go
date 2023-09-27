package article

import "layanan-kependudukan-api/helper"

type ArticleFormatter struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Content   string `json:"content"`
	ImageUrl  string `json:"image_url"`
	Tag       string `json:"tag"`
	CreatedAt string `json:"created_at"`
}

func FormatArticle(article Article) ArticleFormatter {
	formatter := ArticleFormatter{
		ID:        article.ID,
		Title:     article.Title,
		Author:    article.Author,
		Content:   article.Content,
		ImageUrl:  article.ImageUrl,
		Tag:       article.Tag,
		CreatedAt: helper.FormatDateToString(article.CreatedAt),
	}

	return formatter
}

func FormatArticles(Articles []Article) []ArticleFormatter {
	var ArticlesFormatter []ArticleFormatter

	for _, Article := range Articles {
		ArticleFormatter := FormatArticle(Article)
		ArticlesFormatter = append(ArticlesFormatter, ArticleFormatter)
	}

	return ArticlesFormatter
}
