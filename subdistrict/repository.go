package subdistrict

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, distritId int) (helper.Pagination, error)
	FindByID(id int) (SubDistrict, error)
	Save(subDistrict SubDistrict) (SubDistrict, error)
	Update(subDistrict SubDistrict) (SubDistrict, error)
	Delete(subDistrict SubDistrict) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, districtId int) (helper.Pagination, error) {
	var subDistricts []SubDistrict

	db := r.db
	if districtId > 0 {
		db = db.Where("kota_id = ?", districtId)
	}
	err := db.Preload(clause.Associations).Scopes(helper.Paginate(subDistricts, &pagination, r.db)).Find(&subDistricts).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = subDistricts
	return pagination, err
}

func (r *repository) FindByID(ID int) (SubDistrict, error) {
	var subDistrict SubDistrict
	err := r.db.Where("id = ?", ID).First(&subDistrict).Error
	if err != nil {
		return subDistrict, err
	}

	return subDistrict, nil
}

func (r *repository) Save(subDistrict SubDistrict) (SubDistrict, error) {
	err := r.db.Create(&subDistrict).Error
	if err != nil {
		return subDistrict, err
	}

	return subDistrict, nil
}

func (r *repository) Update(subDistrict SubDistrict) (SubDistrict, error) {
	err := r.db.Save(&subDistrict).Error
	if err != nil {
		return subDistrict, err
	}

	return subDistrict, nil
}

func (r *repository) Delete(subDistrict SubDistrict) error {
	err := r.db.Delete(&subDistrict).Error
	if err != nil {
		return err
	}

	return nil
}
