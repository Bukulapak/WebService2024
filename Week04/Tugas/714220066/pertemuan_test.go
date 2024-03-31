package ws

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestInsertPertemuan(t *testing.T) {
	data := Pertemuan{
		ID:    "PT67",
		Judul: "Aritmatika Dasar",
	}

	err := InsertPertemuan(data)
	if err != nil {
		t.Errorf("InsertMhs failed: %v", err)
	}
}

func TestGetPertemuan(t *testing.T) {
	filter := bson.M{}
	result, err := GetPertemuan(filter)
	if err != nil {
		t.Errorf("GetPertemuan failed: %v", err)
	}

	fmt.Println("Pertemuan data:")
	for _, m := range result {
		fmt.Printf("ID: %s, Judul: %s\n", m.ID, m.Judul)
	}
}
