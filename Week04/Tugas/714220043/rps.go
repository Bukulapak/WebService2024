package _14220043

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"

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

func InsertRPS(mata_kuliah Matakuliah, deskripi string, tujuan string) (InsertedID interface{}) {
	var rps RPS
	rps.Mata_Kuliah = mata_kuliah
	rps.Deskripsi = deskripi
	rps.Tujuan = tujuan
	return InsertOneDoc("t4", "rps", rps)
}

func GetRPSFromNamaKRKLM(nama_kurikulum string) (rps RPS, err error) {
	filter := bson.M{"mata_kuliah.kurikulum.nama_kurikulum": nama_kurikulum}

	err = MongoConnect("t4").Collection("rps").FindOne(context.TODO(), filter).Decode(&rps)
	if err != nil {
		fmt.Printf("GetRPSFromNamaKRKLM: %v\n", err)
	}
	return rps, err
}

func GetTujuanFromNamaMataKuliah(namaMataKuliah string) (tujuan string, err error) {
	filter := bson.M{"mata_kuliah.nama_matakuliah": namaMataKuliah}

	var rps RPS
	err = MongoConnect("t4").Collection("rps").FindOne(context.TODO(), filter).Decode(&rps)
	if err != nil {
		fmt.Printf("GetTujuanFromNamaMataKuliah: %v\n", err)
		return "", err
	}
	return rps.Tujuan, nil
}

func GetMataKuliahFromDeskripsi(deskripsi string) (namaMataKuliah string, err error) {
	filter := bson.M{"mata_kuliah.deskripsi": deskripsi}

	var rps RPS
	err = MongoConnect("t4").Collection("rps").FindOne(context.TODO(), filter).Decode(&rps)
	if err != nil {
		fmt.Printf("GetMataKuliahFromDeskripsi: %v\n", err)
		return "", err
	}
	return rps.Mata_Kuliah.Nama_Mata_kuliah, nil
}

func GetAllRPS() (data []RPS) {
	rps := MongoConnect("t4").Collection("rps")
	filter := bson.M{}
	cursor, err := rps.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
