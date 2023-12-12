package kelahiran

type GetKelahiranDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateKelahiranInput struct {
	NIK               string `form:"nik"`
	Nama              string `form:"nama"`
	BirthDate         string `form:"birth_date"`
	BirthPlace        string `form:"birth_place"`
	AnakKe            int    `form:"anak_ke"`
	Jam               string `form:"jam"`
	JK                string `form:"jk"`
	NikAyah           string `form:"nik_ayah"`
	NikIbu            string `form:"nik_ibu"`
	LampiranBukuNikah string `form:"lampiran_buku_nikah_path"`
	LampiranKetRs     string `form:"lampiran_ket_rs_path"`
	KecamatanID       int    `form:"subdistrict_id"`
	KotaID            int    `form:"district_id"`
	ProvinsiID        int    `form:"province_id"`
	Keterangan        string `form:"keterangan"`
}
