package _714220054

import (
	"fmt"
	"testing"
)

func TestInsertRapatMakrab(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	ukm := "Logic Coffee"
	alamat := "Cafe Logic"
	ket := Undangan_Rapat{
		Judul : "Bahas Villa Mkarab",
	}
	hasil:=InsertPresensi(long ,lat , ukm , alamat , ket )
	fmt.Println(hasil)
}

func TestGetAll(t *testing.T) {
	data := GetAllRapatMakrab()
	fmt.Println(data)
}

