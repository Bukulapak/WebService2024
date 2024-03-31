package _714220023

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertPelanggan(t *testing.T) {
	nama := "Hyunmin "
	phoneNumber := "08123456789"
	alamat := "Jalan Dirgantara No. 50"
	email := []string{"Hyunmin256@example.com", "Hyunmin67@example.com"}
	insertedID := InsertPelanggan(nama, phoneNumber, alamat, email)
	fmt.Println(insertedID)
}

func TestGetPelangganByID(t *testing.T) {
	pelangganID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	pelanggan := GetPelangganByID(pelangganID)
	fmt.Println(pelanggan)
}
func TestGetAllPelanggan(t *testing.T) {
	pelanggans := GetAllPelanggan()
	fmt.Println(pelanggans)
}

func TestInsertProduk(t *testing.T) {
	namaProduk := "Martabak Manis Kacang Coklat"
	deskripsi := "Martabak manis melimpah akan toping"
	harga := 25000
	insertedID := InsertProduk(namaProduk, deskripsi, harga)
	fmt.Println(insertedID)
}

func TestGetProdukByID(t *testing.T) {
	produkID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	produk := GetProdukByID(produkID)
	fmt.Println(produk)
}

func TestGetAllProduk(t *testing.T) {
	produks := GetAllProduk()
	fmt.Println(produks)
}

func TestInsertTransaksi(t *testing.T) {
	metodePembayaran := "Transfer Bank"
	tanggalWaktu := "2024-03-28 10:00:00"
	insertedID := InsertTransaksi(metodePembayaran, tanggalWaktu)
	fmt.Println(insertedID)
}

func TestGetTransaksiByID(t *testing.T) {
	transaksiID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	transaksi := GetTransaksiByID(transaksiID)
	fmt.Println(transaksi)
}
func TestGetAllTransaksi(t *testing.T) {
	transaksis := GetAllTransaksi()
	fmt.Println(transaksis)
}
