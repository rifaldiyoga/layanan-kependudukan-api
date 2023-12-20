package rw

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/kelurahan"
)

type RWFormatter struct {
	ID          int                          `json:"id"`
	Code        string                       `json:"code"`
	Name        string                       `json:"name"`
	CreatedAt   string                       `json:"created_at"`
	KelurahanID int                          `json:"kelurahan_id"`
	Kelurahan   kelurahan.KelurahanFormatter `json:"kelurahan"`
}

func FormatRW(rw RW) RWFormatter {
	formatter := RWFormatter{
		ID:          rw.ID,
		Code:        rw.Code,
		Name:        rw.Name,
		KelurahanID: rw.KelurahanID,
		CreatedAt:   helper.FormatDateToString(rw.CreatedAt),
		Kelurahan:   kelurahan.FormatKelurahan(rw.Kelurahan),
	}

	return formatter
}

func FormatRWs(rws []RW) []RWFormatter {
	var rwsFormatter []RWFormatter

	for _, rw := range rws {
		rwFormatter := FormatRW(rw)
		rwsFormatter = append(rwsFormatter, rwFormatter)
	}

	return rwsFormatter
}
