package aparatur_desa

import (
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/position"
)

type AparaturDesaFormatter struct {
	ID        int                        `json:"id"`
	NIP       string                     `json:"nip"`
	Nama      string                     `json:"nama"`
	JabatanID int                        `json:"jabatan_id"`
	CreatedAt string                     `json:"created_at"`
	Jabatan   position.PositionFormatter `json:"jabatan" `
}

func FormatAparaturDesa(kelurahan AparaturDesa) AparaturDesaFormatter {
	formatter := AparaturDesaFormatter{
		ID:        kelurahan.ID,
		NIP:       kelurahan.Nip,
		JabatanID: kelurahan.JabatanID,
		Nama:      kelurahan.Nama,
		CreatedAt: helper.FormatDateToString(kelurahan.CreatedAt),
		Jabatan:   position.FormatPosition(kelurahan.Jabatan),
	}

	return formatter
}

func FormatAparaturDesas(kelurahans []AparaturDesa) []AparaturDesaFormatter {
	var kelurahansFormatter []AparaturDesaFormatter

	for _, kelurahan := range kelurahans {
		kelurahanFormatter := FormatAparaturDesa(kelurahan)
		kelurahansFormatter = append(kelurahansFormatter, kelurahanFormatter)
	}

	return kelurahansFormatter
}
