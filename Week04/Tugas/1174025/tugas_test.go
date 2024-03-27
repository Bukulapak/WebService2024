package _1174025

import (
	"fmt"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertKaryawanAndGetKaryawanFromPhoneNumber(t *testing.T) {
	// Insert Karyawan
	nama := "John Doe"
	phone_number := "68122221814"
	jabatan := "Staff IT"
	jamKerja := []JamKerja{
		{Durasi: 8, Jam_masuk: "08:00", Jam_keluar: "16:00", Gmt: 7, Hari: []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat"}, Shift: 1, Piket_tim: "Tim A"},
		{Durasi: 8, Jam_masuk: "08:00", Jam_keluar: "16:00", Gmt: 7, Hari: []string{"Sabtu"}, Shift: 1, Piket_tim: "Tim B"},
	}
	hariKerja := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}

	karyawanID, err := InsertKaryawan(nama, phone_number, jabatan, jamKerja, hariKerja)
	if err != nil {
		t.Errorf("Error inserting karyawan: %v", err)
	}
	fmt.Println("Inserted Karyawan ID:", karyawanID)

	// Get Karyawan
	biodata, err := GetKaryawanFromPhoneNumber(phone_number)
	if err != nil {
		t.Errorf("Error getting karyawan: %v", err)
	}
	fmt.Println("Karyawan found:", biodata)
}

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
	hasil, err := InsertPresensi(long, lat, lokasi, phonenumber, checkin, biodata)
	if err != nil {
		t.Errorf("Error inserting presensi: %v", err)
	}
	fmt.Println("Inserted ID:", hasil)
}

func TestInsertPenggajian(t *testing.T) {
	karyawanID := primitive.NewObjectID()
	bulan := "Maret"
	tahun := 2024
	gajiPokok := 5000000.0
	tunjangan := 1000000.0
	potongan := 500000.0
	totalGaji := gajiPokok + tunjangan - potongan
	tanggalBayar := primitive.NewDateTimeFromTime(time.Now().UTC())

	hasil, err := InsertPenggajian(karyawanID, bulan, tahun, gajiPokok, tunjangan, potongan, totalGaji, tanggalBayar)
	if err != nil {
		t.Errorf("Error inserting penggajian: %v", err)
	}
	fmt.Println("Inserted ID:", hasil)
}

func TestInsertHonor(t *testing.T) {
	karyawanID := primitive.NewObjectID()
	pekerjaan := "Pembuatan Website"
	tanggalMulai := primitive.NewDateTimeFromTime(time.Now().UTC())
	tanggalSelesai := primitive.NewDateTimeFromTime(time.Now().AddDate(0, 1, 0).UTC())
	honorPerjam := 150000.0
	totalHonor := honorPerjam * 40.0 // Misalkan 40 jam kerja dalam 1 bulan
	tanggalBayar := primitive.NewDateTimeFromTime(time.Now().UTC())

	hasil, err := InsertHonor(karyawanID, pekerjaan, tanggalMulai, tanggalSelesai, honorPerjam, totalHonor, tanggalBayar)
	if err != nil {
		t.Errorf("Error inserting honor: %v", err)
	}
	fmt.Println("Inserted ID:", hasil)
}

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "68122221814"
	biodata, err := GetKaryawanFromPhoneNumber(phonenumber)
	if err != nil {
		t.Errorf("Error getting karyawan: %v", err)
	}
	fmt.Println("Karyawan found:", biodata)
}
