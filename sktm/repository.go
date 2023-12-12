package sktm

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (SKTM, error)
	FindLast() (SKTM, error)
	Save(sktm SKTM) (SKTM, error)
	Update(sktm SKTM) (SKTM, error)
	Delete(sktm SKTM) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var sktms []SKTM

	db := r.db.Debug()

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Scopes(helper.Paginate(sktms, &pagination, r.db)).Where("status = true").Find(&sktms).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = sktms
	return pagination, err
}

func (r *repository) FindByID(ID int) (SKTM, error) {
	var sktms SKTM
	err := r.db.Where("id = ?", ID).Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi").First(&sktms).Error
	if err != nil {
		return sktms, err
	}

	return sktms, nil
}

func (r *repository) FindLast() (SKTM, error) {
	var sktms SKTM
	err := r.db.Last(&sktms).Error
	if err != nil {
		return sktms, err
	}

	return sktms, nil
}

func (r *repository) Save(sktm SKTM) (SKTM, error) {
	err := r.db.Create(&sktm).Error
	if err != nil {
		return sktm, err
	}

	return sktm, nil
}

func (r *repository) Update(sktm SKTM) (SKTM, error) {
	err := r.db.Save(&sktm).Error
	if err != nil {
		return sktm, err
	}

	return sktm, nil
}

func (r *repository) Delete(sktm SKTM) error {
	err := r.db.Delete(&sktm).Error
	if err != nil {
		return err
	}

	return nil
}
