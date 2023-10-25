package layanan

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetLayananByID(ID int) (Layanan, error)
	GetLayanansPaging(pagination helper.Pagination) (helper.Pagination, error)
	GetLayanans() ([]Layanan, error)
	GetRekomLayanans() ([]Layanan, error)
	GetTypes() ([]string, error)
	CreateLayanan(input CreateLayananInput) (Layanan, error)
	UpdateLayanan(ID GetLayananDetailInput, input CreateLayananInput) (Layanan, error)
	DeleteLayanan(ID GetLayananDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetLayananByID(ID int) (Layanan, error) {
	Layanan, err := s.repository.FindByID(ID)
	if err != nil {
		return Layanan, err
	}

	return Layanan, nil
}

func (s *service) CreateLayanan(input CreateLayananInput) (Layanan, error) {
	Layanan := Layanan{}

	Layanan.Code = input.Code
	Layanan.Name = input.Name
	Layanan.Type = input.Type
	Layanan.IsConfirm = input.IsConfirm
	Layanan.CreatedAt = time.Now()

	newLayanan, err := s.repository.Save(Layanan)
	return newLayanan, err
}

func (s *service) UpdateLayanan(inputDetail GetLayananDetailInput, input CreateLayananInput) (Layanan, error) {
	Layanan, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return Layanan, err
	}

	Layanan.Code = input.Code
	Layanan.Name = input.Name
	Layanan.Type = input.Type
	Layanan.IsConfirm = input.IsConfirm
	Layanan.UpdatedAt = time.Now()

	newLayanan, err := s.repository.Update(Layanan)
	return newLayanan, err
}

func (s *service) DeleteLayanan(inputDetail GetLayananDetailInput) error {
	Layanan, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(Layanan)
	return err
}

func (s *service) GetRekomLayanans() ([]Layanan, error) {
	layanan, err := s.repository.FindRecom()
	return layanan, err
}

func (s *service) GetLayanans() ([]Layanan, error) {
	layanan, err := s.repository.FindAll()
	return layanan, err
}

func (s *service) GetTypes() ([]string, error) {
	layanan, err := s.repository.FindByType()
	return layanan, err
}

func (s *service) GetLayanansPaging(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAllPaging(pagination)

	return pagination, err
}
