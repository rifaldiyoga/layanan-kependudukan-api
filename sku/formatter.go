package sku

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/penduduk"
)

type SKUFormatter struct {
	SKU
	CreatedAt string                     `json:"created_at"`
	Penduduk  penduduk.PendudukFormatter `json:"penduduk"`
}

func FormatSKU(sku SKU) SKUFormatter {
	formatter := SKUFormatter{
		SKU:       sku,
		CreatedAt: helper.FormatDateToString(sku.CreatedAt),
		Penduduk:  penduduk.FormatPenduduk(sku.Penduduk),
	}

	return formatter
}

func FormatSKUs(skus []SKU) []SKUFormatter {
	var skusFormatter []SKUFormatter

	for _, sku := range skus {
		skuFormatter := FormatSKU(sku)
		skusFormatter = append(skusFormatter, skuFormatter)
	}

	return skusFormatter
}
