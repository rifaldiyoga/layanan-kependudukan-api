package sistem

import (
	"layanan-kependudukan-api/helper"
)

type Service interface {
	GetSistemByID(ID int) (Sistem, error)
	GetSistems(pagination helper.Pagination) (helper.Pagination, error)
	CreateSistem(input CreateSistemInput) (Sistem, error)
	UpdateSistem(ID GetSistemDetailInput, input CreateSistemInput) (Sistem, error)
	DeleteSistem(ID GetSistemDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetSistemByID(ID int) (Sistem, error) {
	sistem, err := s.repository.FindByID(ID)

	return sistem, err
}

func (s *service) CreateSistem(input CreateSistemInput) (Sistem, error) {
	sistem := Sistem{}

	sistem.Code = input.Code

	newSistem, err := s.repository.Save(sistem)

	return newSistem, err
}

func (s *service) UpdateSistem(inputDetail GetSistemDetailInput, input CreateSistemInput) (Sistem, error) {
	sistem := Sistem{}
	sistem.ID = inputDetail.ID
	sistem.Code = input.Code

	newSistem, err := s.repository.Update(sistem)
	return newSistem, err
}

func (s *service) DeleteSistem(inputDetail GetSistemDetailInput) error {
	sistem, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}
	sistem.ID = inputDetail.ID

	err := s.repository.Delete(sistem)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetSistems(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
