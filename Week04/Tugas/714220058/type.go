package _714220058

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	NomorTelepon string             `bson:"nomor_telepon,omitempty" json:"nomor_telepon,omitempty"`
	ProgramStudi string             `bson:"program_studi,omitempty" json:"program_studi,omitempty"`
	JamKuliah    []JamKuliah        `bson:"jam_kuliah,omitempty" json:"jam_kuliah,omitempty"`
	HariKuliah   []string           `bson:"hari_kuliah,omitempty" json:"hari_kuliah,omitempty"`
}

type JamKuliah struct {
	Durasi    int      `bson:"durasi,omitempty" json:"durasi,omitempty"`
	JamMasuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	JamKeluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	GMT       int      `bson:"gmt,omitempty" json:"gmt,omitempty"`
	Hari      []string `bson:"hari,omitempty" json:"hari,omitempty"`
	Shift     int      `bson:"shift,omitempty" json:"shift,omitempty"`
	Kelas     string   `bson:"kelas,omitempty" json:"kelas,omitempty"`
}

type Pendaftaran struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Longitude        float64            `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude         float64            `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Alamat           string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	NomorTelepon     string             `bson:"nomor_telepon,omitempty" json:"nomor_telepon,omitempty"`
	TanggalDaftar    primitive.DateTime `bson:"tanggal_daftar,omitempty" json:"tanggal_daftar,omitempty"`
	Status           string             `bson:"status,omitempty" json:"status,omitempty"`
	BiodataMahasiswa Mahasiswa          `bson:"biodata_mahasiswa,omitempty" json:"biodata_mahasiswa,omitempty"`
}
