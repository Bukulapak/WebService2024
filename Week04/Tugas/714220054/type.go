package _714220054

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Undangan_Rapat struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kepada   string             `bson:"judul,omitempty" json:"judul,omitempty"`
	Divisi   string            `bson:"divisi,omitempty" json:"divisi,omitempty"`
}

type Tanggal struct {
	Waktu time.Time          `bson:"waktu,omitempty" json:"waktu,omitempty"`
}

type Rapat_Makrab struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Longitude  float64            `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude   float64            `bson:"latitude,omitempty" json:"latitude,omitempty"`
	UKM        string             `bson:"location,omitempty" json:"location,omitempty"`
	Alamat     string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Kepada    Undangan_Rapat      `bson:"undangan_rapat,omitempty" json:"undangan_rapat,omitempty"`
	Waktu		Tanggal			  `bson:"waktu,omitempty" json:"waktu,omitempty"`
}
