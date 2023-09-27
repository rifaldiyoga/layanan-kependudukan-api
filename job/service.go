package job

import (
	"layanan-kependudukan-api/helper"
	"time"
)

type Service interface {
	GetJobByID(ID int) (Job, error)
	GetJobs(pagination helper.Pagination) (helper.Pagination, error)
	CreateJob(input CreateJobInput) (Job, error)
	UpdateJob(ID GetJobDetailInput, input CreateJobInput) (Job, error)
	DeleteJob(ID GetJobDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetJobByID(ID int) (Job, error) {
	Job, err := s.repository.FindByID(ID)
	if err != nil {
		return Job, err
	}

	return Job, nil
}

func (s *service) CreateJob(input CreateJobInput) (Job, error) {
	job := Job{}

	job.Code = input.Code
	job.Name = input.Name
	job.CreatedAt = time.Now()

	newJob, err := s.repository.Save(job)
	return newJob, err
}

func (s *service) UpdateJob(inputDetail GetJobDetailInput, input CreateJobInput) (Job, error) {
	job, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return job, err
	}

	job.Code = input.Code
	job.Name = input.Name
	job.UpdatedAt = time.Now()

	newJob, err := s.repository.Update(job)
	return newJob, err
}

func (s *service) DeleteJob(inputDetail GetJobDetailInput) error {
	job, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(job)
	return err
}

func (s *service) GetJobs(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
