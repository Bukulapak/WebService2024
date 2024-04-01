package _714220062

import (
	"context"
	"fmt"
	"os"

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

func InsertBuku(judul string, penulis string, tahun int, jumlah int, dipinjam int) (InsertedID interface{}) {
	var buku Buku
	buku.ID = primitive.NewObjectID()
	buku.Judul = judul
	buku.Penulis = penulis
	buku.Tahun = tahun
	buku.Jumlah = jumlah
	buku.Dipinjam = dipinjam
	return InsertOneDoc("tugasweek04", "buku", buku)
}

func GetBukuByID(id primitive.ObjectID) (buku Buku) {
	b := MongoConnect("tugasweek04").Collection("buku")
	filter := bson.M{"_id": id}
	err := b.FindOne(context.TODO(), filter).Decode(&buku)
	if err != nil {
		fmt.Printf("GetBukuByID: %v\n", err)
	}
	return buku
}

func GetAllBuku() (data []Buku) {
	b := MongoConnect("tugasweek04").Collection("buku")
	filter := bson.M{}
	cursor, err := b.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetAllBuku:", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
