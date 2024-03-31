package ws

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

type MockMongoDB struct{}

func TestInsertMhs(t *testing.T) {
	mockMahasiswa := Mahasiswa{
		NPM:   "714220066",
		Name:  "Alif Fathur",
		Phone: "123456789",
	}

	err := InsertMhs(mockMahasiswa)
	if err != nil {
		t.Errorf("InsertMhs failed: %v", err)
	}
}

func TestGetMhs(t *testing.T) {
	filter := bson.M{}
	result, err := GetMhs(filter)
	if err != nil {
		t.Errorf("GetMhs failed: %v", err)
	}

	fmt.Println("Mahasiswa data:")
	for _, m := range result {
		fmt.Printf("NPM: %s, Name: %s, Phone: %s\n", m.NPM, m.Name, m.Phone)
	}
}
