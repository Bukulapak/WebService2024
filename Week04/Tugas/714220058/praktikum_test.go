package _714220058

import (
	"fmt"
	"testing"
)

func TestInsertPendaftaran(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	alamat := "Garut"
	nomorTelepon := "082127854156"
	status := "Mahasiswa Baru"
	biodata := Mahasiswa{
		Nama:         "Muhammad Qinthar",
		NomorTelepon: "082127854156",
		ProgramStudi: "D4 Manajemen Informatika",
	}
	hasil := InsertPendaftaran(long, lat, alamat, nomorTelepon, status, biodata)
	fmt.Println(hasil)
}

func TestGetMahasiswaByNomorTelepon(t *testing.T) {
	nomorTelepon := "082127854156"
	biodata := GetMahasiswaByNomorTelepon(nomorTelepon)
	fmt.Println(biodata)
}

func TestGetAllPendaftaran(t *testing.T) {
	data := GetAllPendaftaran()
	fmt.Println(data)
}

func TestGetAllMahasiswa(t *testing.T) {
	data := GetAllMahasiswa()
	fmt.Println(data)
}

func TestGetAllJamKuliah(t *testing.T) {
	data := GetAllJamKuliah()
	fmt.Println(data)
}