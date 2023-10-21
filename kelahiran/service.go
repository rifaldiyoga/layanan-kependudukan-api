package kelahiran

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetKelahiranByID(ID int) (Kelahiran, error)
	GetKelahirans(pagination helper.Pagination) (helper.Pagination, error)
	CreateKelahiran(input CreateKelahiranInput) (Kelahiran, error)
	UpdateKelahiran(ID GetKelahiranDetailInput, input CreateKelahiranInput) (Kelahiran, error)
	DeleteKelahiran(ID GetKelahiranDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetKelahiranByID(ID int) (Kelahiran, error) {
	penduduk, err := s.repository.FindByID(ID)

	return penduduk, err
}

func (s *service) CreateKelahiran(input CreateKelahiranInput) (Kelahiran, error) {
	penduduk := Kelahiran{}

	penduduk.UpdatedAt = time.Now()

	newKelahiran, err := s.repository.Save(penduduk)

	return newKelahiran, err
}

func (s *service) UpdateKelahiran(inputDetail GetKelahiranDetailInput, input CreateKelahiranInput) (Kelahiran, error) {
	penduduk := Kelahiran{}
	penduduk.ID = inputDetail.ID
	penduduk.UpdatedAt = time.Now()

	newKelahiran, err := s.repository.Update(penduduk)
	return newKelahiran, err
}

func (s *service) DeleteKelahiran(inputDetail GetKelahiranDetailInput) error {
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

func (s *service) GetKelahirans(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
