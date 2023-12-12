package domisili

type GetDomisiliDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateDomisiliInput struct {
	NIK              string `form:"nik"`
	KodeSurat        string `form:"kode_surat"`
	Type             string `form:"type" binding:"required"`
	NamaPerusahaan   string `form:"nama_perusahaan"`
	JenisPerusahaan  string `form:"jenis_perusahaan"`
	TelpPerusahaan   string `form:"telp_perusahaan"`
	StatusBangunan   string `form:"status_bangunan"`
	AktaPerusahaan   string `form:"akta_perusahaan"`
	SKPengesahan     string `form:"sk_pengesahan"`
	AlamatPerusahaan string `form:"alamat_perusahaan"`
	PenanggungJawab  string `form:"penanggung_jawab"`
	LampiranPath     string `form:"lampiran_path"`
	Keterangan       string `form:"keterangan" binding:"required"`
}
