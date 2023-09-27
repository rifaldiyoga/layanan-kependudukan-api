package education

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (Education, error)
	Save(education Education) (Education, error)
	Update(education Education) (Education, error)
	Delete(education Education) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var educationss []Education

	err := r.db.Scopes(helper.Paginate(educationss, &pagination, r.db)).Find(&educationss).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = educationss
	return pagination, err
}

func (r *repository) FindByID(ID int) (Education, error) {
	var educations Education
	err := r.db.Where("id = ?", ID).First(&educations).Error
	if err != nil {
		return educations, err
	}

	return educations, nil
}

func (r *repository) Save(education Education) (Education, error) {
	err := r.db.Create(&education).Error
	if err != nil {
		return education, err
	}

	return education, nil
}

func (r *repository) Update(Education Education) (Education, error) {
	err := r.db.Save(&Education).Error
	if err != nil {
		return Education, err
	}

	return Education, nil
}

func (r *repository) Delete(education Education) error {
	err := r.db.Delete(&education).Error
	if err != nil {
		return err
	}

	return nil
}
