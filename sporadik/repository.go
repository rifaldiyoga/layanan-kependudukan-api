package sporadik

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (Sporadik, error)
	FindLast() (Sporadik, error)
	Save(sporadik Sporadik) (Sporadik, error)
	Update(sporadik Sporadik) (Sporadik, error)
	Delete(sporadik Sporadik) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var sporadiks []Sporadik

	db := r.db.Debug()

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Scopes(helper.Paginate(sporadiks, &pagination, r.db)).Where("status = true").Find(&sporadiks).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = sporadiks
	return pagination, err
}

func (r *repository) FindByID(ID int) (Sporadik, error) {
	var sporadiks Sporadik
	err := r.db.Where("id = ?", ID).Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi").First(&sporadiks).Error
	if err != nil {
		return sporadiks, err
	}

	return sporadiks, nil
}

func (r *repository) FindLast() (Sporadik, error) {
	var sporadiks Sporadik
	err := r.db.Last(&sporadiks).Error
	if err != nil {
		return sporadiks, err
	}

	return sporadiks, nil
}

func (r *repository) Save(sporadik Sporadik) (Sporadik, error) {
	err := r.db.Create(&sporadik).Error
	if err != nil {
		return sporadik, err
	}

	return sporadik, nil
}

func (r *repository) Update(sporadik Sporadik) (Sporadik, error) {
	err := r.db.Save(&sporadik).Error
	if err != nil {
		return sporadik, err
	}

	return sporadik, nil
}

func (r *repository) Delete(sporadik Sporadik) error {
	err := r.db.Delete(&sporadik).Error
	if err != nil {
		return err
	}

	return nil
}
