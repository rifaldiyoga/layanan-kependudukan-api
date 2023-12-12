package kematian

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (Kematian, error)
	FindLast() (Kematian, error)
	Save(kematian Kematian) (Kematian, error)
	Update(kematian Kematian) (Kematian, error)
	Delete(kematian Kematian) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var kematians []Kematian

	db := r.db.Debug()

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Scopes(helper.Paginate(kematians, &pagination, r.db)).Where("status = true").Find(&kematians).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = kematians
	return pagination, err
}

func (r *repository) FindByID(ID int) (Kematian, error) {
	var kematians Kematian
	db := r.db.Where("id = ?", ID).Preload(clause.Associations)
	db = db.Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi")
	db = db.Preload("Jenazah.Religion").Preload("Jenazah.Job").Preload("Jenazah.Education").Preload("Jenazah.RT").Preload("Jenazah.RW").Preload("Jenazah.Kelurahan").Preload("Jenazah.Kecamatan").Preload("Jenazah.Kota").Preload("Jenazah.Provinsi")
	err := db.First(&kematians).Error
	if err != nil {
		return kematians, err
	}

	return kematians, nil
}

func (r *repository) FindLast() (Kematian, error) {
	var kematians Kematian
	err := r.db.Last(&kematians).Error
	if err != nil {
		return kematians, err
	}

	return kematians, nil
}

func (r *repository) Save(kematian Kematian) (Kematian, error) {
	err := r.db.Create(&kematian).Error
	if err != nil {
		return kematian, err
	}

	return kematian, nil
}

func (r *repository) Update(kematian Kematian) (Kematian, error) {
	err := r.db.Save(&kematian).Error
	if err != nil {
		return kematian, err
	}

	return kematian, nil
}

func (r *repository) Delete(kematian Kematian) error {
	err := r.db.Delete(&kematian).Error
	if err != nil {
		return err
	}

	return nil
}
