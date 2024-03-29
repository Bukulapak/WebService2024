package _714220054

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Undangan_Rapat struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Judul   string             `bson:"judul,omitempty" json:"judul,omitempty"`
	Tanggal string            `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
}

type Tanggal struct {
	Tanggal    string `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
}

type Rapat_Makrab struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Longitude  float64            `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude   float64            `bson:"latitude,omitempty" json:"latitude,omitempty"`
	UKM        string             `bson:"location,omitempty" json:"location,omitempty"`
	Alamat     string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Keterangan Undangan_Rapat     `bson:"undangan_rapat,omitempty" json:"undangan_rapat,omitempty"`
}
