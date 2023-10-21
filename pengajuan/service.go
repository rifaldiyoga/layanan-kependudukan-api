package pengajuan

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/user"
	"time"
)

type Service interface {
	GetPengajuanByID(ID int) (Pengajuan, error)
	GetPengajuans(pagination helper.Pagination, user user.User) (helper.Pagination, error)
	CreatePengajuan(input CreatePengajuanInput, user user.User) (Pengajuan, error)
	UpdatePengajuan(ID GetPengajuanDetailInput, input CreatePengajuanInput, user user.User) (Pengajuan, error)
	DeletePengajuan(ID GetPengajuanDetailInput) error
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
	Pengajuan.Layanan = input.Layanan
	Pengajuan.LayananID = input.LayananID
	Pengajuan.CreatedBy = user.ID
	Pengajuan.CreatedAt = time.Now()
	Pengajuan.UpdatedAt = time.Now()
	if input.Status == "" {
		Pengajuan.Status = "PENDING"
	} else {
		Pengajuan.Status = input.Status + "_" + user.Role
	}
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
	Pengajuan.Layanan = lastPengajuan.Layanan
	Pengajuan.LayananID = lastPengajuan.LayananID
	Pengajuan.CreatedBy = lastPengajuan.CreatedBy
	Pengajuan.CreatedAt = lastPengajuan.CreatedAt
	Pengajuan.UpdatedAt = time.Now()
	if input.Status == "" {
		Pengajuan.Status = "PENDING"
	} else {
		if user.Role != "PENDUDUK" {
			Pengajuan.Status = input.Status + "_" + user.Role
		}
	}
	Pengajuan.NIK = lastPengajuan.NIK

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

func (s *service) GetPengajuans(pagination helper.Pagination, user user.User) (helper.Pagination, error) {
	pagination, err := s.repository.FindByUser(pagination, user)

	return pagination, err
}
