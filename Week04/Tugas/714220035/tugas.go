package _714220035

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertDosen(d Dosen) interface{} {
	insertResult, err := MongoConnect("perwalianDB").Collection("dosen").InsertOne(context.TODO(), d)
	if err != nil {
		fmt.Printf("InsertDosen: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetDosenByMatkul(matkul string) (d Dosen) {
	col := MongoConnect("perwalianDB").Collection("dosen")
	filter := bson.M{"matkul": matkul}
	err := col.FindOne(context.TODO(), filter).Decode(&d)
	if err != nil {
		fmt.Printf("GetDosenByMatkul: %v\n", err)
	}
	return d
}

func InsertMahasiswa(m Mahasiswa) interface{} {
	insertResult, err := MongoConnect("perwalianDB").Collection("mahasiswa").InsertOne(context.TODO(), m)
	if err != nil {
		fmt.Printf("InsertMahasiswa: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetMahasiswaByNIM(nim string) (m Mahasiswa) {
	col := MongoConnect("perwalianDB").Collection("mahasiswa")
	filter := bson.M{"nim": nim}
	err := col.FindOne(context.TODO(), filter).Decode(&m)
	if err != nil {
		fmt.Printf("GetMahasiswaByNIM: %v\n", err)
	}
	return m
}

func InsertPerwalian(p Perwalian) interface{} {
	insertResult, err := MongoConnect("perwalianDB").Collection("perwalian").InsertOne(context.TODO(), p)
	if err != nil {
		fmt.Printf("InsertPerwalian: %v\n", err)
	}
	return insertResult.InsertedID
}

func GetPerwalianByMatkul(matkul string) (p Perwalian) {
	col := MongoConnect("perwalianDB").Collection("perwalian")
	filter := bson.M{"matkul": matkul}
	err := col.FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		fmt.Printf("GetPerwalianByMatkul: %v\n", err)
	}
	return p
}
