package keluarga

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetKeluargaByID(ID int) (Keluarga, error)
	GetKeluargas(pagination helper.Pagination) (helper.Pagination, error)
	CreateKeluarga(input CreateKeluargaInput) (Keluarga, error)
	UpdateKeluarga(ID GetKeluargaDetailInput, input CreateKeluargaInput) (Keluarga, error)
	DeleteKeluarga(ID GetKeluargaDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetKeluargaByID(ID int) (Keluarga, error) {
	penduduk, err := s.repository.FindByID(ID)

	return penduduk, err
}

func (s *service) CreateKeluarga(input CreateKeluargaInput) (Keluarga, error) {
	penduduk := Keluarga{}

	penduduk.UpdatedAt = time.Now()

	newKeluarga, err := s.repository.Save(penduduk)

	return newKeluarga, err
}

func (s *service) UpdateKeluarga(inputDetail GetKeluargaDetailInput, input CreateKeluargaInput) (Keluarga, error) {
	penduduk := Keluarga{}
	penduduk.ID = inputDetail.ID
	penduduk.UpdatedAt = time.Now()

	newKeluarga, err := s.repository.Update(penduduk)
	return newKeluarga, err
}

func (s *service) DeleteKeluarga(inputDetail GetKeluargaDetailInput) error {
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

func (s *service) GetKeluargas(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
