package _714220061

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

//  koneksi

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

// insert data pemmbeli

func InsertPembeli(nama string, nomortelepon string, alamat string) (InsertedID interface{}) {
	var pembeli DataPembeli
	pembeli.Nama = nama
	pembeli.NomorTelepon = nomortelepon
	pembeli.Alamat = alamat
	return InsertOneDoc("InsertData", "pembeli", pembeli)
}

// insert data tagihan

func InsertTagihan(nama DataPembeli, total float64) (InsertedID interface{}) {
	var tagihan DataTagihan
	tagihan.Nama = nama
	tagihan.Tanggal = primitive.NewDateTimeFromTime(time.Now().UTC())
	tagihan.Total = total
	return InsertOneDoc("InsertData", "tagihan", tagihan)
}

// insert data barang

func InsertBarang(namabarang string, harga float64) (InsertID interface{}) {
	var barang DataBarang
	barang.NamaBarang = namabarang
	barang.HargaBarang = harga
	return InsertOneDoc("InsertData", "barang", barang)
}

// get all data pembeli

func GetAllPembeli() (data []DataPembeli) {
	pembeli := MongoConnect("GetAllData").Collection("pembeli")
	filter := bson.M{}
	cursor, err := pembeli.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// get all data tagihan

func GetAllTagihan() (data []DataTagihan) {
	tagihan := MongoConnect("GetAllData").Collection("tagihan")
	filter := bson.M{}
	cursor, err := tagihan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

//get all data barang

func GetAllBarang() (data []DataBarang) {
	barang := MongoConnect("GetAllData").Collection("barang")
	filter := bson.M{}
	cursor, err := barang.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
