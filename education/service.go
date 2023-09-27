package education

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetEducationByID(ID int) (Education, error)
	GetEducations(pagination helper.Pagination) (helper.Pagination, error)
	CreateEducation(input CreateEducationInput) (Education, error)
	UpdateEducation(ID GetEducationDetailInput, input CreateEducationInput) (Education, error)
	DeleteEducation(ID GetEducationDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetEducationByID(ID int) (Education, error) {
	education, err := s.repository.FindByID(ID)
	if err != nil {
		return education, err
	}

	return education, nil
}

func (s *service) CreateEducation(input CreateEducationInput) (Education, error) {
	education := Education{}

	education.Code = input.Code
	education.Name = input.Name
	education.CreatedAt = time.Now()

	newEducation, err := s.repository.Save(education)
	return newEducation, err
}

func (s *service) UpdateEducation(inputDetail GetEducationDetailInput, input CreateEducationInput) (Education, error) {
	education, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return education, err
	}

	education.Code = input.Code
	education.Name = input.Name
	education.UpdatedAt = time.Now()

	newEducation, err := s.repository.Update(education)
	return newEducation, err
}

func (s *service) DeleteEducation(inputDetail GetEducationDetailInput) error {
	education, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(education)
	return err
}

func (s *service) GetEducations(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)
	return pagination, err
}
