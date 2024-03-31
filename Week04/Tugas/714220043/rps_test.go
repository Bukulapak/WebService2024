package _14220043

import (
	"fmt"
	"testing"
)

func TestInsertRPS(t *testing.T) {
	matakuliah := Matakuliah{
		Nama_Mata_kuliah: "RPL2",
		Deskripsi:        "Mempelajari bahasa pemrograman",
		Jumlah_SKS:       "7",
		Semester:         "3",
		Kurikulum: Kurikulum{
			Nama_Kurikulum: "Kampus Merdeka",
			Deskripsi:      "Meningkatkan pembelajaran",
			Tahun_Mulai:    "2023",
			Tahun_Selesai:  "2024",
		},
	}
	deskripsi := "Semangat Belajar"
	tujuan := "Untuk meningkatkan kualitas mahasiswa"

	hasil := InsertRPS(matakuliah, deskripsi, tujuan)
	fmt.Println(hasil)
}

func TestGetRPSFromNamaKRKLM(t *testing.T) {
	rps, err := GetRPSFromNamaKRKLM("Kampus Merdeka")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Deskripsi:", rps.Deskripsi)
	}
}

func TestGetTujuanFromNamaMataKuliah(t *testing.T) {
	tujuan, err := GetTujuanFromNamaMataKuliah("RPL2")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Tujuan:", tujuan)
	}

}

func TestGetMataKuliahFromDeskripsi(t *testing.T) {
	// Menentukan deskripsi mata kuliah yang ingin diuji
	deskripsiMataKuliah := "Mempelajari bahasa pemrograman" // Ganti dengan deskripsi yang sesuai dengan data Anda

	// Memanggil fungsi yang akan diuji
	namaMataKuliah, err := GetMataKuliahFromDeskripsi(deskripsiMataKuliah)

	// Memeriksa apakah terjadi kesalahan saat memanggil fungsi
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}

	// Menampilkan output sesuai format yang diinginkan
	fmt.Printf("Nama Mata Kuliah: %s\n", namaMataKuliah)

}

func TestGetAllRPS(t *testing.T) {
	data := GetAllRPS()
	fmt.Println(data)
}
