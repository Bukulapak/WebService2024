package _714220063

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	lokasi := "Rumah"
	phonenumber := "68122221814"
	checkin := "masuk"
	biodata := Karyawan{
		Nama:         "Drake",
		Phone_number: "628456456211",
		Jabatan:      "Rakyat Biasa",
	}
	hasil := InsertPresensi(long, lat, lokasi, phonenumber, checkin, biodata)
	fmt.Println(hasil)
}

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "68122221814"
	biodata := GetKaryawanFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}

func TestGetAll(t *testing.T) {
	data := GetAllPresensi()
	fmt.Println(data)
}
