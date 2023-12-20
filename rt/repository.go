package rt

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (RT, error)
	Save(rt RT) (RT, error)
	Update(rt RT) (RT, error)
	Delete(rt RT) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var rts []RT

	err := r.db.Preload(clause.Associations).Preload("RW.Kelurahan").Scopes(helper.Paginate(rts, &pagination, r.db)).Find(&rts).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = rts
	return pagination, err
}

func (r *repository) FindByID(ID int) (RT, error) {
	var rt RT
	err := r.db.Where("id = ?", ID).First(&rt).Error
	if err != nil {
		return rt, err
	}

	return rt, nil
}

func (r *repository) Save(rt RT) (RT, error) {
	err := r.db.Create(&rt).Error
	if err != nil {
		return rt, err
	}

	return rt, nil
}

func (r *repository) Update(rt RT) (RT, error) {
	err := r.db.Save(&rt).Error
	if err != nil {
		return rt, err
	}

	return rt, nil
}

func (r *repository) Delete(rt RT) error {
	err := r.db.Delete(&rt).Error
	if err != nil {
		return err
	}

	return nil
}
