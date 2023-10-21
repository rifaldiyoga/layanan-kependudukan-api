package detail_pengajuan

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/user"
	"time"
)

type Service interface {
	GetDetailPengajuanByID(ID int) (DetailPengajuan, error)
	GetDetailByPengajuan(ID int) ([]DetailPengajuan, error)
	GetDetailPengajuans(pagination helper.Pagination) (helper.Pagination, error)
	CreateDetailPengajuan(pengajuanID int, status string, user user.User) (DetailPengajuan, error)
	UpdateDetailPengajuan(ID GetDetailPengajuanDetailInput, input CreateDetailPengajuanInput) (DetailPengajuan, error)
	DeleteDetailPengajuan(ID GetDetailPengajuanDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetDetailPengajuanByID(ID int) (DetailPengajuan, error) {
	DetailPengajuan, err := s.repository.FindByID(ID)
	if err != nil {
		return DetailPengajuan, err
	}

	return DetailPengajuan, nil
}

func (s *service) GetDetailByPengajuan(ID int) ([]DetailPengajuan, error) {
	DetailPengajuan, err := s.repository.FindByPengajuan(ID)
	if err != nil {
		return DetailPengajuan, err
	}

	return DetailPengajuan, nil
}

func (s *service) CreateDetailPengajuan(pengajuanID int, status string, user user.User) (DetailPengajuan, error) {
	DetailPengajuan := DetailPengajuan{}

	DetailPengajuan.CreatedBy = user.ID

	if status == "" {
		DetailPengajuan.Status = "PENDING"
	} else {
		if user.Role != "PENDUDUK" {
			DetailPengajuan.Status = status + "_" + user.Role
		} else {
			DetailPengajuan.Status = status
		}
	}
	DetailPengajuan.Name = user.Name
	DetailPengajuan.PengajuanID = pengajuanID
	DetailPengajuan.CreatedAt = time.Now()

	newDetailPengajuan, err := s.repository.Save(DetailPengajuan)
	return newDetailPengajuan, err
}

func (s *service) UpdateDetailPengajuan(inputDetail GetDetailPengajuanDetailInput, input CreateDetailPengajuanInput) (DetailPengajuan, error) {
	DetailPengajuan, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return DetailPengajuan, err
	}

	// DetailPengajuan.Code = input.Code
	// DetailPengajuan.Name = input.Name
	// DetailPengajuan.UpdatedAt = time.Now()

	newDetailPengajuan, err := s.repository.Update(DetailPengajuan)
	return newDetailPengajuan, err
}

func (s *service) DeleteDetailPengajuan(inputDetail GetDetailPengajuanDetailInput) error {
	DetailPengajuan, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}

	err := s.repository.Delete(DetailPengajuan)
	return err
}

func (s *service) GetDetailPengajuans(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)
	return pagination, err
}
