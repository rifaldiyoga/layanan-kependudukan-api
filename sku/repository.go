package sku

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (SKU, error)
	FindLast() (SKU, error)
	Save(sku SKU) (SKU, error)
	Update(sku SKU) (SKU, error)
	Delete(sku SKU) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var skus []SKU

	db := r.db.Debug()

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Scopes(helper.Paginate(skus, &pagination, r.db)).Where("status = true").Find(&skus).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = skus
	return pagination, err
}

func (r *repository) FindByID(ID int) (SKU, error) {
	var skus SKU
	err := r.db.Where("id = ?", ID).Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi").First(&skus).Error
	if err != nil {
		return skus, err
	}

	return skus, nil
}

func (r *repository) FindLast() (SKU, error) {
	var skus SKU
	err := r.db.Last(&skus).Error
	if err != nil {
		return skus, err
	}

	return skus, nil
}

func (r *repository) Save(sku SKU) (SKU, error) {
	err := r.db.Create(&sku).Error
	if err != nil {
		return sku, err
	}

	return sku, nil
}

func (r *repository) Update(sku SKU) (SKU, error) {
	err := r.db.Save(&sku).Error
	if err != nil {
		return sku, err
	}

	return sku, nil
}

func (r *repository) Delete(sku SKU) error {
	err := r.db.Delete(&sku).Error
	if err != nil {
		return err
	}

	return nil
}
