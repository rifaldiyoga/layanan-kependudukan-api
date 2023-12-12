package district

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetDistrictByID(ID int) (District, error)
	GetDistricts(pagination helper.Pagination, provinceId int) (helper.Pagination, error)
	CreateDistrict(input CreateDistrictInput) (District, error)
	UpdateDistrict(ID GetDistrictDetailInput, input CreateDistrictInput) (District, error)
	DeleteDistrict(ID GetDistrictDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetDistrictByID(ID int) (District, error) {
	district, err := s.repository.FindByID(ID)
	if err != nil {
		return district, err
	}

	return district, nil
}

func (s *service) CreateDistrict(input CreateDistrictInput) (District, error) {
	district := District{}

	district.Code = input.Code
	district.Name = input.Name
	district.CreatedAt = time.Now()

	newDistrict, err := s.repository.Save(district)
	return newDistrict, err
}

func (s *service) UpdateDistrict(inputDetail GetDistrictDetailInput, input CreateDistrictInput) (District, error) {
	district, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return district, err
	}

	district.Code = input.Code
	district.Name = input.Name
	district.UpdatedAt = time.Now()

	newDistrict, err := s.repository.Update(district)
	return newDistrict, err
}

func (s *service) DeleteDistrict(inputDetail GetDistrictDetailInput) error {
	district, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(district)
	return err
}

func (s *service) GetDistricts(pagination helper.Pagination, provinceId int) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, provinceId)
	return pagination, err
}
