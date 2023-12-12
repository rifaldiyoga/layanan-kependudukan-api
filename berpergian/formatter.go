package berpergian

import (
	"layanan-kependudukan-api/berpergian_detail"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type BerpergianFormatter struct {
	Berpergian
	CreatedAt        string                                        `json:"created_at"`
	TglBerangkat     string                                        `json:"tgl_berangkat"`
	TglKembali       string                                        `json:"tgl_kembali"`
	Penduduk         penduduk.PendudukFormatter                    `json:"penduduk"`
	BerpergianDetail []berpergian_detail.BerpergianDetailFormatter `json:"berpergian_detail"`
}

func FormatBerpergian(berpergian Berpergian) BerpergianFormatter {
	formatter := BerpergianFormatter{
		Berpergian:       berpergian,
		CreatedAt:        helper.FormatDateToString(berpergian.CreatedAt),
		TglBerangkat:     helper.FormatDateToString(berpergian.TglBerangkat),
		TglKembali:       helper.FormatDateToString(berpergian.TglKembali),
		Penduduk:         penduduk.FormatPenduduk(berpergian.Penduduk),
		BerpergianDetail: berpergian_detail.FormatBerpergianDetails(berpergian.BerpergianDetail),
	}

	return formatter
}

func FormatBerpergians(sktms []Berpergian) []BerpergianFormatter {
	var sktmsFormatter []BerpergianFormatter

	for _, sktm := range sktms {
		sktmFormatter := FormatBerpergian(sktm)
		sktmsFormatter = append(sktmsFormatter, sktmFormatter)
	}

	return sktmsFormatter
}
