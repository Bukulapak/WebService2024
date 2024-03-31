package _714220054

import (
	"fmt"
	"testing"
	"time"
)

func TestInsertRapatMakrab(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	ukm := "Logic Coffee"
	alamat := "Cafe Logic"
	kepada := Undangan_Rapat{
		Kepada : "Devi Wulandari",
	}
	waktu := Tanggal{
		Waktu : time.Date(2024, time.March, 30, 10, 0, 0, 0, time.UTC),
	}
	hasil:=InsertRapatMakrab(long ,lat , ukm , alamat , kepada , waktu)
	fmt.Println(hasil)
}

func TestInsertTanggal (t *testing.T){
    waktu := time.Date(2024, time.March, 30, 10, 0, 0, 0, time.UTC)
	hasil:=InsertTanggal(waktu)
	fmt.Println(hasil)
}

func TestInsertUndanganRapat(t *testing.T){
    kepada := "Devi Wulandari"
	divisi := "Humas"
	hasil:=InsertUndanganRapat(kepada, divisi)
	fmt.Println(hasil)
}

func TestGetAll(t *testing.T) {
	data := GetAllRapatMakrab()
	fmt.Println(data)
}

