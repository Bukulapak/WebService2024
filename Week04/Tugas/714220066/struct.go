package ws

import (
	"time"
)

type Mahasiswa struct {
	NPM   string `bson:"npm" json:"npm"`
	Name  string `bson:"name" json:"name"`
	Phone string `bson:"phone" json:"phone"`
}

type Presensi struct {
	ID        string    `bson:"id" json:"id"`
	NPM       string    `bson:"npm" json:"npm"`
	Pertemuan string    `bson:"pertemuan" json:"pertemuan"`
	Time      time.Time `bson:"time" json:"time"`
}

type Pertemuan struct {
	ID    string `bson:"id" json:"id"`
	Judul string `bson:"judul" json:"judul"`
}
