package status

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetStatusByID(ID int) (Status, error)
	GetStatuss(pagination helper.Pagination) (helper.Pagination, error)
	CreateStatus(input CreateStatusInput) (Status, error)
	UpdateStatus(ID GetStatusDetailInput, input CreateStatusInput) (Status, error)
	DeleteStatus(ID GetStatusDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetStatusByID(ID int) (Status, error) {
	status, err := s.repository.FindByID(ID)

	return status, err
}

func (s *service) CreateStatus(input CreateStatusInput) (Status, error) {
	status := Status{}

	status.Code = input.Code
	status.Name = input.Name
	status.UpdatedAt = time.Now()

	newStatus, err := s.repository.Save(status)

	return newStatus, err
}

func (s *service) UpdateStatus(inputDetail GetStatusDetailInput, input CreateStatusInput) (Status, error) {
	status := Status{}
	status.ID = inputDetail.ID
	status.Code = input.Code
	status.Name = input.Name
	status.UpdatedAt = time.Now()

	newStatus, err := s.repository.Update(status)
	return newStatus, err
}

func (s *service) DeleteStatus(inputDetail GetStatusDetailInput) error {
	status, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}
	status.ID = inputDetail.ID

	err := s.repository.Delete(status)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetStatuss(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
