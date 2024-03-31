package _714220058

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(dbname string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(dbname).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPendaftaran(long float64, lat float64, alamat string, nomorTelepon string, status string, biodata Mahasiswa) (InsertedID interface{}) {
	var pendaftaran Pendaftaran
	pendaftaran.ID = primitive.NewObjectID()
	pendaftaran.Latitude = lat
	pendaftaran.Longitude = long
	pendaftaran.Alamat = alamat
	pendaftaran.NomorTelepon = nomorTelepon
	pendaftaran.TanggalDaftar = primitive.NewDateTimeFromTime(time.Now().UTC())
	pendaftaran.Status = status
	pendaftaran.BiodataMahasiswa = biodata
	return InsertOneDoc("tugasweek04", "pmb", pendaftaran)
}

func GetMahasiswaByNomorTelepon(nomorTelepon string) (mahasiswa Pendaftaran) {
	pendaftaran := MongoConnect("tugasweek04").Collection("pmb")
	filter := bson.M{"nomor_telepon": nomorTelepon}
	err := pendaftaran.FindOne(context.TODO(), filter).Decode(&mahasiswa)
	if err != nil {
		fmt.Printf("GetMahasiswaByNomorTelepon: %v\n", err)
	}
	return mahasiswa
}


func GetAllPendaftaran() (data []Pendaftaran) {
	pendaftaran := MongoConnect("tugasweek04").Collection("pmb")
	filter := bson.M{}
	cursor, err := pendaftaran.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllPendaftaran:", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllMahasiswa() (data []Pendaftaran) {
	mahasiswa := MongoConnect("tugasweek04").Collection("pmb")
	filter := bson.M{}
	cursor, err := mahasiswa.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllMahasiswa:", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllJamKuliah() (data []Pendaftaran) {
	jamkuliah := MongoConnect("tugasweek04").Collection("pmb")
	filter := bson.M{}
	cursor, err := jamkuliah.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllJamKuliah:", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}