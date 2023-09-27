package rw

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetRWByID(ID int) (RW, error)
	GetRWs(pagination helper.Pagination) (helper.Pagination, error)
	CreateRW(input CreateRWInput) (RW, error)
	UpdateRW(ID GetRWDetailInput, input CreateRWInput) (RW, error)
	DeleteRW(ID GetRWDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetRWByID(ID int) (RW, error) {
	rw, err := s.repository.FindByID(ID)
	if err != nil {
		return rw, err
	}

	return rw, nil
}

func (s *service) CreateRW(input CreateRWInput) (RW, error) {
	rw := RW{}

	rw.Code = input.Code
	rw.Name = input.Name
	rw.CreatedAt = time.Now()

	newRW, err := s.repository.Save(rw)
	return newRW, err
}

func (s *service) UpdateRW(inputDetail GetRWDetailInput, input CreateRWInput) (RW, error) {
	rw, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return rw, err
	}

	rw.Code = input.Code
	rw.Name = input.Name
	rw.UpdatedAt = time.Now()

	newRW, err := s.repository.Update(rw)
	return newRW, err
}

func (s *service) DeleteRW(inputDetail GetRWDetailInput) error {
	rw, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(rw)
	return err
}

func (s *service) GetRWs(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
