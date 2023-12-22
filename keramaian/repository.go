package keramaian

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (Keramaian, error)
	FindLast() (Keramaian, error)
	Save(keramaian Keramaian) (Keramaian, error)
	Update(keramaian Keramaian) (Keramaian, error)
	Delete(keramaian Keramaian) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var keramaians []Keramaian

	db := r.db.Debug()

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Order("created_at DESC").Preload(clause.Associations).Scopes(helper.Paginate(keramaians, &pagination, r.db)).Where("status = true").Find(&keramaians).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = keramaians
	return pagination, err
}

func (r *repository) FindByID(ID int) (Keramaian, error) {
	var keramaians Keramaian
	err := r.db.Where("id = ?", ID).Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi").First(&keramaians).Error
	if err != nil {
		return keramaians, err
	}

	return keramaians, nil
}

func (r *repository) FindLast() (Keramaian, error) {
	var keramaians Keramaian
	err := r.db.Last(&keramaians).Error
	if err != nil {
		return keramaians, err
	}

	return keramaians, nil
}

func (r *repository) Save(keramaian Keramaian) (Keramaian, error) {
	err := r.db.Create(&keramaian).Error
	if err != nil {
		return keramaian, err
	}

	return keramaian, nil
}

func (r *repository) Update(keramaian Keramaian) (Keramaian, error) {
	err := r.db.Save(&keramaian).Error
	if err != nil {
		return keramaian, err
	}

	return keramaian, nil
}

func (r *repository) Delete(keramaian Keramaian) error {
	err := r.db.Delete(&keramaian).Error
	if err != nil {
		return err
	}

	return nil
}
