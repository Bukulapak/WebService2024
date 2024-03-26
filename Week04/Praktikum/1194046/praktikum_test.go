package _1194046

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	lokasi := "Rumah"
	phonenumber := "62897979797"
	checkin := "masuk"
	biodata := Karyawan{
		Nama:         "Ilham Ambar",
		Phone_number: "62897979797",
		Jabatan:      "Rakyat Biasa",
	}
	hasil := InsertPresensi(long, lat, lokasi, phonenumber, checkin, biodata)
	fmt.Println(hasil)
}

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "62897979797"
	biodata := GetKaryawanFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}

func TestGetAll(t *testing.T) {
	data := GetAllPresensi()
	fmt.Println(data)
}
