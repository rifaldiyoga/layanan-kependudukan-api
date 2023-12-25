package aparatur_desa

type GetAparaturDesaDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateAparaturDesaInput struct {
	NIP       string `json:"nip" binding:"required"`
	JabatanID int    `json:"jabatan_id" binding:"required"`
	Nama      string `json:"nama" binding:"required"`
}
