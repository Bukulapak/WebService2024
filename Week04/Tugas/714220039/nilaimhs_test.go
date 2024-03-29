package _14220039

import (
	"fmt"
	"testing"
)

func TestInsertNilai(t *testing.T) {
	nilai := "85"
	matkul := "Pemograman2"
	phonenumber := "6284455621"
	biodata := Mahasiswa{
		Nama:         "Agung2",
		Phone_number: "628456456211",
		Kelas:        "2B",
		Alamat: Alamat{
			Jalan:    "Jalan Parapat",
			No_rumah: "30",
			RT:       "09",
			RW:       "03",
			KodePos:  21123,
		},
	}
	hasil := InsertNilai(nilai, matkul, phonenumber, biodata)
	fmt.Println(hasil)
}

func TestGetMahasiswaFromPhoneNumber(t *testing.T) {
	phonenumber := "6284455621"
	biodata := GetMahasiswaFromPhone(phonenumber)
	fmt.Println(biodata)
}
func TestGetNilaiFromNama(t *testing.T) {
	mhs, err := GetNilaiFromNama("Agung")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Nilai Mahasiswa:", mhs.Nilai_MHS)
	}
}

func TestGetNamaFromAlamat(t *testing.T) {
	mhs, err := GetNamaFromAlamat("Jalan Parapat")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Nama Mahasiswa:", mhs.Biodata.Nama)
	}
}

func TestGetAll(t *testing.T) {
	data := GetAllNilai()
	fmt.Println(data)
}
