package pernah_menikah

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetPernahMenikahByID(ID int) (PernahMenikah, error)
	GetPernahMenikahs(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastPernahMenikah() (PernahMenikah, error)
	CreatePernahMenikah(input CreatePernahMenikahInput, layanan layanan.Layanan, user user.User) (PernahMenikah, error)
	UpdatePernahMenikah(ID GetPernahMenikahDetailInput, input CreatePernahMenikahInput) (PernahMenikah, error)
	UpdateStatus(ID int) (PernahMenikah, error)
	DeletePernahMenikah(ID GetPernahMenikahDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetPernahMenikahByID(ID int) (PernahMenikah, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastPernahMenikah() (PernahMenikah, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreatePernahMenikah(input CreatePernahMenikahInput, layanan layanan.Layanan, user user.User) (PernahMenikah, error) {
	janda := PernahMenikah{}

	// lastPernahMenikah, _ := s.repository.FindLast()

	janda.Keterangan = input.Keterangan
	// janda.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastPernahMenikah.KodeSurat)
	janda.NIK = input.NIK
	janda.Status = false
	janda.Lampiran = input.Lampiran
	janda.CreatedAt = time.Now()
	janda.CreatedBy = user.ID
	janda.NikSuami = input.NIKSuami
	janda.NikIstri = input.NIKIstri
	newPernahMenikah, err := s.repository.Save(janda)

	return newPernahMenikah, err
}

func (s *service) UpdatePernahMenikah(inputDetail GetPernahMenikahDetailInput, input CreatePernahMenikahInput) (PernahMenikah, error) {
	sktm := PernahMenikah{}

	sktm.Keterangan = input.Keterangan
	sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newPernahMenikah, err := s.repository.Update(sktm)
	return newPernahMenikah, err
}

func (s *service) UpdateStatus(ID int) (PernahMenikah, error) {
	sktm := PernahMenikah{}
	lastPernahMenikah, err := s.repository.FindByID(ID)
	if err != nil {
		return sktm, err
	}

	sktm.Keterangan = lastPernahMenikah.Keterangan
	sktm.KodeSurat = lastPernahMenikah.KodeSurat
	sktm.NIK = lastPernahMenikah.NIK
	sktm.Status = true
	sktm.CreatedAt = time.Now()

	newPernahMenikah, err := s.repository.Update(sktm)
	return newPernahMenikah, err
}

func (s *service) DeletePernahMenikah(inputDetail GetPernahMenikahDetailInput) error {
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

func (s *service) GetPernahMenikahs(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
