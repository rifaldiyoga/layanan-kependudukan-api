package pernah_menikah

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (PernahMenikah, error)
	FindLast() (PernahMenikah, error)
	Save(janda PernahMenikah) (PernahMenikah, error)
	Update(janda PernahMenikah) (PernahMenikah, error)
	Delete(janda PernahMenikah) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var jandas []PernahMenikah

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

func (r *repository) FindByID(ID int) (PernahMenikah, error) {
	var jandas PernahMenikah
	db := r.db.Where("id = ?", ID).Preload(clause.Associations)
	db = db.Preload("Suami.Religion").Preload("Suami.Job").Preload("Suami.Education").Preload("Suami.RT").Preload("Suami.RW").Preload("Suami.Kelurahan").Preload("Suami.Kecamatan").Preload("Suami.Kota").Preload("Suami.Provinsi")
	db = db.Preload("Istri.Religion").Preload("Istri.Job").Preload("Istri.Education").Preload("Istri.RT").Preload("Istri.RW").Preload("Istri.Kelurahan").Preload("Istri.Kecamatan").Preload("Istri.Kota").Preload("Istri.Provinsi")
	db = db.Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi")
	err := db.First(&jandas).Error
	if err != nil {
		return jandas, err
	}

	return jandas, nil
}

func (r *repository) FindLast() (PernahMenikah, error) {
	var jandas PernahMenikah
	err := r.db.Last(&jandas).Error
	if err != nil {
		return jandas, err
	}

	return jandas, nil
}

func (r *repository) Save(janda PernahMenikah) (PernahMenikah, error) {
	err := r.db.Create(&janda).Error
	if err != nil {
		return janda, err
	}

	return janda, nil
}

func (r *repository) Update(janda PernahMenikah) (PernahMenikah, error) {
	err := r.db.Save(&janda).Error
	if err != nil {
		return janda, err
	}

	return janda, nil
}

func (r *repository) Delete(janda PernahMenikah) error {
	err := r.db.Delete(&janda).Error
	if err != nil {
		return err
	}

	return nil
}
