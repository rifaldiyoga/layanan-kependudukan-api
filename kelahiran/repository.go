package kelahiran

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (Kelahiran, error)
	FindLast() (Kelahiran, error)
	Save(kelahiran Kelahiran) (Kelahiran, error)
	Update(kelahiran Kelahiran) (Kelahiran, error)
	Delete(kelahiran Kelahiran) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var kelahirans []Kelahiran

	db := r.db.Debug()

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Order("created_at DESC").Preload(clause.Associations).Scopes(helper.Paginate(kelahirans, &pagination, r.db)).Where("status = true").Find(&kelahirans).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = kelahirans
	return pagination, err
}

func (r *repository) FindByID(ID int) (Kelahiran, error) {
	var kelahirans Kelahiran
	db := r.db.Where("id = ?", ID).Preload(clause.Associations)
	db = db.Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi")
	db = db.Preload("Ayah.Religion").Preload("Ayah.Job").Preload("Ayah.Education").Preload("Ayah.RT").Preload("Ayah.RW").Preload("Ayah.Kelurahan").Preload("Ayah.Kecamatan").Preload("Ayah.Kota").Preload("Ayah.Provinsi")
	db = db.Preload("Ibu.Religion").Preload("Ibu.Job").Preload("Ibu.Education").Preload("Ibu.RT").Preload("Ibu.RW").Preload("Ibu.Kelurahan").Preload("Ibu.Kecamatan").Preload("Ibu.Kota").Preload("Ibu.Provinsi")
	err := db.First(&kelahirans).Error
	if err != nil {
		return kelahirans, err
	}

	return kelahirans, nil
}

func (r *repository) FindLast() (Kelahiran, error) {
	var kelahirans Kelahiran
	err := r.db.Last(&kelahirans).Error
	if err != nil {
		return kelahirans, err
	}

	return kelahirans, nil
}

func (r *repository) Save(kelahiran Kelahiran) (Kelahiran, error) {
	err := r.db.Create(&kelahiran).Error
	if err != nil {
		return kelahiran, err
	}

	return kelahiran, nil
}

func (r *repository) Update(kelahiran Kelahiran) (Kelahiran, error) {
	err := r.db.Save(&kelahiran).Error
	if err != nil {
		return kelahiran, err
	}

	return kelahiran, nil
}

func (r *repository) Delete(kelahiran Kelahiran) error {
	err := r.db.Delete(&kelahiran).Error
	if err != nil {
		return err
	}

	return nil
}
