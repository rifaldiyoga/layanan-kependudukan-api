package sku

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"
)

type Service interface {
	GetSKUByID(ID int) (SKU, error)
	GetSKUs(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	GetLastSKU() (SKU, error)
	CreateSKU(input CreateSKUInput, layanan layanan.Layanan, user user.User) (SKU, error)
	UpdateSKU(ID GetSKUDetailInput, input CreateSKUInput) (SKU, error)
	DeleteSKU(ID GetSKUDetailInput) error
}

type service struct {
	repository Repository
}

func NewService(repository *repository) *service {
	return &service{repository}
}

func (s *service) GetSKUByID(ID int) (SKU, error) {
	sku, err := s.repository.FindByID(ID)

	return sku, err
}

func (s *service) GetLastSKU() (SKU, error) {
	sku, err := s.repository.FindLast()

	return sku, err
}

func (s *service) CreateSKU(input CreateSKUInput, layanan layanan.Layanan, user user.User) (SKU, error) {
	sku := SKU{}

	lastSKU, _ := s.repository.FindLast()

	sku.Keterangan = input.Keterangan
	sku.Usaha = input.Usaha
	sku.KodeSurat = helper.GenerateKodeSurat(layanan.Code, lastSKU.KodeSurat)
	sku.NIK = input.NIK
	sku.CreatedBy = user.ID
	sku.Status = false
	sku.CreatedAt = time.Now()

	newSKU, err := s.repository.Save(sku)

	return newSKU, err
}

func (s *service) UpdateSKU(inputDetail GetSKUDetailInput, input CreateSKUInput) (SKU, error) {
	sku := SKU{}

	sku.Keterangan = input.Keterangan
	sku.KodeSurat = input.KodeSurat
	sku.Usaha = input.Usaha
	sku.NIK = input.NIK
	sku.CreatedAt = time.Now()

	newSKU, err := s.repository.Update(sku)
	return newSKU, err
}

func (s *service) DeleteSKU(inputDetail GetSKUDetailInput) error {
	sku, errId := s.repository.FindByID(inputDetail.ID)
	if errId != nil {
		return errId
	}
	sku.ID = inputDetail.ID

	err := s.repository.Delete(sku)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetSKUs(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	pagination, err := s.repository.FindAll(pagination, params)

	return pagination, err
}
