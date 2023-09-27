package rt

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetRTByID(ID int) (RT, error)
	GetRTs(pagination helper.Pagination) (helper.Pagination, error)
	CreateRT(input CreateRTInput) (RT, error)
	UpdateRT(ID GetRTDetailInput, input CreateRTInput) (RT, error)
	DeleteRT(ID GetRTDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetRTByID(ID int) (RT, error) {
	RT, err := s.repository.FindByID(ID)
	if err != nil {
		return RT, err
	}

	return RT, nil
}

func (s *service) CreateRT(input CreateRTInput) (RT, error) {
	rt := RT{}

	rt.Code = input.Code
	rt.Name = input.Name
	rt.CreatedAt = time.Now()

	newRT, err := s.repository.Save(rt)
	return newRT, err
}

func (s *service) UpdateRT(inputDetail GetRTDetailInput, input CreateRTInput) (RT, error) {
	rt, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return rt, err
	}

	rt.Code = input.Code
	rt.Name = input.Name
	rt.UpdatedAt = time.Now()

	newRT, err := s.repository.Update(rt)
	return newRT, err
}

func (s *service) DeleteRT(inputDetail GetRTDetailInput) error {
	rt, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(rt)
	return err
}

func (s *service) GetRTs(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
