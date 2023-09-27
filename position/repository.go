package position

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (Position, error)
	Save(position Position) (Position, error)
	Update(position Position) (Position, error)
	Delete(position Position) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var positions []Position

	err := r.db.Scopes(helper.Paginate(positions, &pagination, r.db)).Find(&positions).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = positions
	return pagination, err
}

func (r *repository) FindByID(ID int) (Position, error) {
	var position Position
	err := r.db.Where("id = ?", ID).First(&position).Error
	if err != nil {
		return position, err
	}

	return position, nil
}

func (r *repository) Save(position Position) (Position, error) {
	err := r.db.Create(&position).Error
	if err != nil {
		return position, err
	}

	return position, nil
}

func (r *repository) Update(position Position) (Position, error) {
	err := r.db.Save(&position).Error
	if err != nil {
		return position, err
	}

	return position, nil
}

func (r *repository) Delete(position Position) error {
	err := r.db.Delete(&position).Error
	if err != nil {
		return err
	}

	return nil
}
