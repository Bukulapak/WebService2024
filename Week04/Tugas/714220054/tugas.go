package _714220054

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"time"
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

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertRapatMakrab(long float64,lat float64, ukm string, alamat string, kepada Undangan_Rapat, waktu Tanggal) (InsertedID interface{}) {
	var makrab Rapat_Makrab
	makrab.Latitude = long
	makrab.Longitude = lat
	makrab.UKM = ukm
	makrab.Alamat = alamat
	makrab.Kepada = kepada
	makrab.Waktu = waktu
	return InsertOneDoc("undangan_rapat", "rapat_makrab", makrab)
}

func GetAllRapatMakrab() (data []Rapat_Makrab) {
	undangan_rapat := MongoConnect("undangan_rapat").Collection("rapat_makrab")
	filter := bson.M{}
	cursor, err := undangan_rapat.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func InsertTanggal(waktu time.Time) (InsertedID interface{}){
	var tgl Tanggal
	tgl.Waktu = waktu.UTC()

	return InsertOneDoc("undangan_rapat", "rapat_makrab", tgl)
}


func GetAllTanggal() (data []Tanggal) {
	undangan_rapat := MongoConnect("undangan_rapat").Collection("rapat_makrab")
	filter := bson.M{}
	cursor, err := undangan_rapat.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func InsertUndanganRapat(kepada string, divisi string) (InsertedID interface{}){
	var undangan Undangan_Rapat
	undangan.Kepada = kepada
	undangan.Divisi = divisi

	return InsertOneDoc("undangan_rapat", "rapat_makrab", undangan)
}

func GetAllUndanganRapat() (data []Tanggal) {
	undangan_rapat := MongoConnect("undangan_rapat").Collection("rapat_makrab")
	filter := bson.M{}
	cursor, err := undangan_rapat.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
