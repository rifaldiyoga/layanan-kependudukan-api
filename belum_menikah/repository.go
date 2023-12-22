package belum_menikah

import (
	"layanan-kependudukan-api/helper"
	"net/url"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByID(id int) (BelumMenikah, error)
	FindLast() (BelumMenikah, error)
	Save(janda BelumMenikah) (BelumMenikah, error)
	Update(janda BelumMenikah) (BelumMenikah, error)
	Delete(janda BelumMenikah) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var jandas []BelumMenikah
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

func (r *repository) FindByID(ID int) (BelumMenikah, error) {
	var jandas BelumMenikah
	err := r.db.Where("id = ?", ID).Preload("Penduduk.Religion").Preload("Penduduk.Job").Preload("Penduduk.Education").Preload("Penduduk.RT").Preload("Penduduk.RW").Preload("Penduduk.Kelurahan").Preload("Penduduk.Kecamatan").Preload("Penduduk.Kota").Preload("Penduduk.Provinsi").First(&jandas).Error
	if err != nil {
		return jandas, err
	}

	return jandas, nil
}

func (r *repository) FindLast() (BelumMenikah, error) {
	var jandas BelumMenikah
	err := r.db.Last(&jandas).Error
	if err != nil {
		return jandas, err
	}

	return jandas, nil
}

func (r *repository) Save(janda BelumMenikah) (BelumMenikah, error) {
	err := r.db.Create(&janda).Error
	if err != nil {
		return janda, err
	}

	return janda, nil
}

func (r *repository) Update(janda BelumMenikah) (BelumMenikah, error) {
	err := r.db.Save(&janda).Error
	if err != nil {
		return janda, err
	}

	return janda, nil
}

func (r *repository) Delete(janda BelumMenikah) error {
	err := r.db.Delete(&janda).Error
	if err != nil {
		return err
	}

	return nil
}
