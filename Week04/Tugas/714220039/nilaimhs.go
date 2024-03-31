package _14220039

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func InsertNilai(nil string, matkul string, phonenumber string, biodata Mahasiswa) (InsertedID interface{}) {
	var nilai Nilai
	nilai.Nilai_MHS = nil
	nilai.Mata_Kuliah = matkul
	nilai.Phone_number = phonenumber
	nilai.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	nilai.Biodata = biodata
	return InsertOneDoc("tugas4", "nilai", nilai)
}

func GetMahasiswaFromPhone(phone_number string) (mhs Nilai) {
	mahasiswa := MongoConnect("tugas4").Collection("nilai")
	filter := bson.M{"phone_number": phone_number}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("GetMahasiswaFromPhone: %v\n", err)
	}
	return mhs
}

func GetNilaiFromNama(nama string) (mhs Nilai, err error) {
	nilai := MongoConnect("tugas4").Collection("nilai")
	filter := bson.M{"biodata.nama": nama}
	err = nilai.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("GetNilaiFromNama: %v\n", err)
	}
	return mhs, err
}

func GetNamaFromAlamat(alamat string) (mhs Nilai, err error) {
	nilai := MongoConnect("tugas4").Collection("nilai")
	filter := bson.M{"biodata.alamat.jalan": alamat}
	err = nilai.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("GetNamaFromAlamat: %v\n", err)
	}
	return mhs, err
}

func GetAllNilai() (data []Nilai) {
	mahasiswa := MongoConnect("tugas4").Collection("nilai")
	filter := bson.M{}
	cursor, err := mahasiswa.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
