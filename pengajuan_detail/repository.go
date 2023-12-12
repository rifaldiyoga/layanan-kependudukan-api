package pengajuan_detail

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	FindByID(id int) (DetailPengajuan, error)
	FindByPengajuan(id int) ([]DetailPengajuan, error)
	Save(detailPengajuan DetailPengajuan) (DetailPengajuan, error)
	Update(detailPengajuan DetailPengajuan) (DetailPengajuan, error)
	Delete(detailPengajuan DetailPengajuan) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var detailPengajuans []DetailPengajuan

	err := r.db.Scopes(helper.Paginate(detailPengajuans, &pagination, r.db)).Find(&detailPengajuans).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = detailPengajuans
	return pagination, err
}

func (r *repository) FindByPengajuan(ID int) ([]DetailPengajuan, error) {
	var detailPengajuans []DetailPengajuan

	err := r.db.Where("pengajuan_id = ?", ID).Find(&detailPengajuans).Error
	if err != nil {
		return detailPengajuans, err
	}
	return detailPengajuans, err
}

func (r *repository) FindByID(ID int) (DetailPengajuan, error) {
	var detailPengajuan DetailPengajuan
	err := r.db.Where("id = ?", ID).First(&detailPengajuan).Error
	if err != nil {
		return detailPengajuan, err
	}

	return detailPengajuan, nil
}

func (r *repository) Save(detailPengajuan DetailPengajuan) (DetailPengajuan, error) {
	err := r.db.Create(&detailPengajuan).Error
	if err != nil {
		return detailPengajuan, err
	}

	return detailPengajuan, nil
}

func (r *repository) Update(detailPengajuan DetailPengajuan) (DetailPengajuan, error) {
	err := r.db.Save(&detailPengajuan).Error
	if err != nil {
		return detailPengajuan, err
	}

	return detailPengajuan, nil
}

func (r *repository) Delete(detailPengajuan DetailPengajuan) error {
	err := r.db.Delete(&detailPengajuan).Error
	if err != nil {
		return err
	}

	return nil
}
