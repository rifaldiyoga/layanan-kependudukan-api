package sktm

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetSKTMByID(ID int) (SKTM, error)
	GetSKTMs(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastSKTM() (SKTM, error)
	CreateSKTM(input CreateSKTMInput, layanan layanan.Layanan, user user.User) (SKTM, error)
	UpdateSKTM(ID GetSKTMDetailInput, input CreateSKTMInput) (SKTM, error)
	UpdateStatus(ID int) (SKTM, error)
	DeleteSKTM(ID GetSKTMDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetSKTMByID(ID int) (SKTM, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastSKTM() (SKTM, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreateSKTM(input CreateSKTMInput, layanan layanan.Layanan, user user.User) (SKTM, error) {
	sktm := SKTM{}

	// lastSKTM, _ := s.repository.FindLast()

	sktm.Keterangan = input.Keterangan
	// sktm.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastSKTM.KodeSurat)
	sktm.NIK = input.NIK
	sktm.Status = false
	sktm.CreatedAt = time.Now()
	sktm.CreatedBy = user.ID

	newSKTM, err := s.repository.Save(sktm)

	return newSKTM, err
}

func (s *service) UpdateSKTM(inputDetail GetSKTMDetailInput, input CreateSKTMInput) (SKTM, error) {
	sktm := SKTM{}

	sktm.Keterangan = input.Keterangan
	sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newSKTM, err := s.repository.Update(sktm)
	return newSKTM, err
}

func (s *service) UpdateStatus(ID int) (SKTM, error) {
	sktm := SKTM{}
	lastSKTM, err := s.repository.FindByID(ID)
	if err != nil {
		return sktm, err
	}

	sktm.Keterangan = lastSKTM.Keterangan
	sktm.KodeSurat = lastSKTM.KodeSurat
	sktm.NIK = lastSKTM.NIK
	sktm.Status = true
	sktm.CreatedAt = time.Now()

	newSKTM, err := s.repository.Update(sktm)
	return newSKTM, err
}

func (s *service) DeleteSKTM(inputDetail GetSKTMDetailInput) error {
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

func (s *service) GetSKTMs(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
