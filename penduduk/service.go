package penduduk

import (
	"layanan-kependudukan-api/helper"
	"strconv"
	"time"
)

type Service interface {
	GetPendudukByID(ID int) (Penduduk, error)
	GetPenduduks(pagination helper.Pagination) (helper.Pagination, error)
	CreatePenduduk(input CreatePendudukInput) (Penduduk, error)
	UpdatePenduduk(ID GetPendudukDetailInput, input CreatePendudukInput) (Penduduk, error)
	DeletePenduduk(ID GetPendudukDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetPendudukByID(ID int) (Penduduk, error) {
	penduduk, err := s.repository.FindByID(ID)

	return penduduk, err
}

func (s *service) CreatePenduduk(input CreatePendudukInput) (Penduduk, error) {
	penduduk := Penduduk{}

	birthDate := helper.FormatStringToDate(input.BirthDate)
	mariedDate := helper.FormatStringToDate(input.MariedDate)

	if input.NIK == "" {
		var counter int
		last, err := s.repository.FindByDate(birthDate)
		if err != nil {
			counter = 1
		} else {
			i, _ := strconv.Atoi(last.NIK[len(last.NIK)-4:])
			counter = i + 1
		}

		penduduk.NIK = helper.GenerateNIK(birthDate, counter)
	} else {
		penduduk.NIK = input.NIK
	}
	penduduk.Fullname = input.FullName
	penduduk.BirthDate = birthDate
	penduduk.BirthPlace = input.BirthPlace
	penduduk.JK = input.JK
	penduduk.ReligionID = input.ReligionID
	penduduk.Nationality = input.Nationality
	penduduk.PekerjaanID = input.JobID
	penduduk.PendidikanID = input.EducationID
	penduduk.Address = input.Address
	penduduk.RtID = input.RtID
	penduduk.RwID = input.RwID
	penduduk.MariedType = input.MariedType
	penduduk.MariedDate = mariedDate
	penduduk.BloodType = input.BloodType
	penduduk.KelurahanID = input.KelurahanID
	penduduk.KecamatanID = input.KecamatanID
	penduduk.CreatedAt = time.Now()
	penduduk.UpdatedAt = time.Now()

	newPenduduk, err := s.repository.Save(penduduk)

	return newPenduduk, err
}

func (s *service) UpdatePenduduk(inputDetail GetPendudukDetailInput, input CreatePendudukInput) (Penduduk, error) {
	penduduk := Penduduk{}

	birthDate := helper.FormatStringToDate(input.BirthDate)
	mariedDate := helper.FormatStringToDate(input.MariedDate)
	penduduk.ID = inputDetail.ID
	penduduk.NIK = input.NIK
	penduduk.Fullname = input.FullName
	penduduk.BirthDate = birthDate
	penduduk.BirthPlace = input.BirthPlace
	penduduk.JK = input.JK
	penduduk.ReligionID = input.ReligionID
	penduduk.Nationality = input.Nationality
	penduduk.PekerjaanID = input.JobID
	penduduk.PendidikanID = input.EducationID
	penduduk.Address = input.Address
	penduduk.RtID = input.RtID
	penduduk.RwID = input.RwID
	penduduk.MariedType = input.MariedType
	penduduk.MariedDate = mariedDate
	penduduk.BloodType = input.BloodType
	penduduk.KelurahanID = input.KelurahanID
	penduduk.KecamatanID = input.KecamatanID
	penduduk.CreatedAt = time.Now()
	penduduk.UpdatedAt = time.Now()

	newPenduduk, err := s.repository.Update(penduduk)
	return newPenduduk, err
}

func (s *service) DeletePenduduk(inputDetail GetPendudukDetailInput) error {
	penduduk, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}
	penduduk.ID = inputDetail.ID

	err := s.repository.Delete(penduduk)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetPenduduks(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}
