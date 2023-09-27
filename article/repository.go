package article

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (Article, error)
	Save(article Article) (Article, error)
	Update(article Article) (Article, error)
	Delete(article Article) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var articles []Article

	err := r.db.Scopes(helper.Paginate(articles, &pagination, r.db)).Find(&articles).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = articles
	return pagination, err
}

func (r *repository) FindByID(ID int) (Article, error) {
	var article Article
	err := r.db.Where("id = ?", ID).First(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

func (r *repository) Save(article Article) (Article, error) {
	err := r.db.Create(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

func (r *repository) Update(article Article) (Article, error) {
	err := r.db.Save(&article).Error
	if err != nil {
		return article, err
	}

	return article, nil
}

func (r *repository) Delete(article Article) error {
	err := r.db.Delete(&article).Error
	if err != nil {
		return err
	}

	return nil
}
