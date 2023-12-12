package berpergian_detail

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/user"
	"time"
)

type Service interface {
	GetBerpergianDetailByID(ID int) (BerpergianDetail, error)
	GetBerpergianDetails(pagination helper.Pagination) (helper.Pagination, error)
	GetLastBerpergianDetail() (BerpergianDetail, error)
	CreateBerpergianDetail(input CreateBerpergianDetailInput, user user.User) (BerpergianDetail, error)
	UpdateBerpergianDetail(ID GetBerpergianDetailDetailInput, input CreateBerpergianDetailInput) (BerpergianDetail, error)
	UpdateStatus(ID int) (BerpergianDetail, error)
	DeleteBerpergianDetail(ID GetBerpergianDetailDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetBerpergianDetailByID(ID int) (BerpergianDetail, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastBerpergianDetail() (BerpergianDetail, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreateBerpergianDetail(input CreateBerpergianDetailInput, user user.User) (BerpergianDetail, error) {
	berpergian := BerpergianDetail{}

	berpergian.NIK = input.NIK
	berpergian.Nama = input.Nama
	berpergian.BerpergianID = input.BerpergianID
	berpergian.StatusFamily = input.StatusFamily
	berpergian.CreatedAt = time.Now()
	berpergian.CreatedBy = user.ID

	newBerpergianDetail, err := s.repository.Save(berpergian)

	return newBerpergianDetail, err
}

func (s *service) UpdateBerpergianDetail(inputDetail GetBerpergianDetailDetailInput, input CreateBerpergianDetailInput) (BerpergianDetail, error) {
	sktm := BerpergianDetail{}

	// sktm.Keterangan = input.Keterangan
	// sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newBerpergianDetail, err := s.repository.Update(sktm)
	return newBerpergianDetail, err
}

func (s *service) UpdateStatus(ID int) (BerpergianDetail, error) {
	sktm := BerpergianDetail{}
	// // lastBerpergianDetail, err := s.repository.FindByID(ID)
	// if err != nil {
	// 	return sktm, err
	// }

	// sktm.Keterangan = lastBerpergianDetail.Keterangan
	// sktm.KodeSurat = lastBerpergianDetail.KodeSurat
	// sktm.NIK = lastBerpergianDetail.NIK
	// sktm.Status = true
	sktm.CreatedAt = time.Now()

	newBerpergianDetail, err := s.repository.Update(sktm)
	return newBerpergianDetail, err
}

func (s *service) DeleteBerpergianDetail(inputDetail GetBerpergianDetailDetailInput) error {
	sktm, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}
	sktm.ID = inputDetail.ID

	err := s.repository.Delete(sktm)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetBerpergianDetails(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
