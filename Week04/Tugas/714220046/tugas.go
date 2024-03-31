package _714220046

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

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertMataKuliah(kode string, nama string, dosen string, semester int, modules []Module) (insertedID interface{}) {
	var mataKuliah MataKuliah
	mataKuliah.ID = primitive.NewObjectID()
	mataKuliah.Kode = kode
	mataKuliah.Nama = nama
	mataKuliah.Dosen = dosen
	mataKuliah.Semester = semester
	mataKuliah.Modules = modules
	return InsertOneDoc("elearning", "matakuliah", mataKuliah)
}

func GetAllMataKuliah() (mataKuliah []MataKuliah, err error) {
	cursor, err := MongoConnect("elearning").Collection("matakuliah").Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Printf("GetAllMataKuliah: %v\n", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &mataKuliah); err != nil {
		fmt.Printf("GetAllMataKuliah: %v\n", err)
		return nil, err
	}
	return mataKuliah, nil
}

