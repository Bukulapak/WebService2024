package _714220035

import "go.mongodb.org/mongo-driver/bson/primitive"

type Dosen struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama   string             `bson:"nama,omitempty" json:"nama,omitempty"`
	NIP    string             `bson:"nip,omitempty" json:"nip,omitempty"`
	Matkul string             `bson:"matkul,omitempty" json:"matkul,omitempty"`
}

type Mahasiswa struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	NIM      string             `bson:"nim,omitempty" json:"nim,omitempty"`
	Semester int                `bson:"semester,omitempty" json:"semester,omitempty"`
}

type Perwalian struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Mahasiswa Mahasiswa          `bson:"mahasiswa,omitempty" json:"mahasiswa,omitempty"`
	Dosen     Dosen              `bson:"dosen,omitempty" json:"dosen,omitempty"`
	Matkul    string             `bson:"matkul,omitempty" json:"matkul,omitempty"`
	Status    string             `bson:"status,omitempty" json:"status,omitempty"`
}
