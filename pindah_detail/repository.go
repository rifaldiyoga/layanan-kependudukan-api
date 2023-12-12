package pindah_detail

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (PindahDetail, error)
	FindLast() (PindahDetail, error)
	Save(pindah PindahDetail) (PindahDetail, error)
	Update(pindah PindahDetail) (PindahDetail, error)
	Delete(pindah PindahDetail) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var pindahs []PindahDetail

	err := r.db.Scopes(helper.Paginate(pindahs, &pagination, r.db)).Where("status = true").Find(&pindahs).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = pindahs
	return pagination, err
}

func (r *repository) FindByID(ID int) (PindahDetail, error) {
	var pindahs PindahDetail
	err := r.db.Where("id = ?", ID).First(&pindahs).Error
	if err != nil {
		return pindahs, err
	}

	return pindahs, nil
}

func (r *repository) FindLast() (PindahDetail, error) {
	var pindahs PindahDetail
	err := r.db.Last(&pindahs).Error
	if err != nil {
		return pindahs, err
	}

	return pindahs, nil
}

func (r *repository) Save(pindah PindahDetail) (PindahDetail, error) {
	err := r.db.Create(&pindah).Error
	if err != nil {
		return pindah, err
	}

	return pindah, nil
}

func (r *repository) Update(pindah PindahDetail) (PindahDetail, error) {
	err := r.db.Save(&pindah).Error
	if err != nil {
		return pindah, err
	}

	return pindah, nil
}

func (r *repository) Delete(pindah PindahDetail) error {
	err := r.db.Delete(&pindah).Error
	if err != nil {
		return err
	}

	return nil
}
