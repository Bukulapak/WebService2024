package _714220050

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"os"
	//"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var MongoString string = os.Getenv("MONGOSTRING")

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

func IsertMahasiswa(nama string, phonenumber string, jurusan string) (InsertedID interface{}) {
	var mahasiswa Mahasiswa
	mahasiswa.Nama = nama
	mahasiswa.Phone_number = phonenumber
	mahasiswa.Jurusan = jurusan
	return InsertOneDoc("kemahasiswaan", "mahasiswa", mahasiswa)
}

func GetMahasiswaFromPhoneNumber(phone_number string) (staf Mahasiswa) {
	mahasiswa:= MongoConnect("kemahasiswaan").Collection("mahasiswa")
	filter := bson.M{"phone_number": phone_number}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getMahasiswaFromPhoneNumber: %v\n", err)
	}
	return staf
}

func GetAllMahasiswa() (data []Mahasiswa) {
	mahasiswa := MongoConnect("kemahasiswaan").Collection("mahasiswa")
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

func InsertMatkul(namaMatkul string, jamMasuk string, hari []string, sks int, dosen string) (InsertedID interface{}) {
    var matkul Matkul
    matkul.Nama_matkul = namaMatkul
   matkul.Jam_masuk = jamMasuk
    matkul.Hari = hari
    matkul.Sks = sks
    matkul.Dosen = dosen
    return InsertOneDoc("kemahasiswaan", "matkul", matkul)
}



func InsertPresensi(phoneNumber string, datetime primitive.DateTime, biodata Mahasiswa) (InsertedID interface{}) {
    var presensi Presensi
    presensi.Phone_number = phoneNumber
    presensi.Datetime = datetime
    presensi.Biodata = biodata
    return InsertOneDoc("kemahasiswaan", "presensi", presensi)
}
