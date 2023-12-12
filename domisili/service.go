package domisili

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetDomisiliByID(ID int) (Domisili, error)
	GetDomisilis(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastDomisili() (Domisili, error)
	CreateDomisili(input CreateDomisiliInput, layanan layanan.Layanan, user user.User) (Domisili, error)
	UpdateDomisili(ID GetDomisiliDetailInput, input CreateDomisiliInput) (Domisili, error)
	UpdateStatus(ID int) (Domisili, error)
	DeleteDomisili(ID GetDomisiliDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetDomisiliByID(ID int) (Domisili, error) {
	sktm, err := s.repository.FindByID(ID)

	return sktm, err
}

func (s *service) GetLastDomisili() (Domisili, error) {
	sktm, err := s.repository.FindLast()

	return sktm, err
}

func (s *service) CreateDomisili(input CreateDomisiliInput, layanan layanan.Layanan, user user.User) (Domisili, error) {
	domisili := Domisili{}

	// lastDomisili, _ := s.repository.FindLast()

	domisili.Keterangan = input.Keterangan
	// domisili.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastDomisili.KodeSurat)
	domisili.NIK = input.NIK
	domisili.Type = input.Type
	domisili.NamaPerusahaan = input.NamaPerusahaan
	domisili.JenisPerusahaan = input.JenisPerusahaan
	domisili.StatusBangunan = input.StatusBangunan
	domisili.TelpPerusahaan = input.TelpPerusahaan
	domisili.AktaPerusahaan = input.AktaPerusahaan
	domisili.SKPengesahan = input.SKPengesahan
	domisili.PenanggungJawab = input.PenanggungJawab
	domisili.Alamat = input.AlamatPerusahaan
	domisili.Lampiran = input.LampiranPath
	domisili.Status = false
	domisili.CreatedAt = time.Now()
	domisili.CreatedBy = user.ID

	newDomisili, err := s.repository.Save(domisili)

	return newDomisili, err
}

func (s *service) UpdateDomisili(inputDetail GetDomisiliDetailInput, input CreateDomisiliInput) (Domisili, error) {
	sktm := Domisili{}

	sktm.Keterangan = input.Keterangan
	sktm.KodeSurat = input.KodeSurat
	sktm.NIK = input.NIK
	sktm.CreatedAt = time.Now()

	newDomisili, err := s.repository.Update(sktm)
	return newDomisili, err
}

func (s *service) UpdateStatus(ID int) (Domisili, error) {
	sktm := Domisili{}
	lastDomisili, err := s.repository.FindByID(ID)
	if err != nil {
		return sktm, err
	}

	sktm.Keterangan = lastDomisili.Keterangan
	sktm.KodeSurat = lastDomisili.KodeSurat
	sktm.NIK = lastDomisili.NIK
	sktm.Status = true
	sktm.CreatedAt = time.Now()

	newDomisili, err := s.repository.Update(sktm)
	return newDomisili, err
}

func (s *service) DeleteDomisili(inputDetail GetDomisiliDetailInput) error {
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

func (s *service) GetDomisilis(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
