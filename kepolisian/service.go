package kepolisian

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetKepolisianByID(ID int) (Kepolisian, error)
	GetKepolisians(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastKepolisian() (Kepolisian, error)
	CreateKepolisian(input CreateKepolisianInput, layanan layanan.Layanan, user user.User) (Kepolisian, error)
	UpdateKepolisian(ID GetKepolisianDetailInput, input CreateKepolisianInput) (Kepolisian, error)
	UpdateStatus(ID int) (Kepolisian, error)
	DeleteKepolisian(ID GetKepolisianDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetKepolisianByID(ID int) (Kepolisian, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastKepolisian() (Kepolisian, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreateKepolisian(input CreateKepolisianInput, layanan layanan.Layanan, user user.User) (Kepolisian, error) {
	kepolisian := Kepolisian{}

	// lastKepolisian, _ := s.repository.FindLast()

	kepolisian.Keterangan = input.Keterangan
	// kepolisian.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastKepolisian.KodeSurat)
	kepolisian.NIK = input.NIK
	kepolisian.Status = false
	kepolisian.Lampiran = input.Lampiran
	kepolisian.CreatedAt = time.Now()
	kepolisian.CreatedBy = user.ID

	newKepolisian, err := s.repository.Save(kepolisian)

	return newKepolisian, err
}

func (s *service) UpdateKepolisian(inputDetail GetKepolisianDetailInput, input CreateKepolisianInput) (Kepolisian, error) {
	sktm := Kepolisian{}

	sktm.Keterangan = input.Keterangan
	sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newKepolisian, err := s.repository.Update(sktm)
	return newKepolisian, err
}

func (s *service) UpdateStatus(ID int) (Kepolisian, error) {
	sktm := Kepolisian{}
	lastKepolisian, err := s.repository.FindByID(ID)
	if err != nil {
		return sktm, err
	}

	sktm.Keterangan = lastKepolisian.Keterangan
	sktm.KodeSurat = lastKepolisian.KodeSurat
	sktm.NIK = lastKepolisian.NIK
	sktm.Status = true
	sktm.CreatedAt = time.Now()

	newKepolisian, err := s.repository.Update(sktm)
	return newKepolisian, err
}

func (s *service) DeleteKepolisian(inputDetail GetKepolisianDetailInput) error {
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

func (s *service) GetKepolisians(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
