// praktikum_test.go
package _714220031

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertResponden(t *testing.T) {
	usia := "20"
	JenisKelamin := "Perempuan"
	pekerjaan := "PNS"
	insertedID := InsertResponden(usia, JenisKelamin, pekerjaan)
	fmt.Println(insertedID)
}

func TestGetRespondenByID(t *testing.T) {
	respondenID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	responden := GetRespondenByID(respondenID)
	fmt.Println(responden)
}

func TestGetAllResponden(t *testing.T) {
	respondens := GetAllResponden()
	fmt.Println(respondens)
}

func TestInsertPertanyaan(t *testing.T) {
	NamaPertanyaan := "Bagaimana pendapat anda tentang perkembangan zaman?"
	insertedID := InsertPertanyaan(NamaPertanyaan)
	fmt.Println(insertedID)
}

func TestGetPertanyaanByID(t *testing.T) {
	pertanyaanID, err := primitive.ObjectIDFromHex("615af14ae62f4c488e1d6d14")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	pertanyaan := GetPertanyaanByID(pertanyaanID)
	fmt.Println(pertanyaan)
}

func TestGetAllPertanyaan(t *testing.T) {
	pertanyaans := GetAllPertanyaan()
	fmt.Println(pertanyaans)
}

func TestInsertJawaban(t *testing.T) {
	NamaJawaban := "Perkembangan zaman adalah hal yang alami dan tak terelakkan, membawa perubahan dalam berbagai aspek kehidupan."
	TanggalJawab := "2024-03-30"
	insertedID := InsertJawaban(NamaJawaban, TanggalJawab)
	fmt.Println(insertedID)
}
