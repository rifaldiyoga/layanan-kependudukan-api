package kelahiran

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetKelahiranByID(ID int) (Kelahiran, error)
	GetKelahirans(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastKelahiran() (Kelahiran, error)
	CreateKelahiran(input CreateKelahiranInput, layanan layanan.Layanan, user user.User) (Kelahiran, error)
	UpdateKelahiran(ID GetKelahiranDetailInput, input CreateKelahiranInput) (Kelahiran, error)
	UpdateStatus(ID int) (Kelahiran, error)
	DeleteKelahiran(ID GetKelahiranDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetKelahiranByID(ID int) (Kelahiran, error) {
	kelahiran, err := s.repository.FindByID(ID)

	return kelahiran, err
}

func (s *service) GetLastKelahiran() (Kelahiran, error) {
	kelahiran, err := s.repository.FindLast()

	return kelahiran, err
}

func (s *service) CreateKelahiran(input CreateKelahiranInput, layanan layanan.Layanan, user user.User) (Kelahiran, error) {
	domisili := Kelahiran{}

	// lastKelahiran, _ := s.repository.FindLast()

	domisili.Keterangan = input.Keterangan
	domisili.NIK = user.Nik
	// domisili.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastKelahiran.KodeSurat)
	domisili.Nama = input.Nama
	domisili.BirthDate = helper.FormatStringToDate(input.BirthDate)
	domisili.Jam = input.Jam
	domisili.BirthPlace = input.BirthPlace
	domisili.NikAyah = input.NikAyah
	domisili.NikIbu = input.NikIbu
	domisili.JK = input.JK
	domisili.ProvinsiID = input.ProvinsiID
	domisili.KotaID = input.KotaID
	domisili.KecamatanID = input.KecamatanID
	domisili.LampiranBukuNikah = input.LampiranBukuNikah
	domisili.LampiranKetRs = input.LampiranKetRs
	domisili.Status = false
	domisili.CreatedAt = time.Now()
	domisili.CreatedBy = user.ID

	newKelahiran, err := s.repository.Save(domisili)

	return newKelahiran, err
}

func (s *service) UpdateKelahiran(inputDetail GetKelahiranDetailInput, input CreateKelahiranInput) (Kelahiran, error) {
	kelahiran := Kelahiran{}

	kelahiran.Keterangan = input.Keterangan
	// kelahiran.KodeSurat = input.KodeSurat
	// kelahiran.NIK = input.NIK
	kelahiran.CreatedAt = time.Now()

	newKelahiran, err := s.repository.Update(kelahiran)
	return newKelahiran, err
}

func (s *service) UpdateStatus(ID int) (Kelahiran, error) {
	kelahiran := Kelahiran{}
	lastKelahiran, err := s.repository.FindByID(ID)
	if err != nil {
		return kelahiran, err
	}

	kelahiran.Keterangan = lastKelahiran.Keterangan
	kelahiran.KodeSurat = lastKelahiran.KodeSurat
	kelahiran.NIK = lastKelahiran.NIK
	kelahiran.Status = true
	kelahiran.CreatedAt = time.Now()

	newKelahiran, err := s.repository.Update(kelahiran)
	return newKelahiran, err
}

func (s *service) DeleteKelahiran(inputDetail GetKelahiranDetailInput) error {
	kelahiran, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}
	kelahiran.ID = inputDetail.ID

	err := s.repository.Delete(kelahiran)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetKelahirans(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
