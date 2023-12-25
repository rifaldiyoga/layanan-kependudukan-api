package sistem

type Sistem struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Telp      string `json:"telp"`
	KodePos   string `json:"kode_pos"`
	Kota      string `json:"kota"`
	Kecamatan string `json:"kecamatan"`
}

func (Sistem) TableName() string {
	return "tb_sistem"
}
