package _14220043

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Kurikulum struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Kurikulum string             `bson:"nama_kurikulum,omitempty" json:"nama_kurikulum,omitempty"`
	Deskripsi      string             `bson:"deskripsi" json:"deskripsi"`
	Tahun_Mulai    string             `bson:"tahun_mulai,omitempty" json:"tahun_mulai,omitempty"`
	Tahun_Selesai  string             `bson:"tahun_selesai,omitempty" json:"tahun_selesai,omitempty"`
}

type Matakuliah struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Mata_kuliah string             `bson:"nama_matakuliah,omitempty" json:"nama_matakuliah,omitempty"`
	Deskripsi        string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty"`
	Jumlah_SKS       string             `bson:"jumlah_sks,omitempty" json:"jumlah_sks,omitempty"`
	Semester         string             `bson:"semester,omitempty" json:"semester,omitempty"`
	Kurikulum        Kurikulum          `bson:"kurikulum,omitempty" json:"kurikulum,omitempty"`
}

type RPS struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Mata_Kuliah Matakuliah         `bson:"mata_kuliah,omitempty" json:"mata_kuliah,omitempty"`
	Deskripsi   string             `bson:"deskripsi,omitempty" json:"deskripi,omitempty"`
	Tujuan      string             `bson:"tujuan,omitempty" json:"tujuan,omitempty"`
}
