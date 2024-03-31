package _714220013

import (
	"fmt"
	"testing"
)

func TestInsertJadwal(t *testing.T) {
	hari := "Selasa"
	waktuMulai := "13.00"
	waktuSelesai := "17.00"
	mataKuliah := MataKuliah{
		Kode_Matkul: "A0102",
		Nama_Matkul: "SAP",
		Sks:         3,
		Semester:    4,
	}
	ruangan := 316
	dosen := Dosen{
		Nidn:       "N992002",
		Nama_Dosen: "Nisa Hanum",
	}

	hasil := InsertJadwal(hari, waktuMulai, waktuSelesai, mataKuliah, ruangan, dosen)
	fmt.Println(hasil)
}

func TestGetDosenByWaktuMulai(t *testing.T) {
	dsn, err := GetDosenByWaktuMulai("13.00")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Nama Dosen:", dsn.Nama_Dosen)
	}
}

func TestGetJadwalByMataKuliah(t *testing.T) {
	jdw, err := GetJadwalByMataKuliah("SAP")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Jadwal:", jdw.Hari)
	}
}

func TestGetMataKuliahByNamaDosen(t *testing.T) {
	mk, err := GetMataKuliahByNamaDosen("Nisa Hanum")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Mata Kuliah:", mk.Nama_Matkul)
	}
}

func TestGetAllJadwal(t *testing.T) {
	data := GetAllJadwal()
	fmt.Println(data)
}
