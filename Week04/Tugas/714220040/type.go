package _714220040

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Prodi        string             `bson:"prodi,omitempty" json:"prodi,omitempty"`
}

type DataOrtu struct {
	NamaIbu       string `bson:"namaIbu,omitempty" json:"namaIbu,omitempty"`
	PekerjaanIbu  string `bson:"pekerjaanIbu,omitempty" json:"pekerjaanIbu,omitempty"`
	NamaAyah      string `bson:"namaAyah,omitempty" json:"namaAyah,omitempty"`
	PekerjaanAyah string `bson:"pekerjaanAyah,omitempty" json:"pekerjaanAyah,omitempty"`
}

type Profile struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Npm          string             `bson:"npm,omitempty" json:"npm,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Prodi        string             `bson:"prodi,omitempty" json:"prodi,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	NamaAyah     string             `bson:"namaAyah,omitempty" json:"namaAyah,omitempty"`
	Biodata      Mahasiswa          `bson:"biodata,omitempty" json:"biodata,omitempty"`
}
