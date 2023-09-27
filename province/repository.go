package province

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (Province, error)
	Save(province Province) (Province, error)
	Update(province Province) (Province, error)
	Delete(province Province) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var provinces []Province
	err := r.db.Scopes(helper.Paginate(provinces, &pagination, r.db)).Find(&provinces).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = provinces
	return pagination, err
}

func (r *repository) FindByID(ID int) (Province, error) {
	var province Province
	err := r.db.Where("id = ?", ID).First(&province).Error
	if err != nil {
		return province, err
	}

	return province, nil
}

func (r *repository) Save(province Province) (Province, error) {
	err := r.db.Create(&province).Error
	if err != nil {
		return province, err
	}

	return province, nil
}

func (r *repository) Update(province Province) (Province, error) {
	err := r.db.Save(&province).Error
	if err != nil {
		return province, err
	}

	return province, nil
}

func (r *repository) Delete(province Province) error {
	err := r.db.Delete(&province).Error
	if err != nil {
		return err
	}

	return nil
}
