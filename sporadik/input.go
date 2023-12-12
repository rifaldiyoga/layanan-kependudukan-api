package sporadik

type GetSporadikDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateSporadikInput struct {
	NIK                  string `json:"nik"`
	KodeSurat            string `json:"kode_surat"`
	Keterangan           string `json:"keterangan"`
	LampiranPemohon      string `json:"lampiran_pemohon"`
	LampiranSporadikLama string `json:"lampiran_sporadik_lama"`
	LampiranSporadikBaru string `json:"lampiran_sporadik_baru"`
	LampiranBukti        string `json:"lampiran_bukti"`
	LampiranLunasPbb     string `json:"lampiran_lunas_pbb"`
}
