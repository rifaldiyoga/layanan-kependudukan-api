package layanan

import (
	"layanan-kependudukan-api/helper"
)

const PENDUDUKAN = "PENDUDUKAN"
const RT = "RT"
const RW = "RW"
const ADMIN = "ADMIN"
const KELURAHAN = "KELURAHAN"

type TypeFormatter struct {
	Type string             `json:"type"`
	Data []LayananFormatter `json:"data"`
}

type LayananFormatter struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	IsConfirm bool   `json:"is_confirm"`
	IsSign    bool   `json:"is_sign"`
	Info      string `json:"info"`
	CreatedAt string `json:"created_at"`
}

func FormatLayanan(Layanan Layanan) LayananFormatter {
	formatter := LayananFormatter{
		ID:        Layanan.ID,
		Code:      Layanan.Code,
		Name:      Layanan.Name,
		Type:      Layanan.Type,
		IsConfirm: Layanan.IsConfirm,
		IsSign:    Layanan.IsSign,
		Info:      Layanan.Info,
		CreatedAt: helper.FormatDateToString(Layanan.CreatedAt),
	}

	return formatter
}

func FormatType(Layanan string, Layanans []LayananFormatter) TypeFormatter {
	formatter := TypeFormatter{
		Type: helper.GetType(Layanan),
		Data: Layanans,
	}

	return formatter
}

func FormatTypes(Types []string, Layanans []LayananFormatter) []TypeFormatter {

	var LayanansFormatter []TypeFormatter

	for _, Type := range Types {
		LayananFormatter := FormatType(Type, filterList(Layanans, Type))
		LayanansFormatter = append(LayanansFormatter, LayananFormatter)
	}

	return LayanansFormatter
}

func FormatLayanans(Layanans []Layanan) []LayananFormatter {

	var LayanansFormatter []LayananFormatter

	for _, Layanan := range Layanans {
		LayananFormatter := FormatLayanan(Layanan)
		LayanansFormatter = append(LayanansFormatter, LayananFormatter)
	}

	return LayanansFormatter
}

func filterList(list []LayananFormatter, Type string) []LayananFormatter {
	var result []LayananFormatter

	for _, item := range list {
		if item.Type == Type {
			result = append(result, item)
		}
	}

	return result
}
