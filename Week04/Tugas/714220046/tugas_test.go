package _714220046

import (
	"fmt"
	"testing"
)



func TestInsertMataKuliah(t *testing.T) {
	kode := "AM101"
	nama := "Atom Mining"
	dosen := "Dr. Acula"
	semester := 1
	modules := []Module{
		{
			Nama:      "Atom Mining Basic",
			Materi:    "Materi 1",
			Tugas:     "Tugas 1",
			Deskripsi: "Deskripsi modul 1",
			Kuis:      "Kuis 1",
			Referensi: "Referensi 1",
		},
		{
			Nama:      "Atom Mining Advanced",
			Materi:    "Materi 2",
			Tugas:     "Tugas 2",
			Deskripsi: "Deskripsi modul 2",
			Kuis:      "Kuis 2",
			Referensi: "Referensi 2",
		},
	}
	hasil := InsertMataKuliah(kode, nama, dosen, semester, modules)
	fmt.Println(hasil)
}

func TestGetAllMataKuliah(t *testing.T) {
	mataKuliah, err := GetAllMataKuliah()
	if err != nil {
		t.Errorf("Error fetching Mata Kuliah: %v", err)
		return
	}

	fmt.Println("Data Mata Kuliah:")
	for _, mk := range mataKuliah {
		fmt.Printf("Kode: %s, Nama: %s, Dosen: %s, Semester: %d\n", mk.Kode, mk.Nama, mk.Dosen, mk.Semester)
	}
}
