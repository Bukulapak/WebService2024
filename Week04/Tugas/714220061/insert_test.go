package _714220061

import (
	"fmt"
	"testing"
)

func TestInsertPembeli(t *testing.T) {
	nama := "Mikazu"
	nomortelepon := "082293392304"
	alamat := "Shibakoen, Minato City, Tokyo"
	hasil := InsertPembeli(nama, nomortelepon, alamat)
	fmt.Println(hasil)
}

func TestInsertTagihan(t *testing.T) {
	nama := DataPembeli{
		Nama: "Mikazu",
	}
	total := 23000000.0
	hasil := InsertTagihan(nama, total)
	fmt.Println(hasil)
}

func TestInsertBarang(t *testing.T) {
	namabarang := "Macbook Pro M2"
	harga := 23000000.0
	hasil := InsertBarang(namabarang, harga)
	fmt.Println(hasil)
}
