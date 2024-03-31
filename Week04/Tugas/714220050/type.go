package _714220050

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jurusan      string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
}

type Matkul struct {
	Nama_matkul     string      `bson:"durasi,omitempty" json:"durasi,omitempty"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty"`
	Sks      int      `bson:"shift,omitempty" json:"shift,omitempty"`
	Dosen      string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Biodata      Mahasiswa           `bson:"biodata,omitempty" json:"biodata,omitempty"`
}
