package tanah

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (Tanah, error)
	FindLast() (Tanah, error)
	Save(tanah Tanah) (Tanah, error)
	Update(tanah Tanah) (Tanah, error)
	Delete(tanah Tanah) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var tanahs []Tanah

	db := r.db.Debug()

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Order("created_at DESC").Preload(clause.Associations).Scopes(helper.Paginate(tanahs, &pagination, r.db)).Where("status = true").Find(&tanahs).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = tanahs
	return pagination, err
}

func (r *repository) FindByID(ID int) (Tanah, error) {
	var tanahs Tanah
	err := r.db.Where("id = ?", ID).Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi").First(&tanahs).Error
	if err != nil {
		return tanahs, err
	}

	return tanahs, nil
}

func (r *repository) FindLast() (Tanah, error) {
	var tanahs Tanah
	err := r.db.Last(&tanahs).Error
	if err != nil {
		return tanahs, err
	}

	return tanahs, nil
}

func (r *repository) Save(tanah Tanah) (Tanah, error) {
	err := r.db.Create(&tanah).Error
	if err != nil {
		return tanah, err
	}

	return tanah, nil
}

func (r *repository) Update(tanah Tanah) (Tanah, error) {
	err := r.db.Save(&tanah).Error
	if err != nil {
		return tanah, err
	}

	return tanah, nil
}

func (r *repository) Delete(tanah Tanah) error {
	err := r.db.Delete(&tanah).Error
	if err != nil {
		return err
	}

	return nil
}
