package kematian

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetKematianByID(ID int) (Kematian, error)
	GetKematians(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastKematian() (Kematian, error)
	CreateKematian(input CreateKematianInput, layanan layanan.Layanan, user user.User) (Kematian, error)
	UpdateKematian(ID GetKematianDetailInput, input CreateKematianInput) (Kematian, error)
	UpdateStatus(ID int) (Kematian, error)
	DeleteKematian(ID GetKematianDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetKematianByID(ID int) (Kematian, error) {
	kematian, err := s.repository.FindByID(ID)

	return kematian, err
}

func (s *service) GetLastKematian() (Kematian, error) {
	kematian, err := s.repository.FindLast()

	return kematian, err
}

func (s *service) CreateKematian(input CreateKematianInput, layanan layanan.Layanan, user user.User) (Kematian, error) {
	domisili := Kematian{}

	// lastKematian, _ := s.repository.FindLast()

	domisili.Keterangan = input.Keterangan
	// domisili.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastKematian.KodeSurat)
	domisili.NIK = input.NIK
	domisili.TglKematian = helper.FormatStringToDate(input.TglKematian)
	domisili.Jam = input.Jam
	domisili.Sebab = input.Sebab
	domisili.Tempat = input.Tempat
	domisili.Saksi1 = input.Saksi1
	domisili.Saksi2 = input.Saksi2
	domisili.NikJenazah = input.NikJenazah
	domisili.LampiranKetRs = input.LampiranKetRs
	domisili.Status = false
	domisili.CreatedAt = time.Now()
	domisili.CreatedBy = user.ID

	newKematian, err := s.repository.Save(domisili)

	return newKematian, err
}

func (s *service) UpdateKematian(inputDetail GetKematianDetailInput, input CreateKematianInput) (Kematian, error) {
	kematian := Kematian{}

	kematian.Keterangan = input.Keterangan
	kematian.KodeSurat = input.KodeSurat
	kematian.NIK = input.NIK
	kematian.CreatedAt = time.Now()

	newKematian, err := s.repository.Update(kematian)
	return newKematian, err
}

func (s *service) UpdateStatus(ID int) (Kematian, error) {
	kematian := Kematian{}
	lastKematian, err := s.repository.FindByID(ID)
	if err != nil {
		return kematian, err
	}

	kematian.Keterangan = lastKematian.Keterangan
	kematian.KodeSurat = lastKematian.KodeSurat
	kematian.NIK = lastKematian.NIK
	kematian.Status = true
	kematian.CreatedAt = time.Now()

	newKematian, err := s.repository.Update(kematian)
	return newKematian, err
}

func (s *service) DeleteKematian(inputDetail GetKematianDetailInput) error {
	kematian, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}
	kematian.ID = inputDetail.ID

	err := s.repository.Delete(kematian)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetKematians(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
