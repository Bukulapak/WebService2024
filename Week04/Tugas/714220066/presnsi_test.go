package ws

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestInsertPresensi(t *testing.T) {
	currentTime := time.Now()
	data := Presensi{
		ID:        "PT67-714220066",
		NPM:       "714220066",
		Pertemuan: "714220066",
		Time:      currentTime,
	}

	err := InsertPresensi(data)
	if err != nil {
		t.Errorf("InsertMhs failed: %v", err)
	}
}

func TestGetPresensi(t *testing.T) {
	filter := bson.M{}
	result, err := GetPresensi(filter)
	if err != nil {
		t.Errorf("GetPresensi failed: %v", err)
	}

	fmt.Println("Presensi data:")
	for _, m := range result {
		fmt.Printf("NPM: %s, Pertemuan: %s, Time: %s\n", m.NPM, m.Pertemuan, m.Time)
	}
}
