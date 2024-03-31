package _1174025

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Karyawan struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jabatan      string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Jam_kerja    []JamKerja         `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
	Hari_kerja   []string           `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty"`
}

type JamKerja struct {
	Durasi     int      `bson:"durasi,omitempty" json:"durasi,omitempty"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty"`
	Shift      int      `bson:"shift,omitempty" json:"shift,omitempty"`
	Piket_tim  string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Longitude    float64            `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude     float64            `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin      string             `bson:"checkin,omitempty" json:"checkin,omitempty"`
	Biodata      Karyawan           `bson:"biodata,omitempty" json:"biodata,omitempty"`
}

type Penggajian struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KaryawanID    primitive.ObjectID `bson:"karyawan_id,omitempty" json:"karyawan_id,omitempty"`
	Bulan         string             `bson:"bulan,omitempty" json:"bulan,omitempty"`
	Tahun         int                `bson:"tahun,omitempty" json:"tahun,omitempty"`
	Gaji_pokok    float64            `bson:"gaji_pokok,omitempty" json:"gaji_pokok,omitempty"`
	Tunjangan     float64            `bson:"tunjangan,omitempty" json:"tunjangan,omitempty"`
	Potongan      float64            `bson:"potongan,omitempty" json:"potongan,omitempty"`
	Total_gaji    float64            `bson:"total_gaji,omitempty" json:"total_gaji,omitempty"`
	Tanggal_bayar primitive.DateTime `bson:"tanggal_bayar,omitempty" json:"tanggal_bayar,omitempty"`
}

type Honor struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	KaryawanID      primitive.ObjectID `bson:"karyawan_id,omitempty" json:"karyawan_id,omitempty"`
	Pekerjaan       string             `bson:"pekerjaan,omitempty" json:"pekerjaan,omitempty"`
	Tanggal_mulai   primitive.DateTime `bson:"tanggal_mulai,omitempty" json:"tanggal_mulai,omitempty"`
	Tanggal_selesai primitive.DateTime `bson:"tanggal_selesai,omitempty" json:"tanggal_selesai,omitempty"`
	Honor_perjam    float64            `bson:"honor_perjam,omitempty" json:"honor_perjam,omitempty"`
	Total_honor     float64            `bson:"total_honor,omitempty" json:"total_honor,omitempty"`
	Tanggal_bayar   primitive.DateTime `bson:"tanggal_bayar,omitempty" json:"tanggal_bayar,omitempty"`
}
