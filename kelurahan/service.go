package kelurahan

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetKelurahanByID(ID int) (Kelurahan, error)
	GetKelurahans(pagination helper.Pagination) (helper.Pagination, error)
	CreateKelurahan(input CreateKelurahanInput) (Kelurahan, error)
	UpdateKelurahan(ID GetKelurahanDetailInput, input CreateKelurahanInput) (Kelurahan, error)
	DeleteKelurahan(ID GetKelurahanDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetKelurahanByID(ID int) (Kelurahan, error) {
	Kelurahan, err := s.repository.FindByID(ID)
	if err != nil {
		return Kelurahan, err
	}

	return Kelurahan, nil
}

func (s *service) CreateKelurahan(input CreateKelurahanInput) (Kelurahan, error) {
	kelurahan := Kelurahan{}

	kelurahan.Code = input.Code
	kelurahan.Name = input.Name
	kelurahan.CreatedAt = time.Now()

	newKelurahan, err := s.repository.Save(kelurahan)
	return newKelurahan, err
}

func (s *service) UpdateKelurahan(inputDetail GetKelurahanDetailInput, input CreateKelurahanInput) (Kelurahan, error) {
	Kelurahan, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return Kelurahan, err
	}

	Kelurahan.Code = input.Code
	Kelurahan.Name = input.Name
	Kelurahan.UpdatedAt = time.Now()

	newKelurahan, err := s.repository.Update(Kelurahan)
	return newKelurahan, err
}

func (s *service) DeleteKelurahan(inputDetail GetKelurahanDetailInput) error {
	kelurahan, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(kelurahan)
	return err
}

func (s *service) GetKelurahans(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
