package keramaian

type GetKeramaianDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateKeramaianInput struct {
	NIK        string `form:"nik"`
	KodeSurat  string `form:"kode_surat"`
	NamaAcara  string `form:"nama_acara" binding:"required"`
	Tanggal    string `form:"tanggal" binding:"required"`
	Waktu      string `form:"waktu" binding:"required"`
	Tempat     string `form:"tempat" binding:"required"`
	Alamat     string `form:"alamat" binding:"required"`
	Telpon     string `form:"telpon" binding:"required"`
	Lampiran   string `form:"lampiran_path"`
	Keterangan string `form:"keterangan" binding:"required"`
}
