package domisili

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (Domisili, error)
	FindLast() (Domisili, error)
	Save(domisili Domisili) (Domisili, error)
	Update(domisili Domisili) (Domisili, error)
	Delete(domisili Domisili) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var domisilis []Domisili

	db := r.db.Debug()

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Order("created_at DESC").Preload(clause.Associations).Scopes(helper.Paginate(domisilis, &pagination, r.db)).Where("status = true").Find(&domisilis).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = domisilis
	return pagination, err
}

func (r *repository) FindByID(ID int) (Domisili, error) {
	var domisilis Domisili
	err := r.db.Where("id = ?", ID).Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi").First(&domisilis).Error
	if err != nil {
		return domisilis, err
	}

	return domisilis, nil
}

func (r *repository) FindLast() (Domisili, error) {
	var domisilis Domisili
	err := r.db.Last(&domisilis).Error
	if err != nil {
		return domisilis, err
	}

	return domisilis, nil
}

func (r *repository) Save(domisili Domisili) (Domisili, error) {
	err := r.db.Create(&domisili).Error
	if err != nil {
		return domisili, err
	}

	return domisili, nil
}

func (r *repository) Update(domisili Domisili) (Domisili, error) {
	err := r.db.Save(&domisili).Error
	if err != nil {
		return domisili, err
	}

	return domisili, nil
}

func (r *repository) Delete(domisili Domisili) error {
	err := r.db.Delete(&domisili).Error
	if err != nil {
		return err
	}

	return nil
}
