package aparatur_desa

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetAparaturDesaByID(ID int) (AparaturDesa, error)
	GetAparaturDesas(pagination helper.Pagination) (helper.Pagination, error)
	CreateAparaturDesa(input CreateAparaturDesaInput) (AparaturDesa, error)
	UpdateAparaturDesa(ID GetAparaturDesaDetailInput, input CreateAparaturDesaInput) (AparaturDesa, error)
	DeleteAparaturDesa(ID GetAparaturDesaDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetAparaturDesaByID(ID int) (AparaturDesa, error) {
	AparaturDesa, err := s.repository.FindByID(ID)
	if err != nil {
		return AparaturDesa, err
	}

	return AparaturDesa, nil
}

func (s *service) CreateAparaturDesa(input CreateAparaturDesaInput) (AparaturDesa, error) {
	kelurahan := AparaturDesa{}

	kelurahan.Code = input.Code
	kelurahan.Name = input.Name
	kelurahan.CreatedAt = time.Now()

	newAparaturDesa, err := s.repository.Save(kelurahan)
	return newAparaturDesa, err
}

func (s *service) UpdateAparaturDesa(inputDetail GetAparaturDesaDetailInput, input CreateAparaturDesaInput) (AparaturDesa, error) {
	AparaturDesa, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return AparaturDesa, err
	}

	AparaturDesa.Code = input.Code
	AparaturDesa.Name = input.Name
	AparaturDesa.UpdatedAt = time.Now()

	newAparaturDesa, err := s.repository.Update(AparaturDesa)
	return newAparaturDesa, err
}

func (s *service) DeleteAparaturDesa(inputDetail GetAparaturDesaDetailInput) error {
	kelurahan, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(kelurahan)
	return err
}

func (s *service) GetAparaturDesas(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
