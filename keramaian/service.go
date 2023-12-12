package keramaian

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetKeramaianByID(ID int) (Keramaian, error)
	GetKeramaians(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastKeramaian() (Keramaian, error)
	CreateKeramaian(input CreateKeramaianInput, layanan layanan.Layanan, user user.User) (Keramaian, error)
	UpdateKeramaian(ID GetKeramaianDetailInput, input CreateKeramaianInput) (Keramaian, error)
	UpdateStatus(ID int) (Keramaian, error)
	DeleteKeramaian(ID GetKeramaianDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetKeramaianByID(ID int) (Keramaian, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastKeramaian() (Keramaian, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreateKeramaian(input CreateKeramaianInput, layanan layanan.Layanan, user user.User) (Keramaian, error) {
	keramaian := Keramaian{}

	// lastKeramaian, _ := s.repository.FindLast()

	keramaian.Keterangan = input.Keterangan
	// keramaian.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastKeramaian.KodeSurat)
	keramaian.NIK = input.NIK
	keramaian.Status = false
	keramaian.NamaAcara = input.NamaAcara
	keramaian.Tanggal = input.Tanggal
	keramaian.Waktu = input.Waktu
	keramaian.Tempat = input.Tempat
	keramaian.Telpon = input.Telpon
	keramaian.Alamat = input.Alamat
	keramaian.Lampiran = input.Lampiran
	keramaian.CreatedAt = time.Now()
	keramaian.CreatedBy = user.ID

	newKeramaian, err := s.repository.Save(keramaian)

	return newKeramaian, err
}

func (s *service) UpdateKeramaian(inputDetail GetKeramaianDetailInput, input CreateKeramaianInput) (Keramaian, error) {
	sktm := Keramaian{}

	sktm.Keterangan = input.Keterangan
	sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newKeramaian, err := s.repository.Update(sktm)
	return newKeramaian, err
}

func (s *service) UpdateStatus(ID int) (Keramaian, error) {
	sktm := Keramaian{}
	lastKeramaian, err := s.repository.FindByID(ID)
	if err != nil {
		return sktm, err
	}

	sktm.Keterangan = lastKeramaian.Keterangan
	sktm.KodeSurat = lastKeramaian.KodeSurat
	sktm.NIK = lastKeramaian.NIK
	sktm.Status = true
	sktm.CreatedAt = time.Now()

	newKeramaian, err := s.repository.Update(sktm)
	return newKeramaian, err
}

func (s *service) DeleteKeramaian(inputDetail GetKeramaianDetailInput) error {
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

func (s *service) GetKeramaians(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
