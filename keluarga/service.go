package keluarga

import (
	"layanan-kependudukan-api/helper"
	"strconv"
	"time"
)

type Service interface {
	GetKeluargaByID(ID int) (Keluarga, error)
	GetKeluargas(pagination helper.Pagination) (helper.Pagination, error)
	GetKeluargaUser(NoKK string) (Keluarga, error)
	CreateKeluarga(input CreateKeluargaInput) (Keluarga, error)
	UpdateKeluarga(ID GetKeluargaDetailInput, input CreateKeluargaInput) (Keluarga, error)
	DeleteKeluarga(ID GetKeluargaDetailInput) error
	GetCountKeluarga() (int64, error)
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetKeluargaByID(ID int) (Keluarga, error) {
	keluarga, err := s.repository.FindByID(ID)

	return keluarga, err
}

func (s *service) CreateKeluarga(input CreateKeluargaInput) (Keluarga, error) {
	keluarga := Keluarga{}

	if input.NoKK == "" {
		var counter int
		last, err := s.repository.FindLast()
		if err != nil {
			counter = 1
		} else {
			i, _ := strconv.Atoi(last.NoKK[len(last.NoKK)-4:])
			counter = i + 1
		}

		keluarga.NoKK = helper.GenerateNoKK(counter)
	} else {
		keluarga.NoKK = input.NoKK
	}
	keluarga.NIKKepalaKeluarga = input.NIKKepalaKeluarga
	keluarga.KepalaKeluarga = input.KepalaKeluarga
	keluarga.Address = input.Address
	keluarga.RtID = input.RtID
	keluarga.RwID = input.RwID
	keluarga.KelurahanID = input.KelurahanID
	keluarga.KecamatanID = input.KecamatanID
	keluarga.CreatedAt = time.Now()
	keluarga.UpdatedAt = time.Now()

	newKeluarga, err := s.repository.Save(keluarga)

	return newKeluarga, err
}

func (s *service) UpdateKeluarga(inputDetail GetKeluargaDetailInput, input CreateKeluargaInput) (Keluarga, error) {
	lastKeluarga, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return lastKeluarga, err
	}
	keluarga := Keluarga{}
	keluarga.ID = inputDetail.ID
	keluarga.NoKK = lastKeluarga.NoKK
	keluarga.NIKKepalaKeluarga = input.NIKKepalaKeluarga
	keluarga.KepalaKeluarga = input.KepalaKeluarga
	keluarga.Address = input.Address
	keluarga.RtID = input.RtID
	keluarga.RwID = input.RwID
	keluarga.KelurahanID = input.KelurahanID
	keluarga.KecamatanID = input.KecamatanID
	keluarga.CreatedAt = lastKeluarga.CreatedAt
	keluarga.UpdatedAt = time.Now()

	newKeluarga, err := s.repository.Update(keluarga)
	return newKeluarga, err
}

func (s *service) DeleteKeluarga(inputDetail GetKeluargaDetailInput) error {
	keluarga, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}
	keluarga.ID = inputDetail.ID

	err := s.repository.Delete(keluarga)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetKeluargas(pagination helper.Pagination) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination)

	return pagination, err
}

func (s *service) GetKeluargaUser(NoKK string) (Keluarga, error) {
	keluarga, err := s.repository.FindByNoKK(NoKK)

	return keluarga, err
}

func (s *service) GetCountKeluarga() (int64, error) {
	count, err := s.repository.CountAll()
	return count, err
}
