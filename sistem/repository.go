package sistem

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (Sistem, error)
	Save(sistem Sistem) (Sistem, error)
	Update(sistem Sistem) (Sistem, error)
	Delete(sistem Sistem) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var sistems []Sistem

	err := r.db.Scopes(helper.Paginate(sistems, &pagination, r.db)).Find(&sistems).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = sistems
	return pagination, err
}

func (r *repository) FindByID(ID int) (Sistem, error) {
	var sistems Sistem
	err := r.db.Preload(clause.Associations).Where("id = ?", ID).First(&sistems).Error
	if err != nil {
		return sistems, err
	}

	return sistems, nil
}

func (r *repository) Save(sistem Sistem) (Sistem, error) {
	err := r.db.Create(&sistem).Error
	if err != nil {
		return sistem, err
	}

	return sistem, nil
}

func (r *repository) Update(sistem Sistem) (Sistem, error) {
	err := r.db.Save(&sistem).Error
	if err != nil {
		return sistem, err
	}

	return sistem, nil
}

func (r *repository) Delete(sistem Sistem) error {
	err := r.db.Delete(&sistem).Error
	if err != nil {
		return err
	}

	return nil
}
