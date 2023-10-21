package layanan

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAllPaging(pagination helper.Pagination) (helper.Pagination, error)
	FindAll() ([]Layanan, error)
	FindRecom() ([]Layanan, error)
	FindByType() ([]string, error)
	FindByID(id int) (Layanan, error)
	Save(layanan Layanan) (Layanan, error)
	Update(layanan Layanan) (Layanan, error)
	Delete(layanan Layanan) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllPaging(pagination helper.Pagination) (helper.Pagination, error) {
	var layanans []Layanan

	err := r.db.Scopes(helper.Paginate(layanans, &pagination, r.db)).Find(&layanans).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = layanans
	return pagination, err
}

func (r *repository) FindAll() ([]Layanan, error) {
	var Layanans []Layanan

	err := r.db.Find(&Layanans).Error
	if err != nil {
		return nil, err
	}

	return Layanans, err
}

func (r *repository) FindRecom() ([]Layanan, error) {
	var Layanans []Layanan

	err := r.db.Find(&Layanans).Limit(10).Error
	if err != nil {
		return nil, err
	}

	return Layanans, err
}

func (r *repository) FindByType() ([]string, error) {
	var types []string

	err := r.db.Table("tb_layanan").Select("type").Group("type").Find(&types).Error
	if err != nil {
		return nil, err
	}

	return types, err
}

func (r *repository) FindByID(ID int) (Layanan, error) {
	var Layanan Layanan
	err := r.db.Where("id = ?", ID).First(&Layanan).Error
	if err != nil {
		return Layanan, err
	}

	return Layanan, nil
}

func (r *repository) Save(layanan Layanan) (Layanan, error) {
	err := r.db.Create(&layanan).Error
	if err != nil {
		return layanan, err
	}

	return layanan, nil
}

func (r *repository) Update(layanan Layanan) (Layanan, error) {
	err := r.db.Save(&layanan).Error
	if err != nil {
		return layanan, err
	}

	return layanan, nil
}

func (r *repository) Delete(layanan Layanan) error {
	err := r.db.Delete(&layanan).Error
	if err != nil {
		return err
	}

	return nil
}
