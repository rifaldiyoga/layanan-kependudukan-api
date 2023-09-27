package job

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (Job, error)
	Save(job Job) (Job, error)
	Update(job Job) (Job, error)
	Delete(job Job) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var jobs []Job

	err := r.db.Scopes(helper.Paginate(jobs, &pagination, r.db)).Find(&jobs).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = jobs
	return pagination, err
}

func (r *repository) FindByID(ID int) (Job, error) {
	var job Job
	err := r.db.Where("id = ?", ID).First(&job).Error
	if err != nil {
		return job, err
	}

	return job, nil
}

func (r *repository) Save(job Job) (Job, error) {
	err := r.db.Create(&job).Error
	if err != nil {
		return job, err
	}

	return job, nil
}

func (r *repository) Update(job Job) (Job, error) {
	err := r.db.Save(&job).Error
	if err != nil {
		return job, err
	}

	return job, nil
}

func (r *repository) Delete(job Job) error {
	err := r.db.Delete(&job).Error
	if err != nil {
		return err
	}

	return nil
}
