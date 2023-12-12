package rumah

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (Rumah, error)
	FindLast() (Rumah, error)
	Save(rumah Rumah) (Rumah, error)
	Update(rumah Rumah) (Rumah, error)
	Delete(rumah Rumah) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var rumahs []Rumah

	db := r.db.Debug()

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Scopes(helper.Paginate(rumahs, &pagination, r.db)).Where("status = true").Find(&rumahs).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = rumahs
	return pagination, err
}

func (r *repository) FindByID(ID int) (Rumah, error) {
	var rumahs Rumah
	err := r.db.Where("id = ?", ID).Preload(clause.Associations).Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi").First(&rumahs).Error
	if err != nil {
		return rumahs, err
	}

	return rumahs, nil
}

func (r *repository) FindLast() (Rumah, error) {
	var rumahs Rumah
	err := r.db.Last(&rumahs).Error
	if err != nil {
		return rumahs, err
	}

	return rumahs, nil
}

func (r *repository) Save(rumah Rumah) (Rumah, error) {
	err := r.db.Create(&rumah).Error
	if err != nil {
		return rumah, err
	}

	return rumah, nil
}

func (r *repository) Update(rumah Rumah) (Rumah, error) {
	err := r.db.Save(&rumah).Error
	if err != nil {
		return rumah, err
	}

	return rumah, nil
}

func (r *repository) Delete(rumah Rumah) error {
	err := r.db.Delete(&rumah).Error
	if err != nil {
		return err
	}

	return nil
}
