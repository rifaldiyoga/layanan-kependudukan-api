package sistem

type GetSistemDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateSistemInput struct {
	Code        string `json:"code"`
	Nama        string `json:"nama"`
	Alamat      string `json:"alamat"`
	Telp        string `json:"telp"`
	KodePos     string `json:"kode_pos"`
	KecamatanID int    `json:"subdistrict_id"`
	KotaID      int    `json:"district_id"`
	ProvinsiID  int    `json:"province_id"`
}
