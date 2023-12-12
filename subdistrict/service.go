package subdistrict

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetSubDistrictByID(ID int) (SubDistrict, error)
	GetSubDistricts(pagination helper.Pagination, districtId int) (helper.Pagination, error)
	CreateSubDistrict(input CreateSubDistrictInput) (SubDistrict, error)
	UpdateSubDistrict(ID GetSubDistrictDetailInput, input CreateSubDistrictInput) (SubDistrict, error)
	DeleteSubDistrict(ID GetSubDistrictDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetSubDistrictByID(ID int) (SubDistrict, error) {
	SubDistrict, err := s.repository.FindByID(ID)
	if err != nil {
		return SubDistrict, err
	}

	return SubDistrict, nil
}

func (s *service) CreateSubDistrict(input CreateSubDistrictInput) (SubDistrict, error) {
	subDistrict := SubDistrict{}

	subDistrict.Code = input.Code
	subDistrict.Name = input.Name
	subDistrict.CreatedAt = time.Now()

	newSubDistrict, err := s.repository.Save(subDistrict)
	return newSubDistrict, err
}

func (s *service) UpdateSubDistrict(inputDetail GetSubDistrictDetailInput, input CreateSubDistrictInput) (SubDistrict, error) {
	subDistrict, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return subDistrict, err
	}

	subDistrict.Code = input.Code
	subDistrict.Name = input.Name
	subDistrict.UpdatedAt = time.Now()

	newSubDistrict, err := s.repository.Update(subDistrict)
	return newSubDistrict, err
}

func (s *service) DeleteSubDistrict(inputDetail GetSubDistrictDetailInput) error {
	subDistrict, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(subDistrict)
	return err
}

func (s *service) GetSubDistricts(pagination helper.Pagination, districtId int) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, districtId)

	return pagination, err
}
