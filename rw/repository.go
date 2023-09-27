package rw

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (RW, error)
	Save(rw RW) (RW, error)
	Update(rw RW) (RW, error)
	Delete(rw RW) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var rws []RW

	err := r.db.Scopes(helper.Paginate(rws, &pagination, r.db)).Find(&rws).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = rws
	return pagination, err
}

func (r *repository) FindByID(ID int) (RW, error) {
	var rw RW
	err := r.db.Where("id = ?", ID).First(&rw).Error
	if err != nil {
		return rw, err
	}

	return rw, nil
}

func (r *repository) Save(rw RW) (RW, error) {
	err := r.db.Create(&rw).Error
	if err != nil {
		return rw, err
	}

	return rw, nil
}

func (r *repository) Update(rw RW) (RW, error) {
	err := r.db.Save(&rw).Error
	if err != nil {
		return rw, err
	}

	return rw, nil
}

func (r *repository) Delete(rw RW) error {
	err := r.db.Delete(&rw).Error
	if err != nil {
		return err
	}

	return nil
}
