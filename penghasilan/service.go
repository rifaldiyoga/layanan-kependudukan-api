package penghasilan

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetPenghasilanByID(ID int) (Penghasilan, error)
	GetPenghasilans(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastPenghasilan() (Penghasilan, error)
	CreatePenghasilan(input CreatePenghasilanInput, layanan layanan.Layanan, user user.User) (Penghasilan, error)
	UpdatePenghasilan(ID GetPenghasilanDetailInput, input CreatePenghasilanInput) (Penghasilan, error)
	UpdateStatus(ID int) (Penghasilan, error)
	DeletePenghasilan(ID GetPenghasilanDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetPenghasilanByID(ID int) (Penghasilan, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastPenghasilan() (Penghasilan, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreatePenghasilan(input CreatePenghasilanInput, layanan layanan.Layanan, user user.User) (Penghasilan, error) {
	penghasilan := Penghasilan{}

	// lastPenghasilan, _ := s.repository.FindLast()

	penghasilan.Keterangan = input.Keterangan
	// penghasilan.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastPenghasilan.KodeSurat)
	penghasilan.NIK = input.NIK
	penghasilan.Status = false
	penghasilan.Lampiran = input.Lampiran
	penghasilan.Penghasilan = input.Penghasilan
	penghasilan.CreatedAt = time.Now()
	penghasilan.CreatedBy = user.ID

	newPenghasilan, err := s.repository.Save(penghasilan)

	return newPenghasilan, err
}

func (s *service) UpdatePenghasilan(inputDetail GetPenghasilanDetailInput, input CreatePenghasilanInput) (Penghasilan, error) {
	sktm := Penghasilan{}

	sktm.Keterangan = input.Keterangan
	sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newPenghasilan, err := s.repository.Update(sktm)
	return newPenghasilan, err
}

func (s *service) UpdateStatus(ID int) (Penghasilan, error) {
	sktm := Penghasilan{}
	lastPenghasilan, err := s.repository.FindByID(ID)
	if err != nil {
		return sktm, err
	}

	sktm.Keterangan = lastPenghasilan.Keterangan
	sktm.KodeSurat = lastPenghasilan.KodeSurat
	sktm.NIK = lastPenghasilan.NIK
	sktm.Status = true
	sktm.CreatedAt = time.Now()

	newPenghasilan, err := s.repository.Update(sktm)
	return newPenghasilan, err
}

func (s *service) DeletePenghasilan(inputDetail GetPenghasilanDetailInput) error {
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

func (s *service) GetPenghasilans(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
