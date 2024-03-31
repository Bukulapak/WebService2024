package _714220035

import (
	"fmt"
	"testing"
)

func TestInsertAndGetDosen(t *testing.T) {
	d := Dosen{
		Nama:   "Dr. Samsud",
		NIP:    "123456",
		Matkul: "Machine Learning",
	}
	insertedID := InsertDosen(d)
	fmt.Println("Inserted Dosen ID:", insertedID)

	matkul := "Machine Learning"
	dosen := GetDosenByMatkul(matkul)
	fmt.Println("Dosen:", dosen)
}

func TestInsertAndGetMahasiswa(t *testing.T) {
	m := Mahasiswa{
		Nama:     "Irgi F",
		NIM:      "714220035",
		Semester: 4,
	}
	insertedID := InsertMahasiswa(m)
	fmt.Println("Inserted Mahasiswa ID:", insertedID)

	nim := "714220035"
	mahasiswa := GetMahasiswaByNIM(nim)
	fmt.Println("Mahasiswa:", mahasiswa)
}

func TestInsertAndGetPerwalian(t *testing.T) {
	dosen := Dosen{
		Nama:   "Dr. Samsud",
		NIP:    "123456",
		Matkul: "Machine Learning",
	}

	mahasiswa := Mahasiswa{
		Nama:     "Irgi F",
		NIM:      "714220035",
		Semester: 4,
	}

	perwalian := Perwalian{
		Mahasiswa: mahasiswa,
		Dosen:     dosen,
		Matkul:    "Machine Learning",
		Status:    "PENDING",
	}

	insertedID := InsertPerwalian(perwalian)
	fmt.Println("Inserted Perwalian ID:", insertedID)

	perwalianFromDB := GetPerwalianByMatkul("Machine Learning")
	fmt.Println("Perwalian:", perwalianFromDB)
}
