package _714220062

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertBuku(t *testing.T) {
	id := InsertBuku("Judul Buku", "Penulis Buku", 2022, 10, 2)
	if id == nil {
		t.Error("Expected ID, got nil")
	}
}

func TestGetBukuByID(t *testing.T) {
	id := InsertBuku("Judul Buku", "Penulis Buku", 2022, 10, 2).(primitive.ObjectID)
	buku := GetBukuByID(id)
	if buku.ID != id {
		t.Errorf("Expected ID: %v, got ID: %v", id, buku.ID)
	}
}

func TestGetAllBuku(t *testing.T) {
	buku := GetAllBuku()
	if len(buku) == 0 {
		t.Error("Expected non-zero length, got 0")
	}
}
