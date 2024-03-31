package _714220030

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tiket struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NomorTiket  string             `bson:"nomor_tiket,omitempty" json:"nomor_tiket,omitempty"`
	Prioritas   string             `bson:"prioritas,omitempty" json:"prioritas,omitempty"`
	Kategori    string             `bson:"kategori,omitempty" json:"kategori,omitempty"`
	Deskripsi   string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Status      string             `bson:"status,omitempty" json:"status,omitempty"`
	Datetime    primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Penyelia    string             `bson:"penyelia,omitempty" json:"penyelia,omitempty"`
	AssignedTo  string             `bson:"assigned_to,omitempty" json:"assigned_to,omitempty"`
	Komentar    []Komentar         `bson:"komentar,omitempty" json:"komentar,omitempty"`
}

type Komentar struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Pengirim  string             `bson:"pengirim,omitempty" json:"pengirim,omitempty"`
	Isi       string             `bson:"isi,omitempty" json:"isi,omitempty"`
	Datetime  primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
}

type Helpdesk struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Departemen  string             `bson:"departemen,omitempty" json:"departemen,omitempty"`
	Ext         string             `bson:"ext,omitempty" json:"ext,omitempty"`
	Email       string             `bson:"email,omitempty" json:"email,omitempty"`
	Tiket       []Tiket            `bson:"tiket,omitempty" json:"tiket,omitempty"`
}
