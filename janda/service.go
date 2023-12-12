package janda

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetJandaByID(ID int) (Janda, error)
	GetJandas(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastJanda() (Janda, error)
	CreateJanda(input CreateJandaInput, layanan layanan.Layanan, user user.User) (Janda, error)
	UpdateJanda(ID GetJandaDetailInput, input CreateJandaInput) (Janda, error)
	UpdateStatus(ID int) (Janda, error)
	DeleteJanda(ID GetJandaDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetJandaByID(ID int) (Janda, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastJanda() (Janda, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreateJanda(input CreateJandaInput, layanan layanan.Layanan, user user.User) (Janda, error) {
	janda := Janda{}

	// lastJanda, _ := s.repository.FindLast()

	janda.Keterangan = input.Keterangan
	// janda.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastJanda.KodeSurat)
	janda.NIK = input.NIK
	janda.Status = false
	janda.Lampiran = input.Lampiran
	janda.CreatedAt = time.Now()
	janda.CreatedBy = user.ID

	newJanda, err := s.repository.Save(janda)

	return newJanda, err
}

func (s *service) UpdateJanda(inputDetail GetJandaDetailInput, input CreateJandaInput) (Janda, error) {
	sktm := Janda{}

	sktm.Keterangan = input.Keterangan
	sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newJanda, err := s.repository.Update(sktm)
	return newJanda, err
}

func (s *service) UpdateStatus(ID int) (Janda, error) {
	sktm := Janda{}
	lastJanda, err := s.repository.FindByID(ID)
	if err != nil {
		return sktm, err
	}

	sktm.Keterangan = lastJanda.Keterangan
	sktm.KodeSurat = lastJanda.KodeSurat
	sktm.NIK = lastJanda.NIK
	sktm.Status = true
	sktm.CreatedAt = time.Now()

	newJanda, err := s.repository.Update(sktm)
	return newJanda, err
}

func (s *service) DeleteJanda(inputDetail GetJandaDetailInput) error {
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

func (s *service) GetJandas(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
