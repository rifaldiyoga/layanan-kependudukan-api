package district

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, provinceId int) (helper.Pagination, error)
	FindByID(id int) (District, error)
	Save(district District) (District, error)
	Update(district District) (District, error)
	Delete(district District) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, provinceId int) (helper.Pagination, error) {
	var districts []District

	db := r.db
	if provinceId > 0 {
		db = db.Where("provinsi_id = ?", provinceId)
	}
	err := db.Preload(clause.Associations).Scopes(helper.Paginate(districts, &pagination, r.db)).Find(&districts).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = districts
	return pagination, err
}

func (r *repository) FindByID(ID int) (District, error) {
	var district District
	err := r.db.Where("id = ?", ID).First(&district).Error
	if err != nil {
		return district, err
	}

	return district, nil
}

func (r *repository) Save(district District) (District, error) {
	err := r.db.Create(&district).Error
	if err != nil {
		return district, err
	}

	return district, nil
}

func (r *repository) Update(district District) (District, error) {
	err := r.db.Save(&district).Error
	if err != nil {
		return district, err
	}

	return district, nil
}

func (r *repository) Delete(district District) error {
	err := r.db.Delete(&district).Error
	if err != nil {
		return err
	}

	return nil
}
