package aparatur_desa

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (AparaturDesa, error)
	Save(aparaturDesa AparaturDesa) (AparaturDesa, error)
	Update(aparaturDesa AparaturDesa) (AparaturDesa, error)
	Delete(aparaturDesa AparaturDesa) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var aparaturDesas []AparaturDesa

	err := r.db.Preload(clause.Associations).Scopes(helper.Paginate(aparaturDesas, &pagination, r.db)).Find(&aparaturDesas).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = aparaturDesas
	return pagination, err
}

func (r *repository) FindByID(ID int) (AparaturDesa, error) {
	var aparaturDesa AparaturDesa
	err := r.db.Where("id = ?", ID).Find(&aparaturDesa).Error
	if err != nil {
		return aparaturDesa, err
	}

	return aparaturDesa, nil
}

func (r *repository) Save(aparaturDesa AparaturDesa) (AparaturDesa, error) {
	err := r.db.Create(&aparaturDesa).Error
	if err != nil {
		return aparaturDesa, err
	}

	return aparaturDesa, nil
}

func (r *repository) Update(aparaturDesa AparaturDesa) (AparaturDesa, error) {
	err := r.db.Save(&aparaturDesa).Error
	if err != nil {
		return aparaturDesa, err
	}

	return aparaturDesa, nil
}

func (r *repository) Delete(aparaturDesa AparaturDesa) error {
	err := r.db.Delete(&aparaturDesa).Error
	if err != nil {
		return err
	}

	return nil
}
