package kematian

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (Kematian, error)
	Save(penduduk Kematian) (Kematian, error)
	Update(penduduk Kematian) (Kematian, error)
	Delete(penduduk Kematian) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var kematians []Kematian

	err := r.db.Scopes(helper.Paginate(kematians, &pagination, r.db)).Find(&kematians).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = kematians
	return pagination, err
}

func (r *repository) FindByID(ID int) (Kematian, error) {
	var kematians Kematian
	err := r.db.Where("id = ?", ID).First(&kematians).Error
	if err != nil {
		return kematians, err
	}

	return kematians, nil
}

func (r *repository) Save(penduduk Kematian) (Kematian, error) {
	err := r.db.Create(&penduduk).Error
	if err != nil {
		return penduduk, err
	}

	return penduduk, nil
}

func (r *repository) Update(penduduk Kematian) (Kematian, error) {
	err := r.db.Save(&penduduk).Error
	if err != nil {
		return penduduk, err
	}

	return penduduk, nil
}

func (r *repository) Delete(penduduk Kematian) error {
	err := r.db.Delete(&penduduk).Error
	if err != nil {
		return err
	}

	return nil
}
