package pindah_detail

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/user"
	"time"
)

type Service interface {
	GetPindahDetailByID(ID int) (PindahDetail, error)
	GetPindahDetails(pagination helper.Pagination) (helper.Pagination, error)
	GetLastPindahDetail() (PindahDetail, error)
	CreatePindahDetail(input CreatePindahDetailInput, user user.User) (PindahDetail, error)
	UpdatePindahDetail(ID GetPindahDetailDetailInput, input CreatePindahDetailInput) (PindahDetail, error)
	UpdateStatus(ID int) (PindahDetail, error)
	DeletePindahDetail(ID GetPindahDetailDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetPindahDetailByID(ID int) (PindahDetail, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastPindahDetail() (PindahDetail, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreatePindahDetail(input CreatePindahDetailInput, user user.User) (PindahDetail, error) {
	pindah := PindahDetail{}

	pindah.NIK = input.NIK
	pindah.Nama = input.Nama
	pindah.PindahID = input.PindahID
	pindah.StatusFamily = input.StatusFamily
	pindah.CreatedAt = time.Now()
	pindah.CreatedBy = user.ID

	newPindahDetail, err := s.repository.Save(pindah)

	return newPindahDetail, err
}

func (s *service) UpdatePindahDetail(inputDetail GetPindahDetailDetailInput, input CreatePindahDetailInput) (PindahDetail, error) {
	sktm := PindahDetail{}

	// sktm.Keterangan = input.Keterangan
	// sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newPindahDetail, err := s.repository.Update(sktm)
	return newPindahDetail, err
}

func (s *service) UpdateStatus(ID int) (PindahDetail, error) {
	sktm := PindahDetail{}
	// // lastPindahDetail, err := s.repository.FindByID(ID)
	// if err != nil {
	// 	return sktm, err
	// }

	// sktm.Keterangan = lastPindahDetail.Keterangan
	// sktm.KodeSurat = lastPindahDetail.KodeSurat
	// sktm.NIK = lastPindahDetail.NIK
	// sktm.Status = true
	sktm.CreatedAt = time.Now()

	newPindahDetail, err := s.repository.Update(sktm)
	return newPindahDetail, err
}

func (s *service) DeletePindahDetail(inputDetail GetPindahDetailDetailInput) error {
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

func (s *service) GetPindahDetails(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
