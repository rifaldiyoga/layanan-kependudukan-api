package pengajuan

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/user"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByUser(pagination helper.Pagination, user user.User) (helper.Pagination, error)
	FindByID(id int) (Pengajuan, error)
	Save(pengajuan Pengajuan) (Pengajuan, error)
	Update(pengajuan Pengajuan) (Pengajuan, error)
	Delete(pengajuan Pengajuan) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var pengajuans []Pengajuan

	err := r.db.Scopes(helper.Paginate(pengajuans, &pagination, r.db)).Preload("Detail").Find(&pengajuans).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = pengajuans
	return pagination, err
}

func (r *repository) FindByUser(pagination helper.Pagination, user user.User) (helper.Pagination, error) {
	var pengajuans []Pengajuan

	err := r.db.Scopes(helper.Paginate(pengajuans, &pagination, r.db)).Preload("Detail").Find(&pengajuans).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = pengajuans
	return pagination, err
}

func (r *repository) FindByID(ID int) (Pengajuan, error) {
	var pengajuan Pengajuan
	err := r.db.Debug().Where("id = ?", ID).Preload("Detail").First(&pengajuan).Error
	if err != nil {
		return pengajuan, err
	}

	return pengajuan, nil
}

func (r *repository) Save(pengajuan Pengajuan) (Pengajuan, error) {
	err := r.db.Create(&pengajuan).Error
	if err != nil {
		return pengajuan, err
	}

	return pengajuan, nil
}

func (r *repository) Update(pengajuan Pengajuan) (Pengajuan, error) {
	err := r.db.Save(&pengajuan).Error
	if err != nil {
		return pengajuan, err
	}

	return pengajuan, nil
}

func (r *repository) Delete(pengajuan Pengajuan) error {
	err := r.db.Delete(&pengajuan).Error
	if err != nil {
		return err
	}

	return nil
}