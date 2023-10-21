package kelahiran

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (Kelahiran, error)
	Save(penduduk Kelahiran) (Kelahiran, error)
	Update(penduduk Kelahiran) (Kelahiran, error)
	Delete(penduduk Kelahiran) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var kelahirans []Kelahiran

	err := r.db.Scopes(helper.Paginate(kelahirans, &pagination, r.db)).Find(&kelahirans).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = kelahirans
	return pagination, err
}

func (r *repository) FindByID(ID int) (Kelahiran, error) {
	var kelahirans Kelahiran
	err := r.db.Where("id = ?", ID).First(&kelahirans).Error
	if err != nil {
		return kelahirans, err
	}

	return kelahirans, nil
}

func (r *repository) Save(penduduk Kelahiran) (Kelahiran, error) {
	err := r.db.Create(&penduduk).Error
	if err != nil {
		return penduduk, err
	}

	return penduduk, nil
}

func (r *repository) Update(penduduk Kelahiran) (Kelahiran, error) {
	err := r.db.Save(&penduduk).Error
	if err != nil {
		return penduduk, err
	}

	return penduduk, nil
}

func (r *repository) Delete(penduduk Kelahiran) error {
	err := r.db.Delete(&penduduk).Error
	if err != nil {
		return err
	}

	return nil
}
