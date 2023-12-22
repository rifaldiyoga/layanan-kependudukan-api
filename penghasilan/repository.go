package penghasilan

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (Penghasilan, error)
	FindLast() (Penghasilan, error)
	Save(penghasilan Penghasilan) (Penghasilan, error)
	Update(penghasilan Penghasilan) (Penghasilan, error)
	Delete(penghasilan Penghasilan) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var penghasilans []Penghasilan

	db := r.db.Debug()

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Order("created_at DESC").Preload(clause.Associations).Scopes(helper.Paginate(penghasilans, &pagination, r.db)).Where("status = true").Find(&penghasilans).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = penghasilans
	return pagination, err
}

func (r *repository) FindByID(ID int) (Penghasilan, error) {
	var penghasilans Penghasilan
	err := r.db.Where("id = ?", ID).Preload(clause.Associations).Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi").First(&penghasilans).Error
	if err != nil {
		return penghasilans, err
	}

	return penghasilans, nil
}

func (r *repository) FindLast() (Penghasilan, error) {
	var penghasilans Penghasilan
	err := r.db.Last(&penghasilans).Error
	if err != nil {
		return penghasilans, err
	}

	return penghasilans, nil
}

func (r *repository) Save(penghasilan Penghasilan) (Penghasilan, error) {
	err := r.db.Create(&penghasilan).Error
	if err != nil {
		return penghasilan, err
	}

	return penghasilan, nil
}

func (r *repository) Update(penghasilan Penghasilan) (Penghasilan, error) {
	err := r.db.Save(&penghasilan).Error
	if err != nil {
		return penghasilan, err
	}

	return penghasilan, nil
}

func (r *repository) Delete(penghasilan Penghasilan) error {
	err := r.db.Delete(&penghasilan).Error
	if err != nil {
		return err
	}

	return nil
}
