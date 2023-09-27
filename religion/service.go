package religion

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetReligionByID(ID int) (Religion, error)
	GetReligions(pagination helper.Pagination) (helper.Pagination, error)
	CreateReligion(input CreateReligionInput) (Religion, error)
	UpdateReligion(ID GetReligionDetailInput, input CreateReligionInput) (Religion, error)
	DeleteReligion(ID GetReligionDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetReligionByID(ID int) (Religion, error) {
	religion, err := s.repository.FindByID(ID)

	return religion, err
}

func (s *service) CreateReligion(input CreateReligionInput) (Religion, error) {
	religion := Religion{}

	religion.Code = input.Code
	religion.Name = input.Name
	religion.UpdatedAt = time.Now()

	newReligion, err := s.repository.Save(religion)

	return newReligion, err
}

func (s *service) UpdateReligion(inputDetail GetReligionDetailInput, input CreateReligionInput) (Religion, error) {
	religion := Religion{}
	religion.ID = inputDetail.ID
	religion.Code = input.Code
	religion.Name = input.Name
	religion.UpdatedAt = time.Now()

	newReligion, err := s.repository.Update(religion)
	return newReligion, err
}

func (s *service) DeleteReligion(inputDetail GetReligionDetailInput) error {
	religion, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}
	religion.ID = inputDetail.ID

	err := s.repository.Delete(religion)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetReligions(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
