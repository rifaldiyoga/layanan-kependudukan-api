package belum_menikah

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetBelumMenikahByID(ID int) (BelumMenikah, error)
	GetBelumMenikahs(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastBelumMenikah() (BelumMenikah, error)
	CreateBelumMenikah(input CreateBelumMenikahInput, layanan layanan.Layanan, user user.User) (BelumMenikah, error)
	UpdateBelumMenikah(ID GetBelumMenikahDetailInput, input CreateBelumMenikahInput) (BelumMenikah, error)
	UpdateStatus(ID int) (BelumMenikah, error)
	DeleteBelumMenikah(ID GetBelumMenikahDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetBelumMenikahByID(ID int) (BelumMenikah, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastBelumMenikah() (BelumMenikah, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreateBelumMenikah(input CreateBelumMenikahInput, layanan layanan.Layanan, user user.User) (BelumMenikah, error) {
	janda := BelumMenikah{}

	// lastBelumMenikah, _ := s.repository.FindLast()

	janda.Keterangan = input.Keterangan
	// janda.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastBelumMenikah.KodeSurat)
	janda.NIK = input.NIK
	janda.Status = false
	janda.Lampiran = input.Lampiran
	janda.CreatedAt = time.Now()
	janda.CreatedBy = user.ID

	newBelumMenikah, err := s.repository.Save(janda)

	return newBelumMenikah, err
}

func (s *service) UpdateBelumMenikah(inputDetail GetBelumMenikahDetailInput, input CreateBelumMenikahInput) (BelumMenikah, error) {
	sktm := BelumMenikah{}

	sktm.Keterangan = input.Keterangan
	sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newBelumMenikah, err := s.repository.Update(sktm)
	return newBelumMenikah, err
}

func (s *service) UpdateStatus(ID int) (BelumMenikah, error) {
	sktm := BelumMenikah{}
	lastBelumMenikah, err := s.repository.FindByID(ID)
	if err != nil {
		return sktm, err
	}

	sktm.Keterangan = lastBelumMenikah.Keterangan
	sktm.KodeSurat = lastBelumMenikah.KodeSurat
	sktm.NIK = lastBelumMenikah.NIK
	sktm.Status = true
	sktm.CreatedAt = time.Now()

	newBelumMenikah, err := s.repository.Update(sktm)
	return newBelumMenikah, err
}

func (s *service) DeleteBelumMenikah(inputDetail GetBelumMenikahDetailInput) error {
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

func (s *service) GetBelumMenikahs(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
