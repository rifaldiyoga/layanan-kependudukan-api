package tanah

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetTanahByID(ID int) (Tanah, error)
	GetTanahs(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastTanah() (Tanah, error)
	CreateTanah(input CreateTanahInput, layanan layanan.Layanan, user user.User) (Tanah, error)
	UpdateTanah(ID GetTanahDetailInput, input CreateTanahInput) (Tanah, error)
	UpdateStatus(ID int) (Tanah, error)
	DeleteTanah(ID GetTanahDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetTanahByID(ID int) (Tanah, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastTanah() (Tanah, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreateTanah(input CreateTanahInput, layanan layanan.Layanan, user user.User) (Tanah, error) {
	tanah := Tanah{}

	// lastTanah, _ := s.repository.FindLast()

	tanah.Keterangan = input.Keterangan
	// tanah.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastTanah.KodeSurat)
	tanah.NIK = input.NIK
	tanah.Status = false
	tanah.Lampiran = input.Lampiran
	tanah.Lokasi = input.Lokasi
	tanah.LuasTanah = input.LuasTanah
	tanah.Panjang = 0
	tanah.Lebar = 0
	tanah.BatasBarat = input.BatasBarat
	tanah.BatasTimur = input.BatasTimur
	tanah.BatasUtara = input.BatasUtara
	tanah.BatasSelatan = input.BatasSelatan
	tanah.Saksi1 = input.Saksi1
	tanah.Saksi2 = input.Saksi2
	tanah.Type = input.Type
	tanah.CreatedAt = time.Now()
	tanah.CreatedBy = user.ID

	newTanah, err := s.repository.Save(tanah)

	return newTanah, err
}

func (s *service) UpdateTanah(inputDetail GetTanahDetailInput, input CreateTanahInput) (Tanah, error) {
	sktm := Tanah{}

	sktm.Keterangan = input.Keterangan
	sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newTanah, err := s.repository.Update(sktm)
	return newTanah, err
}

func (s *service) UpdateStatus(ID int) (Tanah, error) {
	sktm := Tanah{}
	lastTanah, err := s.repository.FindByID(ID)
	if err != nil {
		return sktm, err
	}

	sktm.Keterangan = lastTanah.Keterangan
	sktm.KodeSurat = lastTanah.KodeSurat
	sktm.NIK = lastTanah.NIK
	sktm.Status = true
	sktm.CreatedAt = time.Now()

	newTanah, err := s.repository.Update(sktm)
	return newTanah, err
}

func (s *service) DeleteTanah(inputDetail GetTanahDetailInput) error {
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

func (s *service) GetTanahs(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
