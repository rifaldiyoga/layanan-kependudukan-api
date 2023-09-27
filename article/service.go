package article

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetArticleByID(ID int) (Article, error)
	GetArticles(pagination helper.Pagination) (helper.Pagination, error)
	CreateArticle(input CreateArticleInput) (Article, error)
	UpdateArticle(ID GetArticleDetailInput, input CreateArticleInput) (Article, error)
	DeleteArticle(ID GetArticleDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetArticleByID(ID int) (Article, error) {
	Article, err := s.repository.FindByID(ID)
	if err != nil {
		return Article, err
	}

	return Article, nil
}

func (s *service) CreateArticle(input CreateArticleInput) (Article, error) {
	Article := Article{}

	Article.Title = input.Title
	Article.Content = input.Content
	Article.Tag = input.Tag
	Article.ImageUrl = input.ImageUrl
	Article.Author = input.Author
	Article.CreatedAt = time.Now()

	newArticle, err := s.repository.Save(Article)
	return newArticle, err
}

func (s *service) UpdateArticle(inputDetail GetArticleDetailInput, input CreateArticleInput) (Article, error) {
	Article, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return Article, err
	}

	Article.Title = input.Title
	Article.Content = input.Content
	Article.Tag = input.Tag
	Article.ImageUrl = input.ImageUrl
	Article.Author = input.Author

	newArticle, err := s.repository.Update(Article)
	return newArticle, err
}

func (s *service) DeleteArticle(inputDetail GetArticleDetailInput) error {
	Article, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(Article)
	return err
}

func (s *service) GetArticles(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)
	return pagination, err
}
