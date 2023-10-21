package user

import (
	"layanan-kependudukan-api/helper"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination) (helper.Pagination, error)
	Save(user User) (User, error)
	Update(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(id int) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination) (helper.Pagination, error) {
	var users []User

	err := r.db.Scopes(helper.Paginate(users, &pagination, r.db)).Find(&users).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = users
	return pagination, err
}

func (r *repository) Update(subDistrict User) (User, error) {
	err := r.db.Save(&subDistrict).Error
	if err != nil {
		return subDistrict, err
	}

	return subDistrict, nil
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(id int) (User, error) {
	var user User
	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
