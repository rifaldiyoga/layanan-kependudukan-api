package pindah

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetPindahByID(ID int) (Pindah, error)
	GetPindahs(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastPindah() (Pindah, error)
	CreatePindah(input CreatePindahInput, layanan layanan.Layanan, user user.User) (Pindah, error)
	UpdatePindah(ID GetPindahDetailInput, input CreatePindahInput) (Pindah, error)
	UpdateStatus(ID int) (Pindah, error)
	DeletePindah(ID GetPindahDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetPindahByID(ID int) (Pindah, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastPindah() (Pindah, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreatePindah(input CreatePindahInput, layanan layanan.Layanan, user user.User) (Pindah, error) {
	pindah := Pindah{}

	// lastPindah, _ := s.repository.FindLast()

	pindah.NIK = input.NIK
	pindah.NikKepalaKeluarga = input.NikKepalaKeluarga
	// pindah.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastPindah.KodeSurat)
	pindah.Status = false
	pindah.Type = input.Type
	pindah.AlasanPindah = input.AlasanPindah
	pindah.AlamatTujuan = input.AlamatTujuan
	pindah.Rt = input.Rt
	pindah.Rw = input.Rw
	pindah.Kelurahan = input.Kelurahan
	pindah.KecamatanID = input.KecamatanID
	pindah.KotaID = input.KotaID
	pindah.ProvinsiID = input.ProvinsiID
	pindah.KodePos = input.KodePos
	pindah.Telepon = input.Telepon
	pindah.JenisKepindahan = input.JenisKepindahan
	pindah.StatusTidakPindah = input.StatusTidakPindah
	pindah.StatusPindah = input.StatusPindah
	pindah.Lampiran = input.Lampiran
	pindah.CreatedAt = time.Now()
	pindah.CreatedBy = user.ID

	newPindah, err := s.repository.Save(pindah)

	return newPindah, err
}

func (s *service) UpdatePindah(inputDetail GetPindahDetailInput, input CreatePindahInput) (Pindah, error) {
	sktm := Pindah{}

	// sktm.Keterangan = input.Keterangan
	sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newPindah, err := s.repository.Update(sktm)
	return newPindah, err
}

func (s *service) UpdateStatus(ID int) (Pindah, error) {
	sktm := Pindah{}
	lastPindah, err := s.repository.FindByID(ID)
	if err != nil {
		return sktm, err
	}

	// sktm.Keterangan = lastPindah.Keterangan
	sktm.KodeSurat = lastPindah.KodeSurat
	sktm.NIK = lastPindah.NIK
	sktm.Status = true
	sktm.CreatedAt = time.Now()

	newPindah, err := s.repository.Update(sktm)
	return newPindah, err
}

func (s *service) DeletePindah(inputDetail GetPindahDetailInput) error {
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

func (s *service) GetPindahs(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
