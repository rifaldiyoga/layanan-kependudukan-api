package penduduk

import (
	"layanan-kependudukan-api/helper"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll(pagination helper.Pagination, NIK string) (helper.Pagination, error)
	FindByID(id int) (Penduduk, error)
	FindByNIK(id string) (Penduduk, error)
	FindByNoKK(id string) ([]Penduduk, error)
	FindByDate(date time.Time) (Penduduk, error)
	FindByRT(RtID int, RwID int) (Penduduk, error)
	FindByRW(RwID int) (Penduduk, error)
	CountAll() (int64, error)
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
	data = data.Where("active = ?", true)
	err := data.Scopes(helper.Paginate(penduduks, &pagination, data)).Find(&penduduks).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = penduduks
	return pagination, err
}

func (r *repository) FindByID(ID int) (Penduduk, error) {
	var penduduks Penduduk
	// err := r.db.Where("id = ?", ID).Preload("Religion").Preload("Job").Preload("Education").Preload("RT").Preload("RW").Preload("Kelurahan").Preload("Kecamatan").Preload("Kota").Preload("Provinsi").First(&penduduks).Error
	err := r.db.Where("id = ?", ID).Preload(clause.Associations).First(&penduduks).Error
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

func (r *repository) FindByNoKK(ID string) ([]Penduduk, error) {
	var penduduks []Penduduk
	err := r.db.Where("no_kk = ? AND active = ?", ID, true).Find(&penduduks).Error
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

func (r *repository) FindByRT(RtID int, RwID int) (Penduduk, error) {
	var penduduks Penduduk
	err := r.db.Debug().Joins("JOIN tb_user ON tb_user.nik = tb_penduduk.nik").Where("tb_penduduk.rt_id = ? AND tb_penduduk.rw_id = ? AND tb_user.role = 'RT'", RtID, RwID).First(&penduduks).Error
	if err != nil {
		return penduduks, err
	}

	return penduduks, nil
}

func (r *repository) FindByRW(RwID int) (Penduduk, error) {
	var penduduks Penduduk
	err := r.db.Debug().Joins("JOIN tb_user ON tb_user.nik = tb_penduduk.nik").Where("tb_penduduk.rw_id = ? AND tb_user.role = 'RW'", RwID).First(&penduduks).Error
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

func (r *repository) CountAll() (int64, error) {
	var count int64
	err := r.db.Model(&Penduduk{}).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}
