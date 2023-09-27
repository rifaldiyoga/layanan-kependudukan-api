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
	UpdatePengajuan(ID GetPengajuanDetailInput, input CreatePengajuanInput) (Pengajuan, error)
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
	Pengajuan.CreatedBy = user.ID
	Pengajuan.Name = user.Name
	Pengajuan.Layanan = input.Layanan
	Pengajuan.LayananID = input.LayananID
	Pengajuan.CreatedAt = time.Now()
	Pengajuan.UpdatedAt = time.Now()

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

func (s *service) UpdatePengajuan(inputDetail GetPengajuanDetailInput, input CreatePengajuanInput) (Pengajuan, error) {
	Pengajuan, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return Pengajuan, err
	}

	// Pengajuan.Code = input.Code
	// Pengajuan.Name = input.Name
	Pengajuan.UpdatedAt = time.Now()

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
