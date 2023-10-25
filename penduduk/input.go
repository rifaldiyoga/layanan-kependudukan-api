package penduduk

type GetPendudukDetailInput struct {
	ID int `json:"id" binding:"required"`
}

type CreatePendudukInput struct {
	NIK          string `json:"nik"`
	NoKK         string `json:"no_kk"`
	FullName     string `json:"fullname" binding:"required"`
	BirthPlace   string `json:"birth_place" binding:"required"`
	BirthDate    string `json:"birth_date" binding:"required"`
	ReligionID   int    `json:"religion_id" binding:"required"`
	EducationID  int    `json:"education_id" binding:"required"`
	JobID        int    `json:"job_id" binding:"required"`
	Nationality  string `json:"nationality" binding:"required"`
	MariedType   string `json:"maried_type" binding:"required"`
	MariedDate   string `json:"maried_date"`
	BloodType    string `json:"blood_type"`
	Address      string `json:"address" binding:"required"`
	RtID         int    `json:"rt_id" binding:"required"`
	RwID         int    `json:"rw_id" binding:"required"`
	KelurahanID  int    `json:"kelurahan_id" binding:"required"`
	KecamatanID  int    `json:"subdistrict_id" binding:"required"`
	JK           string `json:"jk" binding:"required"`
	StatusFamily string `json:"status_family" `
}
