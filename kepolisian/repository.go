package kepolisian

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (Kepolisian, error)
	FindLast() (Kepolisian, error)
	Save(kepolisian Kepolisian) (Kepolisian, error)
	Update(kepolisian Kepolisian) (Kepolisian, error)
	Delete(kepolisian Kepolisian) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var kepolisians []Kepolisian

	db := r.db.Debug()

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Scopes(helper.Paginate(kepolisians, &pagination, r.db)).Where("status = true").Find(&kepolisians).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = kepolisians
	return pagination, err
}

func (r *repository) FindByID(ID int) (Kepolisian, error) {
	var kepolisians Kepolisian
	err := r.db.Where("id = ?", ID).Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi").First(&kepolisians).Error
	if err != nil {
		return kepolisians, err
	}

	return kepolisians, nil
}

func (r *repository) FindLast() (Kepolisian, error) {
	var kepolisians Kepolisian
	err := r.db.Last(&kepolisians).Error
	if err != nil {
		return kepolisians, err
	}

	return kepolisians, nil
}

func (r *repository) Save(kepolisian Kepolisian) (Kepolisian, error) {
	err := r.db.Create(&kepolisian).Error
	if err != nil {
		return kepolisian, err
	}

	return kepolisian, nil
}

func (r *repository) Update(kepolisian Kepolisian) (Kepolisian, error) {
	err := r.db.Save(&kepolisian).Error
	if err != nil {
		return kepolisian, err
	}

	return kepolisian, nil
}

func (r *repository) Delete(kepolisian Kepolisian) error {
	err := r.db.Delete(&kepolisian).Error
	if err != nil {
		return err
	}

	return nil
}
