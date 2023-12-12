package rumah

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetRumahByID(ID int) (Rumah, error)
	GetRumahs(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastRumah() (Rumah, error)
	CreateRumah(input CreateRumahInput, layanan layanan.Layanan, user user.User) (Rumah, error)
	UpdateRumah(ID GetRumahDetailInput, input CreateRumahInput) (Rumah, error)
	UpdateStatus(ID int) (Rumah, error)
	DeleteRumah(ID GetRumahDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetRumahByID(ID int) (Rumah, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastRumah() (Rumah, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreateRumah(input CreateRumahInput, layanan layanan.Layanan, user user.User) (Rumah, error) {
	rumah := Rumah{}

	// lastRumah, _ := s.repository.FindLast()

	rumah.Keterangan = input.Keterangan
	// rumah.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastRumah.KodeSurat)
	rumah.NIK = input.NIK
	rumah.Status = false
	rumah.Lampiran = input.Lampiran
	rumah.CreatedAt = time.Now()
	rumah.CreatedBy = user.ID

	newRumah, err := s.repository.Save(rumah)

	return newRumah, err
}

func (s *service) UpdateRumah(inputDetail GetRumahDetailInput, input CreateRumahInput) (Rumah, error) {
	sktm := Rumah{}

	sktm.Keterangan = input.Keterangan
	sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newRumah, err := s.repository.Update(sktm)
	return newRumah, err
}

func (s *service) UpdateStatus(ID int) (Rumah, error) {
	sktm := Rumah{}
	lastRumah, err := s.repository.FindByID(ID)
	if err != nil {
		return sktm, err
	}

	sktm.Keterangan = lastRumah.Keterangan
	sktm.KodeSurat = lastRumah.KodeSurat
	sktm.NIK = lastRumah.NIK
	sktm.Status = true
	sktm.CreatedAt = time.Now()

	newRumah, err := s.repository.Update(sktm)
	return newRumah, err
}

func (s *service) DeleteRumah(inputDetail GetRumahDetailInput) error {
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

func (s *service) GetRumahs(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
