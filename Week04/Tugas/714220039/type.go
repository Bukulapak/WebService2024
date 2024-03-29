package _14220039

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama          string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number  string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Kelas         string             `bson:"kelas,omitempty" json:"kelas,omitempty"`
	Alamat        Alamat             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Jenis_kelamin string             `bson:"jeniskelamin,omitempty" json:"jeniskelamin,omitempty"`
}

type Alamat struct {
	Jalan    string `bson:"jalan,omitempty" json:"jalan,omitempty"`
	No_rumah string `bson:"no_rumah,omitempty" json:"no_rumah,omitempty"`
	RT       string `bson:"rt,omitempty" json:"rt,omitempty"`
	RW       string `bson:"rw,omitempty" json:"rw,omitempty"`
	KodePos  int    `bson:"kodepos,omitempty" json:"kodepos,omitempty"`
}

type Nilai struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nilai_MHS    string             `bson:"nilai_mhs,omitempty" json:"nilai_mhs,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Mata_Kuliah  string             `bson:"mata_kuliah,omitempty" json:"mata_kuliah,omitempty"`
	Biodata      Mahasiswa          `bson:"biodata,omitempty" json:"biodata,omitempty"`
}
