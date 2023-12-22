package janda

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (Janda, error)
	FindLast() (Janda, error)
	Save(janda Janda) (Janda, error)
	Update(janda Janda) (Janda, error)
	Delete(janda Janda) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var jandas []Janda

	db := r.db.Debug()

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Order("created_at DESC").Preload(clause.Associations).Scopes(helper.Paginate(jandas, &pagination, r.db)).Where("status = true").Find(&jandas).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = jandas
	return pagination, err
}

func (r *repository) FindByID(ID int) (Janda, error) {
	var jandas Janda
	err := r.db.Where("id = ?", ID).Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi").First(&jandas).Error
	if err != nil {
		return jandas, err
	}

	return jandas, nil
}

func (r *repository) FindLast() (Janda, error) {
	var jandas Janda
	err := r.db.Last(&jandas).Error
	if err != nil {
		return jandas, err
	}

	return jandas, nil
}

func (r *repository) Save(janda Janda) (Janda, error) {
	err := r.db.Create(&janda).Error
	if err != nil {
		return janda, err
	}

	return janda, nil
}

func (r *repository) Update(janda Janda) (Janda, error) {
	err := r.db.Save(&janda).Error
	if err != nil {
		return janda, err
	}

	return janda, nil
}

func (r *repository) Delete(janda Janda) error {
	err := r.db.Delete(&janda).Error
	if err != nil {
		return err
	}

	return nil
}
