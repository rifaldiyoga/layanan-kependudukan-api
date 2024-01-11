package pengajuan

import (
	"layanan-kependudukan-api/belum_menikah"
	"layanan-kependudukan-api/berpergian"
	"layanan-kependudukan-api/domisili"
	"layanan-kependudukan-api/helper"
	"layanan-kependudukan-api/janda"
	"layanan-kependudukan-api/kelahiran"
	"layanan-kependudukan-api/kematian"
	"layanan-kependudukan-api/kepolisian"
	"layanan-kependudukan-api/keramaian"
	"layanan-kependudukan-api/layanan"
	"layanan-kependudukan-api/penduduk"
	"layanan-kependudukan-api/penghasilan"
	"layanan-kependudukan-api/pernah_menikah"
	"layanan-kependudukan-api/pindah"
	"layanan-kependudukan-api/rumah"
	"layanan-kependudukan-api/sistem"
	"layanan-kependudukan-api/sktm"
	"layanan-kependudukan-api/sku"
	"layanan-kependudukan-api/sporadik"
	"layanan-kependudukan-api/tanah"
	"layanan-kependudukan-api/user"
	"net/url"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error)
	FindByUser(pagination helper.Pagination, user user.User) (helper.Pagination, error)
	FindByID(id int) (Pengajuan, error)
	CountAll() (int64, error)
	Save(pengajuan Pengajuan) (Pengajuan, error)
	Update(pengajuan Pengajuan) (Pengajuan, error)
	UpdateStatus(pengajuan Pengajuan) (Pengajuan, error)
	Delete(pengajuan Pengajuan) error
}

type repository struct {
	db *gorm.DB
}

func NewRepsitory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(pagination helper.Pagination, params url.Values) (helper.Pagination, error) {
	var pengajuans []Pengajuan

	var stat []string

	db := r.db.Debug()

	if params.Get("status") == "PENDING" {
		stat = append(stat, "APPROVED_RW", "PENDING_ADMIN")
	}

	if params.Get("status") == "REJECTED" {
		stat = append(stat, "REJECTED")
	}

	if params.Get("status") == "VALID" {
		stat = append(stat, "VALID")
	}

	if params.Get("status") == "ALL" {
		stat = append(stat, "VALID", "APPROVED_RW", "PENDING_ADMIN", "REJECTED")
	}

	if len(stat) > 0 {
		db = db.Where("status IN (?)", stat)
	}

	if params.Get("start_date") != "" && params.Get("end_date") != "" {
		db = db.Where("created_at between ? and ?", helper.FormatStringToDate(params.Get("start_date")), helper.FormatStringToDate(params.Get("end_date")))
	}

	err := db.Order("created_at DESC").Scopes(helper.Paginate(pengajuans, &pagination, r.db)).Preload("Detail").Find(&pengajuans).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = pengajuans
	return pagination, err
}

func (r *repository) FindByUser(pagination helper.Pagination, user user.User) (helper.Pagination, error) {
	var pengajuans []Pengajuan

	db := r.db.Debug()
	if user.Role != "PENDUDUK" {
		var penduduk penduduk.Penduduk
		err := r.db.Where("nik = ?", user.Nik).Find(&penduduk).Error
		if err != nil {
			return pagination, err
		}

		var stat []string
		if user.Role == "RT" {
			stat = append(stat, "PENDING_RT", "REJECTED_RT", "APPROVED_RT")
		}
		if user.Role == "RW" {
			stat = append(stat, "APPROVED_RT", "REJECTED_RW")
		}
		if user.Role == "ADMIN" || user.Role == "KELURAHAN" {
			stat = append(stat, "APPROVED_RW", "PENDING_ADMIN", "VALID", "REJECTED")
		}

		db = db.Joins("JOIN tb_penduduk ON tb_penduduk.nik = tb_pengajuan.nik")
		if user.Role == "ADMIN" || user.Role == "KELURAHAN" {
			db = db.Where(" status IN (?)", stat)
		} else if user.Role == "RW" {
			db = db.Where("tb_penduduk.rw_id = ? AND (tb_pengajuan.nik = ? OR status IN (?))", penduduk.RwID, user.Nik, stat)
		} else {
			db = db.Where("tb_penduduk.rt_id = ? AND tb_penduduk.rw_id = ? AND (tb_pengajuan.nik = ? OR status IN (?))", penduduk.RtID, penduduk.RwID, user.Nik, stat)
		}

	} else {
		db = db.Where("tb_pengajuan.nik = ?", user.Nik)
	}

	err := db.Scopes(helper.Paginate(pengajuans, &pagination, db)).Order("tb_pengajuan.created_at DESC").Preload("Detail").Preload("Penduduk").Find(&pengajuans).Error
	if err != nil {
		return pagination, err
	}
	pagination.Data = pengajuans
	return pagination, err
}

func (r *repository) FindByID(ID int) (Pengajuan, error) {
	var pengajuan Pengajuan
	err := r.db.Debug().Where("id = ?", ID).Preload("Detail").Preload("Penduduk").First(&pengajuan).Error
	if err != nil {
		return pengajuan, err
	}

	return pengajuan, nil
}

func (r *repository) Save(pengajuan Pengajuan) (Pengajuan, error) {
	err := r.db.Create(&pengajuan).Error
	if err != nil {
		return pengajuan, err
	}

	return pengajuan, nil
}

func (r *repository) Update(pengajuan Pengajuan) (Pengajuan, error) {
	err := r.db.Updates(&pengajuan).Error
	if err != nil {
		return pengajuan, err
	}

	return pengajuan, nil
}

func (r *repository) Delete(pengajuan Pengajuan) error {
	err := r.db.Delete(&pengajuan).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateStatus(pengajuan Pengajuan) (Pengajuan, error) {
	var currentLayanan layanan.Layanan
	err := r.db.Where("id = ?", pengajuan.LayananID).First(&currentLayanan).Error
	if err != nil {
		return pengajuan, err
	}

	var sistem sistem.Sistem
	err = r.db.First(&sistem).Error
	if err != nil {
		return pengajuan, err
	}

	if currentLayanan.Code == "SKTM" {
		var currentSKTM sktm.SKTM
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentSKTM).Error

		if err != nil {
			return pengajuan, err
		}

		var lastSKTM sktm.SKTM
		_ = r.db.Where("kode_surat != ''").Last(&lastSKTM).Error

		currentSKTM.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastSKTM.KodeSurat)
		currentSKTM.Status = true
		r.db.Save(&currentSKTM)

	}

	if currentLayanan.Code == "SKU" {
		var currentSKU sku.SKU
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentSKU).Error

		if err != nil {
			return pengajuan, err
		}

		var lastSKU sku.SKU
		_ = r.db.Where("kode_surat != ''").Last(&lastSKU).Error

		currentSKU.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastSKU.KodeSurat)
		currentSKU.Status = true
		r.db.Save(&currentSKU)

	}
	if currentLayanan.Code == "SKD" {
		var currentSKD domisili.Domisili
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentSKD).Error

		if err != nil {
			return pengajuan, err
		}

		var lastSKD domisili.Domisili
		_ = r.db.Where("kode_surat != ''").Last(&lastSKD).Error

		currentSKD.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastSKD.KodeSurat)
		currentSKD.Status = true
		r.db.Save(&currentSKD)

	}
	if currentLayanan.Code == "SKKBBK" {
		var currentBerpergian berpergian.Berpergian
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentBerpergian).Error

		if err != nil {
			return pengajuan, err
		}

		var lastBerpergian berpergian.Berpergian
		_ = r.db.Where("kode_surat != ''").Last(&lastBerpergian).Error

		currentBerpergian.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastBerpergian.KodeSurat)
		currentBerpergian.Status = true
		r.db.Save(&currentBerpergian)

	}
	if currentLayanan.Code == "SKKT" || currentLayanan.Code == "SPORADIK" {
		var currentTanah tanah.Tanah
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentTanah).Error

		if err != nil {
			return pengajuan, err
		}

		var lastTanah tanah.Tanah
		_ = r.db.Where("kode_surat != ''").Last(&lastTanah).Error

		currentTanah.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastTanah.KodeSurat)
		currentTanah.Status = true
		r.db.Save(&currentTanah)

	}
	if currentLayanan.Code == "SSP" {
		var currentSporadik sporadik.Sporadik
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentSporadik).Error

		if err != nil {
			return pengajuan, err
		}

		var lastSporadik sporadik.Sporadik
		_ = r.db.Where("kode_surat != ''").Last(&lastSporadik).Error

		currentSporadik.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastSporadik.KodeSurat)
		currentSporadik.Status = true
		r.db.Save(&currentSporadik)

	}
	if currentLayanan.Code == "SIK" {
		var currentKeramaian keramaian.Keramaian
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentKeramaian).Error

		if err != nil {
			return pengajuan, err
		}

		var lastKeramaian keramaian.Keramaian
		_ = r.db.Where("kode_surat != ''").Last(&lastKeramaian).Error

		currentKeramaian.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastKeramaian.KodeSurat)
		currentKeramaian.Status = true
		r.db.Save(&currentKeramaian)

	}
	if currentLayanan.Code == "SKPD" || currentLayanan.Code == "SKPK" {
		var currentPindah pindah.Pindah
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentPindah).Error

		if err != nil {
			return pengajuan, err
		}

		var lastPindah pindah.Pindah
		_ = r.db.Where("kode_surat != ''").Last(&lastPindah).Error

		currentPindah.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastPindah.KodeSurat)
		currentPindah.Status = true
		r.db.Save(&currentPindah)

	}
	if currentLayanan.Code == "SKPN" {
		var currentPernahMenikah pernah_menikah.PernahMenikah
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentPernahMenikah).Error

		if err != nil {
			return pengajuan, err
		}

		var lastPernahMenikah pernah_menikah.PernahMenikah
		_ = r.db.Where("kode_surat != ''").Last(&lastPernahMenikah).Error

		currentPernahMenikah.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastPernahMenikah.KodeSurat)
		currentPernahMenikah.Status = true
		r.db.Save(&currentPernahMenikah)

	}
	if currentLayanan.Code == "SKBPN" {
		var currentBelumMenikah belum_menikah.BelumMenikah
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentBelumMenikah).Error

		if err != nil {
			return pengajuan, err
		}

		var lastBelumMenikah belum_menikah.BelumMenikah
		_ = r.db.Where("kode_surat != ''").Last(&lastBelumMenikah).Error

		currentBelumMenikah.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastBelumMenikah.KodeSurat)
		currentBelumMenikah.Status = true
		r.db.Save(&currentBelumMenikah)

	}
	if currentLayanan.Code == "SKKH" {
		var currentKelahiran kelahiran.Kelahiran
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentKelahiran).Error

		if err != nil {
			return pengajuan, err
		}

		var lastKelahiran kelahiran.Kelahiran
		_ = r.db.Where("kode_surat != ''").Last(&lastKelahiran).Error

		currentKelahiran.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastKelahiran.KodeSurat)
		currentKelahiran.Status = true
		r.db.Save(&currentKelahiran)

	}
	if currentLayanan.Code == "SKKM" {
		var currentKematian kematian.Kematian
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentKematian).Error

		if err != nil {
			return pengajuan, err
		}

		var lastKematian kematian.Kematian
		_ = r.db.Where("kode_surat != ''").Last(&lastKematian).Error

		currentKematian.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastKematian.KodeSurat)
		currentKematian.Status = true
		r.db.Save(&currentKematian)

		//update penduduk
		var currentPenduduk penduduk.Penduduk
		err = r.db.Where("nik = ?", currentKematian.NikJenazah).First(&currentPenduduk).Error

		if err != nil {
			return pengajuan, err
		}
		currentPenduduk.StatusFamily = "MENINGGAL"
		currentPenduduk.Active = false
		r.db.Save(&currentPenduduk)

	}
	if currentLayanan.Code == "SPCK" {
		var currentKepolisian kepolisian.Kepolisian
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentKepolisian).Error

		if err != nil {
			return pengajuan, err
		}

		var lastKepolisian kepolisian.Kepolisian
		_ = r.db.Where("kode_surat != ''").Last(&lastKepolisian).Error

		currentKepolisian.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastKepolisian.KodeSurat)
		currentKepolisian.Status = true
		r.db.Save(&currentKepolisian)

	}
	if currentLayanan.Code == "SKPOT" {
		var currentPenghasilan penghasilan.Penghasilan
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentPenghasilan).Error

		if err != nil {
			return pengajuan, err
		}

		var lastPenghasilan penghasilan.Penghasilan
		_ = r.db.Where("kode_surat != ''").Last(&lastPenghasilan).Error

		currentPenghasilan.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastPenghasilan.KodeSurat)
		currentPenghasilan.Status = true
		r.db.Save(&currentPenghasilan)

	}
	if currentLayanan.Code == "SKJD" {
		var currentJanda janda.Janda
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentJanda).Error

		if err != nil {
			return pengajuan, err
		}

		var lastJanda janda.Janda
		_ = r.db.Where("kode_surat != ''").Last(&lastJanda).Error

		currentJanda.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastJanda.KodeSurat)
		currentJanda.Status = true
		r.db.Save(&currentJanda)

	}
	if currentLayanan.Code == "SKTMR" {
		var currentRumah rumah.Rumah
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentRumah).Error

		if err != nil {
			return pengajuan, err
		}

		var lastRumah rumah.Rumah
		_ = r.db.Where("kode_surat != ''").Last(&lastRumah).Error

		currentRumah.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastRumah.KodeSurat)
		currentRumah.Status = true
		r.db.Save(&currentRumah)

	}
	if currentLayanan.Code == "SKBBK" {
		var currentBerpergian berpergian.Berpergian
		err := r.db.Where("id = ?", pengajuan.RefID).First(&currentBerpergian).Error

		if err != nil {
			return pengajuan, err
		}

		var lastBerpergian berpergian.Berpergian
		_ = r.db.Where("kode_surat != ''").Last(&lastBerpergian).Error

		currentBerpergian.KodeSurat = helper.GenerateKodeSurat(currentLayanan.KodeSurat, sistem.Code, lastBerpergian.KodeSurat)
		currentBerpergian.Status = true
		r.db.Save(&currentBerpergian)

	}

	return pengajuan, err

}

func (r *repository) CountAll() (int64, error) {
	var count int64
	today := time.Now().UTC().Truncate(24 * time.Hour)
	err := r.db.Model(&Pengajuan{}).Where("(status = 'APPROVED_RW' OR status = 'PENDING_ADMIN' OR status = 'VALID' OR status = 'REJECTED') AND DATE(created_at) = ?", today.Format("2006-01-02")).Count(&count).Error
	if err != nil {
		return count, err
	}

	return count, nil
}
