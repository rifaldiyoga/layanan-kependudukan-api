package sporadik

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetSporadikByID(ID int) (Sporadik, error)
	GetSporadiks(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastSporadik() (Sporadik, error)
	CreateSporadik(input CreateSporadikInput, layanan layanan.Layanan, user user.User) (Sporadik, error)
	UpdateSporadik(ID GetSporadikDetailInput, input CreateSporadikInput) (Sporadik, error)
	UpdateStatus(ID int) (Sporadik, error)
	DeleteSporadik(ID GetSporadikDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetSporadikByID(ID int) (Sporadik, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastSporadik() (Sporadik, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreateSporadik(input CreateSporadikInput, layanan layanan.Layanan, user user.User) (Sporadik, error) {
	sporadik := Sporadik{}

	// lastSporadik, _ := s.repository.FindLast()

	sporadik.Keterangan = input.Keterangan
	// sporadik.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastSporadik.KodeSurat)
	sporadik.NIK = input.NIK
	sporadik.Status = false
	sporadik.LampiranBukti = input.LampiranBukti
	sporadik.LampiranPemohon = input.LampiranPemohon
	sporadik.LampiranSporadikBaru = input.LampiranSporadikBaru
	sporadik.LampiranSporadikLama = input.LampiranSporadikLama
	sporadik.LampiranLunasPbb = input.LampiranLunasPbb
	sporadik.CreatedAt = time.Now()
	sporadik.CreatedBy = user.ID

	newSporadik, err := s.repository.Save(sporadik)

	return newSporadik, err
}

func (s *service) UpdateSporadik(inputDetail GetSporadikDetailInput, input CreateSporadikInput) (Sporadik, error) {
	sktm := Sporadik{}

	sktm.Keterangan = input.Keterangan
	sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newSporadik, err := s.repository.Update(sktm)
	return newSporadik, err
}

func (s *service) UpdateStatus(ID int) (Sporadik, error) {
	sktm := Sporadik{}
	lastSporadik, err := s.repository.FindByID(ID)
	if err != nil {
		return sktm, err
	}

	sktm.Keterangan = lastSporadik.Keterangan
	sktm.KodeSurat = lastSporadik.KodeSurat
	sktm.NIK = lastSporadik.NIK
	sktm.Status = true
	sktm.CreatedAt = time.Now()

	newSporadik, err := s.repository.Update(sktm)
	return newSporadik, err
}

func (s *service) DeleteSporadik(inputDetail GetSporadikDetailInput) error {
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

func (s *service) GetSporadiks(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
