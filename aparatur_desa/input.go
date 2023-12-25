package aparatur_desa

type GetAparaturDesaDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreateAparaturDesaInput struct {
	NIP       string `json:"nip" `
	JabatanID string `json:"jabatan_id" `
	Nama      string `json:"nama" `
}
