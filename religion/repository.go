package religion

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (Religion, error)
	Save(religion Religion) (Religion, error)
	Update(religion Religion) (Religion, error)
	Delete(religion Religion) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var religions []Religion

	err := r.db.Scopes(helper.Paginate(religions, &pagination, r.db)).Find(&religions).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = religions
	return pagination, err
}

func (r *repository) FindByID(ID int) (Religion, error) {
	var religions Religion
	err := r.db.Where("id = ?", ID).First(&religions).Error
	if err != nil {
		return religions, err
	}

	return religions, nil
}

func (r *repository) Save(religion Religion) (Religion, error) {
	err := r.db.Create(&religion).Error
	if err != nil {
		return religion, err
	}

	return religion, nil
}

func (r *repository) Update(religion Religion) (Religion, error) {
	err := r.db.Save(&religion).Error
	if err != nil {
		return religion, err
	}

	return religion, nil
}

func (r *repository) Delete(religion Religion) error {
	err := r.db.Delete(&religion).Error
	if err != nil {
		return err
	}

	return nil
}
