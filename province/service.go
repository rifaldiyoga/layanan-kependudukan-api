package province

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetProvinceByID(ID int) (Province, error)
	GetProvinces(pagination helper.Pagination) (helper.Pagination, error)
	CreateProvince(input CreateProvinceInput) (Province, error)
	UpdateProvince(ID GetProvinceDetailInput, input CreateProvinceInput) (Province, error)
	DeleteProvince(ID GetProvinceDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetProvinceByID(ID int) (Province, error) {
	Province, err := s.repository.FindByID(ID)
	if err != nil {
		return Province, err
	}

	return Province, nil
}

func (s *service) CreateProvince(input CreateProvinceInput) (Province, error) {
	province := Province{}

	province.Code = input.Code
	province.Name = input.Name
	province.CreatedAt = time.Now()

	newProvince, err := s.repository.Save(province)
	return newProvince, err
}

func (s *service) UpdateProvince(inputDetail GetProvinceDetailInput, input CreateProvinceInput) (Province, error) {
	province, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return province, err
	}

	province.Code = input.Code
	province.Name = input.Name
	province.UpdatedAt = time.Now()

	newProvince, err := s.repository.Update(province)
	return newProvince, err
}

func (s *service) DeleteProvince(inputDetail GetProvinceDetailInput) error {
	province, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(province)
	return err
}

func (s *service) GetProvinces(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
