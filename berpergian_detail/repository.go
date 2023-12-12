package berpergian_detail

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (BerpergianDetail, error)
	FindLast() (BerpergianDetail, error)
	Save(berpergian BerpergianDetail) (BerpergianDetail, error)
	Update(berpergian BerpergianDetail) (BerpergianDetail, error)
	Delete(berpergian BerpergianDetail) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var berpergians []BerpergianDetail

	err := r.db.Scopes(helper.Paginate(berpergians, &pagination, r.db)).Where("status = true").Find(&berpergians).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = berpergians
	return pagination, err
}

func (r *repository) FindByID(ID int) (BerpergianDetail, error) {
	var berpergians BerpergianDetail
	err := r.db.Where("id = ?", ID).First(&berpergians).Error
	if err != nil {
		return berpergians, err
	}

	return berpergians, nil
}

func (r *repository) FindLast() (BerpergianDetail, error) {
	var berpergians BerpergianDetail
	err := r.db.Last(&berpergians).Error
	if err != nil {
		return berpergians, err
	}

	return berpergians, nil
}

func (r *repository) Save(berpergian BerpergianDetail) (BerpergianDetail, error) {
	err := r.db.Create(&berpergian).Error
	if err != nil {
		return berpergian, err
	}

	return berpergian, nil
}

func (r *repository) Update(berpergian BerpergianDetail) (BerpergianDetail, error) {
	err := r.db.Save(&berpergian).Error
	if err != nil {
		return berpergian, err
	}

	return berpergian, nil
}

func (r *repository) Delete(berpergian BerpergianDetail) error {
	err := r.db.Delete(&berpergian).Error
	if err != nil {
		return err
	}

	return nil
}
