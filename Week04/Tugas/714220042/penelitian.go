package _714220042

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

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}


func InsertPenelitian(judul string, institusi string, penulis string, ringkasan string, biodata Peneliti) (insertedID interface{}) {
	penelitian := penelitian{
		Judul:     judul,
		Institusi: institusi,
		Penulis:   penulis,
		Datetime:  primitive.NewDateTimeFromTime(time.Now()),
		Ringkasan: ringkasan,
		Biodata:   biodata,
	}
	return InsertOneDoc("tugas04_penelitian", "penelitian", penelitian)
}
func GetPenelitiFromInstitusi(institusi string) (penelitian penelitian, err error) {
    collection := MongoConnect("tugas04_penelitian").Collection("penelitian")
    filter := bson.M{"institusi": institusi}
    err = collection.FindOne(context.TODO(), filter).Decode(&penelitian)
    if err != nil {
        fmt.Printf("GetPenelitiFromInstitusi: %v\n", err)
    }
    return penelitian, err
}

func GetPenelitianFromNama(nama string) (penelitian penelitian, err error) {
    collection := MongoConnect("tugas04_penelitian").Collection("penelitian")
    filter := bson.M{"biodata.nama": nama}
    err = collection.FindOne(context.TODO(), filter).Decode(&penelitian)
    if err != nil {
        fmt.Printf("GetPenelitianFromNama: %v\n", err)
    }
    return penelitian, err
}

func GetNamaFromPublikasi(publikasi string) (Penelitian penelitian, err error) {
    collection := MongoConnect("tugas04_penelitian").Collection("penelitian")
    filter := bson.M{"biodata.publikasi.judul": publikasi}
    err = collection.FindOne(context.TODO(), filter).Decode(&Penelitian)
    if err != nil {
        fmt.Printf("GetNamaFromPublikasi: %v\n", err)
    }
    return Penelitian, err
}

func GetAllPenelitian() (data []penelitian) {
	collection := MongoConnect("tugas04_penelitian").Collection("penelitian")
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetAllPenelitian: %v\n", err)
		return data
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var penelitian penelitian
		err := cursor.Decode(&penelitian)
		if err != nil {
			fmt.Printf("Error decoding penelitian: %v\n", err)
			continue
		}
		data = append(data, penelitian)
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("Cursor error: %v\n", err)
	}
	return data
}