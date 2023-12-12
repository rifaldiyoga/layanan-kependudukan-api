package pengajuan

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetPengajuanByID(ID int) (Pengajuan, error)
	GetPengajuan(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetPengajuanUser(pagination helper.Pagination, user user.User) (helper.Pagination, error)
	CreatePengajuan(input CreatePengajuanInput, user user.User) (Pengajuan, error)
	UpdatePengajuan(ID GetPengajuanDetailInput, input CreatePengajuanInput, user user.User) (Pengajuan, error)
	DeletePengajuan(ID GetPengajuanDetailInput) error
	GetCountPengajuan() (int64, error)
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetPengajuanByID(ID int) (Pengajuan, error) {
	Pengajuan, err := s.repository.FindByID(ID)
	if err != nil {
		return Pengajuan, err
	}

	return Pengajuan, nil
}

func (s *service) CreatePengajuan(input CreatePengajuanInput, user user.User) (Pengajuan, error) {
	Pengajuan := Pengajuan{}

	Pengajuan.Keterangan = input.Keterangan
	Pengajuan.Name = user.Name
	Pengajuan.Code = input.Code
	Pengajuan.Layanan = input.Layanan
	Pengajuan.LayananID = input.LayananID
	Pengajuan.CreatedBy = user.ID
	Pengajuan.RefID = input.RefID
	Pengajuan.CreatedAt = time.Now()
	Pengajuan.UpdatedAt = time.Now()
	Pengajuan.Status = input.Status
	Pengajuan.NIK = user.Nik

	Pengajuan, err := s.repository.Save(Pengajuan)
	if err != nil {
		return Pengajuan, err
	}
	newPengajuan, err := s.repository.FindByID(Pengajuan.ID)
	if err != nil {
		return newPengajuan, err
	}
	return newPengajuan, err
}

func (s *service) UpdatePengajuan(inputDetail GetPengajuanDetailInput, input CreatePengajuanInput, user user.User) (Pengajuan, error) {
	lastPengajuan, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return lastPengajuan, err
	}

	Pengajuan := Pengajuan{}
	Pengajuan.ID = inputDetail.ID
	Pengajuan.Keterangan = input.Keterangan
	Pengajuan.Name = lastPengajuan.Name

	Pengajuan.Code = input.Code
	Pengajuan.Layanan = lastPengajuan.Layanan
	Pengajuan.LayananID = lastPengajuan.LayananID
	Pengajuan.RefID = input.RefID
	Pengajuan.CreatedBy = lastPengajuan.CreatedBy
	Pengajuan.CreatedAt = lastPengajuan.CreatedAt
	Pengajuan.UpdatedAt = time.Now()
	Pengajuan.Status = input.Status
	Pengajuan.Note = input.Note
	Pengajuan.NIK = lastPengajuan.NIK

	if input.Status == "VALID" {
		s.repository.UpdateStatus(lastPengajuan)
	}

	newPengajuan, err := s.repository.Update(Pengajuan)

	return newPengajuan, err
}

func (s *service) DeletePengajuan(inputDetail GetPengajuanDetailInput) error {
	Pengajuan, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(Pengajuan)
	return err
}

func (s *service) GetPengajuanUser(pagination helper.Pagination, user user.User) (helper.Pagination, error) {
	pagination, err := s.repository.FindByUser(pagination, user)

	return pagination, err
}

func (s *service) GetPengajuan(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}

func (s *service) GetCountPengajuan() (int64, error) {
	count, err := s.repository.CountAll()
	return count, err
}
