package _714220013

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MataKuliah struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_Matkul string             `bson:"kode_matkul,omitempty" json:"kode_matkul,omitempty"`
	Nama_Matkul string             `bson:"nama_matkul,omitempty" json:"nama_matkul,omitempty"`
	Sks         int                `bson:"sks,omitempty" json:"sks,omitempty"`
	Semester    int                `bson:"semester,omitempty" json:"semester,omitempty"`
}

type Dosen struct {
	Nidn        string   				`bson:"nidn,omitempty" json:"nidn,omitempty"`
	Nama_Dosen  string   				`bson:"nama_dosen,omitempty" json:"nama_dosen,omitempty"`
}

type Jadwal struct {
	ID            primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
	Hari          string            	`bson:"hari,omitempty" json:"hari,omitempty"`
	Waktu_Mulai   string             	`bson:"waktu_mulai,omitempty" json:"waktu_mulai,omitempty"`
	Waktu_Selesai string             	`bson:"waktu_selesai,omitempty" json:"waktu_selesai,omitempty"`
	Mata_Kuliah   MataKuliah            `bson:"mata_kuliah,omitempty" json:"mata_kuliah,omitempty"`
	Ruangan       int				 	`bson:"ruangan,omitempty" json:"ruangan,omitempty"`
	Dosen         Dosen             	`bson:"dosen,omitempty" json:"dosen,omitempty"`
}
