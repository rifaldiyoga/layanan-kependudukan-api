package pindah

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (Pindah, error)
	FindLast() (Pindah, error)
	Save(pindah Pindah) (Pindah, error)
	Update(pindah Pindah) (Pindah, error)
	Delete(pindah Pindah) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var pindahs []Pindah

	db := r.db.Debug()

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Order("created_at DESC").Preload(clause.Associations).Scopes(helper.Paginate(pindahs, &pagination, r.db)).Where("status = true").Find(&pindahs).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = pindahs
	return pagination, err
}

func (r *repository) FindByID(ID int) (Pindah, error) {
	var pindahs Pindah
	err := r.db.Where("id = ?", ID).Preload(clause.Associations).Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi").First(&pindahs).Error
	if err != nil {
		return pindahs, err
	}

	return pindahs, nil
}

func (r *repository) FindLast() (Pindah, error) {
	var pindahs Pindah
	err := r.db.Last(&pindahs).Error
	if err != nil {
		return pindahs, err
	}

	return pindahs, nil
}

func (r *repository) Save(pindah Pindah) (Pindah, error) {
	err := r.db.Create(&pindah).Error
	if err != nil {
		return pindah, err
	}

	return pindah, nil
}

func (r *repository) Update(pindah Pindah) (Pindah, error) {
	err := r.db.Save(&pindah).Error
	if err != nil {
		return pindah, err
	}

	return pindah, nil
}

func (r *repository) Delete(pindah Pindah) error {
	err := r.db.Delete(&pindah).Error
	if err != nil {
		return err
	}

	return nil
}
