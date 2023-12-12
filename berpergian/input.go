package berpergian

type GetBerpergianDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateBerpergianInput struct {
	NIK          string `form:"nik"`
	KodeSurat    string `form:"kode_surat"`
	Lampiran     string `form:"lampiran_path"`
	Keterangan   string `form:"keterangan"`
	Tujuan       string `form:"tujuan"`
	TglBerangkat string `form:"tgl_berangkat"`
	TglKembali   string `form:"tgl_kembali"`
}
