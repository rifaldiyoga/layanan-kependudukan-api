package position

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetPositionByID(ID int) (Position, error)
	GetPositions(pagination helper.Pagination) (helper.Pagination, error)
	CreatePosition(input CreatePositionInput) (Position, error)
	UpdatePosition(ID GetPositionDetailInput, input CreatePositionInput) (Position, error)
	DeletePosition(ID GetPositionDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetPositionByID(ID int) (Position, error) {
	position, err := s.repository.FindByID(ID)
	if err != nil {
		return position, err
	}

	return position, nil
}

func (s *service) CreatePosition(input CreatePositionInput) (Position, error) {
	position := Position{}

	position.Code = input.Code
	position.Jabatan = input.Jabatan
	position.CreatedAt = time.Now()

	newPosition, err := s.repository.Save(position)
	return newPosition, err
}

func (s *service) UpdatePosition(inputDetail GetPositionDetailInput, input CreatePositionInput) (Position, error) {
	position, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return position, err
	}

	position.Code = input.Code
	position.Jabatan = input.Jabatan
	position.UpdatedAt = time.Now()

	newPosition, err := s.repository.Update(position)
	return newPosition, err
}

func (s *service) DeletePosition(inputDetail GetPositionDetailInput) error {
	position, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(position)
	return err
}

func (s *service) GetPositions(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
