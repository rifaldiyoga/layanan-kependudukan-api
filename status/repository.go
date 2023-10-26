package status

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (Status, error)
	Save(status Status) (Status, error)
	Update(status Status) (Status, error)
	Delete(status Status) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var statuss []Status

	err := r.db.Scopes(helper.Paginate(statuss, &pagination, r.db)).Find(&statuss).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = statuss
	return pagination, err
}

func (r *repository) FindByID(ID int) (Status, error) {
	var statuss Status
	err := r.db.Where("id = ?", ID).First(&statuss).Error
	if err != nil {
		return statuss, err
	}

	return statuss, nil
}

func (r *repository) Save(status Status) (Status, error) {
	err := r.db.Create(&status).Error
	if err != nil {
		return status, err
	}

	return status, nil
}

func (r *repository) Update(status Status) (Status, error) {
	err := r.db.Save(&status).Error
	if err != nil {
		return status, err
	}

	return status, nil
}

func (r *repository) Delete(status Status) error {
	err := r.db.Delete(&status).Error
	if err != nil {
		return err
	}

	return nil
}
