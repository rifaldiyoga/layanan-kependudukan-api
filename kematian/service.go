package kematian

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetKematianByID(ID int) (Kematian, error)
	GetKematians(pagination helper.Pagination) (helper.Pagination, error)
	CreateKematian(input CreateKematianInput) (Kematian, error)
	UpdateKematian(ID GetKematianDetailInput, input CreateKematianInput) (Kematian, error)
	DeleteKematian(ID GetKematianDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetKematianByID(ID int) (Kematian, error) {
	penduduk, err := s.repository.FindByID(ID)

	return penduduk, err
}

func (s *service) CreateKematian(input CreateKematianInput) (Kematian, error) {
	penduduk := Kematian{}

	penduduk.UpdatedAt = time.Now()

	newKematian, err := s.repository.Save(penduduk)

	return newKematian, err
}

func (s *service) UpdateKematian(inputDetail GetKematianDetailInput, input CreateKematianInput) (Kematian, error) {
	penduduk := Kematian{}
	penduduk.ID = inputDetail.ID
	penduduk.UpdatedAt = time.Now()

	newKematian, err := s.repository.Update(penduduk)
	return newKematian, err
}

func (s *service) DeleteKematian(inputDetail GetKematianDetailInput) error {
	penduduk, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}
	penduduk.ID = inputDetail.ID

	err := s.repository.Delete(penduduk)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetKematians(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
