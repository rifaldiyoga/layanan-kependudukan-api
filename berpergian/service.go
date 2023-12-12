package berpergian

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetBerpergianByID(ID int) (Berpergian, error)
	GetBerpergians(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastBerpergian() (Berpergian, error)
	CreateBerpergian(input CreateBerpergianInput, layanan layanan.Layanan, user user.User) (Berpergian, error)
	UpdateBerpergian(ID GetBerpergianDetailInput, input CreateBerpergianInput) (Berpergian, error)
	UpdateStatus(ID int) (Berpergian, error)
	DeleteBerpergian(ID GetBerpergianDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetBerpergianByID(ID int) (Berpergian, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastBerpergian() (Berpergian, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreateBerpergian(input CreateBerpergianInput, layanan layanan.Layanan, user user.User) (Berpergian, error) {
	berpergian := Berpergian{}

	// lastBerpergian, _ := s.repository.FindLast()

	berpergian.Keterangan = input.Keterangan
	// berpergian.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastBerpergian.KodeSurat)
	berpergian.NIK = input.NIK
	berpergian.Status = false
	berpergian.Lampiran = input.Lampiran
	berpergian.Tujuan = input.Tujuan
	berpergian.TglBerangkat = helper.FormatStringToDate(input.TglBerangkat)
	berpergian.TglKembali = helper.FormatStringToDate(input.TglKembali)
	berpergian.CreatedAt = time.Now()
	berpergian.CreatedBy = user.ID

	newBerpergian, err := s.repository.Save(berpergian)

	return newBerpergian, err
}

func (s *service) UpdateBerpergian(inputDetail GetBerpergianDetailInput, input CreateBerpergianInput) (Berpergian, error) {
	sktm := Berpergian{}

	sktm.Keterangan = input.Keterangan
	sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newBerpergian, err := s.repository.Update(sktm)
	return newBerpergian, err
}

func (s *service) UpdateStatus(ID int) (Berpergian, error) {
	sktm := Berpergian{}
	lastBerpergian, err := s.repository.FindByID(ID)
	if err != nil {
		return sktm, err
	}

	sktm.Keterangan = lastBerpergian.Keterangan
	sktm.KodeSurat = lastBerpergian.KodeSurat
	sktm.NIK = lastBerpergian.NIK
	sktm.Status = true
	sktm.CreatedAt = time.Now()

	newBerpergian, err := s.repository.Update(sktm)
	return newBerpergian, err
}

func (s *service) DeleteBerpergian(inputDetail GetBerpergianDetailInput) error {
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

func (s *service) GetBerpergians(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
