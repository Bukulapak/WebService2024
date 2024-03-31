// type.go
package _714220031

import "go.mongodb.org/mongo-driver/bson/primitive"

// "time"

type Responden struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Usia         string               `bson:"usia,omitempty" json:"usia,omitempty"`
	JenisKelamin string             `bson:"jenis_kelamin,omitempty" json:"jenis_kelamin,omitempty"`
	Pekerjaan    string             `bson:"pekerjaan,omitempty" json:"pekerjaan,omitempty"`
	// TanggalPengisian primitive.DateTime `bson:"tanggal_pengisian,omitempty" json:"tanggal_pengisian,omitempty"`
}

type Pertanyaan struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaPertanyaan string             `bson:"nama_pertanyaan,omitempty" json:"pertanyaan,omitempty"`
}

type Jawaban struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	IDResponden  primitive.ObjectID `bson:"id_responden,omitempty" json:"id_responden,omitempty"`
	IDPertanyaan primitive.ObjectID `bson:"id_pertanyaan,omitempty" json:"id_pertanyaan,omitempty"`
	NamaJawaban  string             `bson:"jawaban,omitempty" json:"jawaban,omitempty"`
	TanggalJawab string `bson:"tanggal_jawab,omitempty" json:"tanggal_jawab,omitempty"`
}
