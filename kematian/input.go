package kematian

type GetKematianDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateKematianInput struct {
	NIK           string `form:"nik"`
	NikJenazah    string `form:"nik_jenazah"`
	KodeSurat     string `form:"kode_surat"`
	Keterangan    string `form:"keterangan"`
	TglKematian   string `form:"tgl_kematian"`
	Jam           string `form:"jam"`
	Sebab         string `form:"sebab"`
	Tempat        string `form:"tempat"`
	Saksi1        string `form:"saksi1"`
	Saksi2        string `form:"saksi2"`
	LampiranKetRs string `form:"lampiran_ket_rs_path"`
}
