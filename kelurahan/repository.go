package kelurahan

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (Kelurahan, error)
	Save(kelurahan Kelurahan) (Kelurahan, error)
	Update(kelurahan Kelurahan) (Kelurahan, error)
	Delete(kelurahan Kelurahan) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var kelurahans []Kelurahan

	err := r.db.Scopes(helper.Paginate(kelurahans, &pagination, r.db)).Find(&kelurahans).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = kelurahans
	return pagination, err
}

func (r *repository) FindByID(ID int) (Kelurahan, error) {
	var kelurahan Kelurahan
	err := r.db.Where("id = ?", ID).Find(&kelurahan).Error
	if err != nil {
		return kelurahan, err
	}

	return kelurahan, nil
}

func (r *repository) Save(kelurahan Kelurahan) (Kelurahan, error) {
	err := r.db.Create(&kelurahan).Error
	if err != nil {
		return kelurahan, err
	}

	return kelurahan, nil
}

func (r *repository) Update(kelurahan Kelurahan) (Kelurahan, error) {
	err := r.db.Save(&kelurahan).Error
	if err != nil {
		return kelurahan, err
	}

	return kelurahan, nil
}

func (r *repository) Delete(kelurahan Kelurahan) error {
	err := r.db.Delete(&kelurahan).Error
	if err != nil {
		return err
	}

	return nil
}
