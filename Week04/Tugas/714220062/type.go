package _714220062

import "go.mongodb.org/mongo-driver/bson/primitive"

// Buku represents a book in the library.
type Buku struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Judul    string             `bson:"judul"`
	Penulis  string             `bson:"penulis"`
	Tahun    int                `bson:"tahun"`
	Jumlah   int                `bson:"jumlah"`
	Dipinjam int                `bson:"dipinjam"`
}

// Mahasiswa represents a student who can borrow books from the library.
type Mahasiswa struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty"`
	Nama          string               `bson:"nama"`
	NomorInduk    string               `bson:"nomor_induk"`
	ProgramStudi  string               `bson:"program_studi"`
	TahunMasuk    int                  `bson:"tahun_masuk"`
	JumlahPinjam  int                  `bson:"jumlah_pinjam"`
	BukuDipinjam  []primitive.ObjectID `bson:"buku_dipinjam"`
}

// Transaksi represents a borrowing transaction of a book by a student.
type Transaksi struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Mahasiswa primitive.ObjectID `bson:"mahasiswa"`
	Buku      primitive.ObjectID `bson:"buku"`
	Tanggal   primitive.DateTime `bson:"tanggal"`
	Kembali   bool               `bson:"kembali"`
}
