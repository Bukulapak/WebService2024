package _714220046

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NPM           string             `bson:"npm,omitempty" json:"npm,omitempty"`
	Nama          string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Email         string             `bson:"email,omitempty" json:"email,omitempty"`
	Phone_number  int                `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Kelas         string             `bson:"kelas,omitempty" json:"kelas,omitempty"`
	Jurusan       string             `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
	Jenis_kelamin string             `bson:"jeniskelamin,omitempty" json:"jeniskelamin,omitempty"`
}

type MataKuliah struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode     string             `bson:"kode,omitempty" json:"kode,omitempty"`
	Nama     string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Dosen    string             `bson:"dosen,omitempty" json:"dosen,omitempty"`
	Semester int                `bson:"semester,omitempty" json:"semester,omitempty"`
	Modules  []Module           `bson:"modules,omitempty" json:"modules,omitempty"`
}

type Module struct {
	Nama       string `bson:"nama,omitempty" json:"nama,omitempty"`
	Deskripsi  string `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Materi     string `bson:"materi,omitempty" json:"materi,omitempty"`
	Tugas      string `bson:"tugas,omitempty" json:"tugas,omitempty"`
	Kuis       string `bson:"kuis,omitempty" json:"kuis,omitempty"`
	Referensi  string `bson:"referensi,omitempty" json:"referensi,omitempty"`
}

type Enrollment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Mahasiswa Mahasiswa          `bson:"mahasiswa,omitempty" json:"mahasiswa,omitempty"`
	Kuliah    MataKuliah         `bson:"matakuliah,omitempty" json:"matakuliah,omitempty"`
	Progress  float64            `bson:"progress,omitempty" json:"progress,omitempty"`
	Status    string             `bson:"status,omitempty" json:"status,omitempty"`
}
