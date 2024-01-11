package penduduk

import (
	"layanan-kependudukan-api/helper"
	"strconv"
	"time"
)

type Service interface {
	GetPendudukByID(ID int) (Penduduk, error)
	GetPendudukByNIK(NIK string) (Penduduk, error)
	GetPendudukByNoKK(NIK string) ([]Penduduk, error)
	GetPenduduks(pagination helper.Pagination, NIK string) (helper.Pagination, error)
	GetRTByPengaju(RtID int, RwID int) (Penduduk, error)
	GetRWByPengaju(RwID int) (Penduduk, error)
	CreatePenduduk(input CreatePendudukInput) (Penduduk, error)
	UpdatePenduduk(ID GetPendudukDetailInput, input CreatePendudukInput) (Penduduk, error)
	UpdatePenduduks(ID int, input Penduduk) (Penduduk, error)
	DeletePenduduk(ID GetPendudukDetailInput) error
	GetCountPenduduk() (int64, error)
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

func (s *service) GetPendudukByNIK(ID string) (Penduduk, error) {
	penduduk, err := s.repository.FindByNIK(ID)

	return penduduk, err
}

func (s *service) GetPendudukByNoKK(ID string) ([]Penduduk, error) {
	penduduk, err := s.repository.FindByNoKK(ID)

	return penduduk, err
}

func (s *service) GetRTByPengaju(RtID int, RwID int) (Penduduk, error) {
	penduduk, err := s.repository.FindByRT(RtID, RwID)

	return penduduk, err
}

func (s *service) GetRWByPengaju(RwID int) (Penduduk, error) {
	penduduk, err := s.repository.FindByRW(RwID)

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
	penduduk.KotaID = input.KotaID
	penduduk.ProvinsiID = input.ProvinsiID
	penduduk.NoKK = input.NoKK
	penduduk.StatusFamily = input.StatusFamily
	penduduk.CreatedAt = time.Now()
	penduduk.UpdatedAt = time.Now()
	penduduk.Active = true

	newPenduduk, err := s.repository.Save(penduduk)

	return newPenduduk, err
}

func (s *service) UpdatePenduduk(inputDetail GetPendudukDetailInput, input CreatePendudukInput) (Penduduk, error) {
	penduduk := Penduduk{}

	birthDate := helper.FormatStringToDate(input.BirthDate)
	mariedDate := helper.FormatStringToDate(input.MariedDate)
	penduduk.ID = inputDetail.ID
	penduduk.NIK = input.NIK
	penduduk.NoKK = input.NoKK
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
	penduduk.KotaID = input.KotaID
	penduduk.ProvinsiID = input.ProvinsiID
	penduduk.StatusFamily = input.StatusFamily
	penduduk.CreatedAt = time.Now()
	penduduk.UpdatedAt = time.Now()
	penduduk.Active = true

	newPenduduk, err := s.repository.Update(penduduk)
	return newPenduduk, err
}

func (s *service) UpdatePenduduks(inputDetail int, input Penduduk) (Penduduk, error) {
	penduduk := Penduduk{}

	birthDate := input.BirthDate
	mariedDate := input.MariedDate
	penduduk.ID = inputDetail
	penduduk.NIK = input.NIK
	penduduk.NoKK = input.NoKK
	penduduk.Fullname = input.Fullname
	penduduk.BirthDate = birthDate
	penduduk.BirthPlace = input.BirthPlace
	penduduk.JK = input.JK
	penduduk.ReligionID = input.ReligionID
	penduduk.Nationality = input.Nationality
	penduduk.PekerjaanID = input.PekerjaanID
	penduduk.PendidikanID = input.PendidikanID
	penduduk.Address = input.Address
	penduduk.RtID = input.RtID
	penduduk.RwID = input.RwID
	penduduk.MariedType = input.MariedType
	penduduk.MariedDate = mariedDate
	penduduk.BloodType = input.BloodType
	penduduk.KelurahanID = input.KelurahanID
	penduduk.KecamatanID = input.KecamatanID
	penduduk.CreatedAt = input.CreatedAt
	penduduk.StatusFamily = input.StatusFamily
	penduduk.UpdatedAt = time.Now()
	penduduk.Active = true

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

func (s *service) GetPenduduks(pagination helper.Pagination, NIK string) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, NIK)

	return pagination, err
}

func (s *service) GetCountPenduduk() (int64, error) {
	count, err := s.repository.CountAll()
	return count, err
}
