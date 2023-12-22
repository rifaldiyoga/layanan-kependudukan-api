package berpergian

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (Berpergian, error)
	FindLast() (Berpergian, error)
	Save(berpergian Berpergian) (Berpergian, error)
	Update(berpergian Berpergian) (Berpergian, error)
	Delete(berpergian Berpergian) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var berpergians []Berpergian
	db := r.db.Debug()

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Order("created_at DESC").Preload(clause.Associations).Preload(clause.Associations).Scopes(helper.Paginate(berpergians, &pagination, r.db)).Where("status = true").Find(&berpergians).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = berpergians
	return pagination, err
}

func (r *repository) FindByID(ID int) (Berpergian, error) {
	var berpergians Berpergian
	err := r.db.Where("id = ?", ID).Preload(clause.Associations).Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi").First(&berpergians).Error
	if err != nil {
		return berpergians, err
	}

	return berpergians, nil
}

func (r *repository) FindLast() (Berpergian, error) {
	var berpergians Berpergian
	err := r.db.Last(&berpergians).Error
	if err != nil {
		return berpergians, err
	}

	return berpergians, nil
}

func (r *repository) Save(berpergian Berpergian) (Berpergian, error) {
	err := r.db.Create(&berpergian).Error
	if err != nil {
		return berpergian, err
	}

	return berpergian, nil
}

func (r *repository) Update(berpergian Berpergian) (Berpergian, error) {
	err := r.db.Save(&berpergian).Error
	if err != nil {
		return berpergian, err
	}

	return berpergian, nil
}

func (r *repository) Delete(berpergian Berpergian) error {
	err := r.db.Delete(&berpergian).Error
	if err != nil {
		return err
	}

	return nil
}
