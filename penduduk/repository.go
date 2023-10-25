package penduduk

import (
	"layanan-kependudukan-api/helper"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination, NIK string) (helper.Pagination, error)
	FindByID(id int) (Penduduk, error)
	FindByNIK(id string) (Penduduk, error)
	FindByDate(date time.Time) (Penduduk, error)
	Save(penduduk Penduduk) (Penduduk, error)
	Update(penduduk Penduduk) (Penduduk, error)
	Delete(penduduk Penduduk) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, NIK string) (helper.Pagination, error) {
	var penduduks []Penduduk
	data := r.db.Debug()
	if NIK != "" {
		data = data.Where("no_kk = ?", NIK)
	}
	err := data.Scopes(helper.Paginate(penduduks, &pagination, data)).Find(&penduduks).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = penduduks
	return pagination, err
}

func (r *repository) FindByID(ID int) (Penduduk, error) {
	var penduduks Penduduk
	err := r.db.Where("id = ?", ID).First(&penduduks).Error
	if err != nil {
		return penduduks, err
	}

	return penduduks, nil
}

func (r *repository) FindByNIK(ID string) (Penduduk, error) {
	var penduduks Penduduk
	err := r.db.Where("nik = ?", ID).First(&penduduks).Error
	if err != nil {
		return penduduks, err
	}

	return penduduks, nil
}

func (r *repository) FindByDate(date time.Time) (Penduduk, error) {
	var penduduks Penduduk
	err := r.db.Where("birth_date = ?", date.Format("2006-01-02")).Order("id DESC").First(&penduduks).Error
	if err != nil {
		return penduduks, err
	}

	return penduduks, nil
}

func (r *repository) Save(penduduk Penduduk) (Penduduk, error) {
	err := r.db.Create(&penduduk).Error
	if err != nil {
		return penduduk, err
	}

	return penduduk, nil
}

func (r *repository) Update(penduduk Penduduk) (Penduduk, error) {
	err := r.db.Save(&penduduk).Error
	if err != nil {
		return penduduk, err
	}

	return penduduk, nil
}

func (r *repository) Delete(penduduk Penduduk) error {
	err := r.db.Delete(&penduduk).Error
	if err != nil {
		return err
	}

	return nil
}
