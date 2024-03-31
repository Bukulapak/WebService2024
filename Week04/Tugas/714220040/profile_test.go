package _714220040

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertProfile(t *testing.T) {
	nama := "Salzsa"
	npm := "654321"
	prodi := "TI"
	phone_number := "68122221814"
	biodata := Mahasiswa{
		Nama: "Drake",
	}
	hasil := InsertProfile(nama, npm, prodi, phone_number, biodata)
	fmt.Println(hasil)
}

func TestGetProfileByID(t *testing.T) {
	ID, err := primitive.ObjectIDFromHex("7615132")
	if err != nil {
		fmt.Println("Error; %v\n", err)
		return
	}
	profile := GetProfileByID(ID)
	fmt.Println(profile)
}

func TestGetMahasiswaFromID(t *testing.T) {
	ID, err := primitive.ObjectIDFromHex("7652")
	if err != nil {
		fmt.Println("Error; %v\n", err)
		return
	}
	mahasiswa := GetMahasiswaFromID(ID)
	fmt.Println(mahasiswa)
}

func TestGetAll(t *testing.T) {
	data := GetAllProfile()
	fmt.Println(data)
}
