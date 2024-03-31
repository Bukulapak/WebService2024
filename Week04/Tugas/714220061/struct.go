package _714220061

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DataPembeli struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	NomorTelepon string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Alamat       string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
}

type DataTagihan struct {
	ID    		 primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama  		 DataPembeli        `bson:"nama,omitempty" json:"nama,omitempty"`
	Tanggal		 primitive.DateTime `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
	Total  		 float64            `bson:"durasi,omitempty" json:"durasi,omitempty"`
}

type DataBarang struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaBarang	 string				`bson:"nama_barang,omitempty" json:"nama_barang,omitempty"`
	HargaBarang	 float64			`bson:"harga_barang,omitempty" json:"harga_barang,omitempty"`
}
