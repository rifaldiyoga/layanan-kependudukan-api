package keluarga

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (Keluarga, error)
	Save(penduduk Keluarga) (Keluarga, error)
	Update(penduduk Keluarga) (Keluarga, error)
	Delete(penduduk Keluarga) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var keluargas []Keluarga

	err := r.db.Scopes(helper.Paginate(keluargas, &pagination, r.db)).Find(&keluargas).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = keluargas
	return pagination, err
}

func (r *repository) FindByID(ID int) (Keluarga, error) {
	var keluargas Keluarga
	err := r.db.Where("id = ?", ID).First(&keluargas).Error
	if err != nil {
		return keluargas, err
	}

	return keluargas, nil
}

func (r *repository) Save(penduduk Keluarga) (Keluarga, error) {
	err := r.db.Create(&penduduk).Error
	if err != nil {
		return penduduk, err
	}

	return penduduk, nil
}

func (r *repository) Update(penduduk Keluarga) (Keluarga, error) {
	err := r.db.Save(&penduduk).Error
	if err != nil {
		return penduduk, err
	}

	return penduduk, nil
}

func (r *repository) Delete(penduduk Keluarga) error {
	err := r.db.Delete(&penduduk).Error
	if err != nil {
		return err
	}

	return nil
}
