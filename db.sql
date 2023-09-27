CREATE TABLE "tb_jabatan" (
  "id" serial NOT NULL,
  "jabatan" varchar(255) NOT NULL,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_kelahiran" (
  "id" serial NOT NULL,
  "nama" varchar(255) NOT NULL,
  "gender_id" int NOT NULL,
  "birth_date" varchar(255) NOT NULL,
  "birth_place" varchar(255) NOT NULL,
  "nik_ayah" varchar(255) NOT NULL,
  "nik_ibu" varchar(255) NOT NULL,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_kematian" (
  "id" serial NOT NULL,
  "nik" varchar(255) NOT NULL,
  "tgl_kematian" timestamp NOT NULL,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_layanan" (
  "id" serial NOT NULL,
  "name" varchar(255) DEFAULT NULL ,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_provinsi" (
  "id" serial NOT NULL,
  "code" varchar(255) DEFAULT NULL ,
  "name" varchar(255) DEFAULT NULL ,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_kota" (
  "id" serial NOT NULL,
  "code" varchar(255) DEFAULT NULL ,
  "name" varchar(255) DEFAULT NULL ,
  "provinsi_id" int DEFAULT NULL,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_pengajuan" (
  "id" serial NOT NULL,
  "nik" varchar(255) DEFAULT NULL ,
  "layanan_id" int DEFAULT NULL,
  "status" varchar(255) DEFAULT NULL ,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_aparatur_desa" (
  "id" serial NOT NULL,
  "nip" varchar(255) NOT NULL,
  "nama" varchar(255) NOT NULL,
  "jabatan_id" int NOT NULL,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_kecamatan" (
  "id" serial NOT NULL,
  "code" varchar(255) DEFAULT NULL ,
  "name" varchar(255) DEFAULT NULL ,
  "kota_id" int DEFAULT NULL,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_penduduk" (
  "id" serial NOT NULL,
  "nik" varchar(255) NOT NULL ,
  "fullname" varchar(255) DEFAULT NULL ,
  "gender" int DEFAULT NULL,
  "birth_place" varchar(255) DEFAULT NULL ,
  "birth_date" date DEFAULT NULL,
  "religion_id" int DEFAULT NULL,
  "pendidikan_id" int DEFAULT NULL,
  "pekerjaan_id" int DEFAULT NULL,
  "blood_type" varchar(255) DEFAULT NULL ,
  "nationality" varchar(255) DEFAULT NULL ,
  "maried_type" int DEFAULT NULL,
  "maried_date" date DEFAULT NULL,
  "kk_id" int DEFAULT NULL,
  "rt_id" int DEFAULT NULL,
  "rw_id" int DEFAULT NULL,
  "kelurahan_id" int DEFAULT NULL,
  "kecamatan_id" int DEFAULT NULL,
  "kota_id" int DEFAULT NULL,
  "provinsi_id" int DEFAULT NULL,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id", "nik")
);

CREATE TABLE "tb_agama" (
  "id" serial NOT NULL,
  "code" varchar(255) DEFAULT NULL ,
  "name" varchar(255) DEFAULT NULL ,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_kartu_keluarga" (
  "id" serial NOT NULL,
  "no_kk" varchar(255) DEFAULT NULL ,
  "kepala_keluarga_id" int DEFAULT NULL,
  "kepala_kelarga" varchar(255) DEFAULT NULL ,
  "rt_id" int DEFAULT NULL,
  "rw_id" int DEFAULT NULL,
  "kelurahan_id" int DEFAULT NULL,
  "kecamatan_id" int DEFAULT NULL,
  "kota_id" int DEFAULT NULL,
  "provinsi_id" int DEFAULT NULL,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_kelurahan" (
  "id" serial NOT NULL,
  "code" varchar(255) DEFAULT NULL ,
  "name" varchar(255) DEFAULT NULL ,
  "kecamatan_id" int DEFAULT NULL,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_nikah" (
  "id" serial NOT NULL,
  "nik_istri" varchar(255) NOT NULL,
  "nik_suami" varchar(255) NOT NULL,
  "status" varchar(255) NOT NULL,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_pekerjaan" (
  "id" serial NOT NULL,
  "code" varchar(255) DEFAULT NULL ,
  "name" varchar(255) DEFAULT NULL ,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_pendidikan" (
  "id" serial NOT NULL,
  "code" varchar(255) DEFAULT NULL ,
  "name" varchar(255) DEFAULT NULL ,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_tanah" (
  "id" serial NOT NULL,
  "nik" varchar(255) DEFAULT NULL ,
  "nama_pemilik" varchar(255) DEFAULT NULL ,
  "lokasi" varchar(255) DEFAULT NULL ,
  "luas_tanah" int DEFAULT NULL,
  "panjang" numeric DEFAULT NULL,
  "lebar" numeric DEFAULT NULL,
  "batas_barat" varchar(255) DEFAULT NULL ,
  "batas_timuir" varchar(255) DEFAULT NULL ,
  "batas_utara" varchar(255) DEFAULT NULL ,
  "batas_selatan" varchar(255) DEFAULT NULL ,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_usaha" (
  "id" serial NOT NULL,
  "nik" varchar(255) DEFAULT NULL ,
  "nama_pemilik" varchar(255) DEFAULT NULL ,
  "nama_usaha" varchar(255) DEFAULT NULL ,
  "jenis_usaha" varchar(255) DEFAULT NULL ,
  "pekerjaan" varchar(255) DEFAULT NULL ,
  "tgl_berlaku" date DEFAULT NULL,
  "npwp" varchar(255) DEFAULT NULL ,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_user" (
  "id" serial NOT NULL,
  "nik" varchar(255) DEFAULT NULL ,
  "name" varchar(255) DEFAULT NULL ,
  "avatar_file_name" varchar(255) DEFAULT NULL ,
  "role" varchar(255) DEFAULT NULL ,
  "email" varchar(255) DEFAULT NULL ,
  "password" varchar(255) DEFAULT NULL ,
  "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_rw" (
  "id" serial NOT NULL,
  "code" varchar(255) DEFAULT NULL ,
  "name" varchar(255) DEFAULT NULL ,
  "kelurahan_id" int DEFAULT NULL,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_tempat_usaha" (
  "id" serial NOT NULL,
  "usaha_id" int DEFAULT NULL,
  "nik" varchar(255) DEFAULT NULL ,
  "nama_pemilik" varchar(255) DEFAULT NULL ,
  "nama_usaha" varchar(255) DEFAULT NULL ,
  "jenis_usaha" varchar(255) DEFAULT NULL ,
  "tgl_berlaku" date DEFAULT NULL,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);

CREATE TABLE "tb_rt" (
  "id" serial NOT NULL,
  "code" varchar(255) DEFAULT NULL ,
  "name" varchar(255) DEFAULT NULL ,
  "rw_id" int DEFAULT NULL,
    "created_at" timestamp,
  "updated_at" timestamp,
  PRIMARY KEY ("id")
);