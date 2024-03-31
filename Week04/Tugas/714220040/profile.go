package _714220040

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

func InsertProfile(nama string, npm string, prodi string, phonenumber string, biodata Mahasiswa) (InsertedID interface{}) {
	var profile Profile
	profile.Nama = nama
	profile.Npm = npm
	profile.Prodi = prodi
	profile.Phone_number = phonenumber
	profile.Biodata = biodata
	return InsertOneDoc("Profile", "profile", profile)
}

func GetProfileByID(ID primitive.ObjectID) (profile Profile) {
	collection := MongoConnect("Profile").Collection("profile")
	filter := bson.M{"_id": ID}
	err := collection.FindOne(context.TODO(), filter).Decode(&profile)
	if err != nil {
		fmt.Printf("getProfileByID: %v\n", err)
	}
	return profile
}

func GetAllProfile() (data []Profile) {
	mahasiswa := MongoConnect("Profile").Collection("profile")
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

func InsertMahasiswa(nama string, phone_number string, prodi string, data_ortu []DataOrtu, NamaAyah []string) (InsertedID interface{}) {
	var mahasiswa Mahasiswa
	mahasiswa.Nama = nama
	mahasiswa.Phone_number = phone_number
	mahasiswa.Prodi = prodi
	return InsertOneDoc("Profile", "mahasiswa", mahasiswa)
}

func GetMahasiswaFromID(ID primitive.ObjectID) (mhs Profile) {
	mahasiswa := MongoConnect("Profile").Collection("mahasiswa")
	filter := bson.M{"_id": ID}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("getMahasiswaFromPhoneNumber: %v\n", err)
	}
	return mhs
}

func GetAllMahasiswa() (data []Profile) {
	mahasiswa := MongoConnect("Profile").Collection("mahasiswa")
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

func InsertDataOrtu(namaAyah string, pekerjaanAyah string, namaIbu string, pekerjaanIbu string) (InsertedID interface{}) {
	var dataOrtu DataOrtu
	dataOrtu.NamaAyah = namaAyah
	dataOrtu.PekerjaanAyah = pekerjaanAyah
	dataOrtu.NamaIbu = namaIbu
	dataOrtu.PekerjaanIbu = pekerjaanIbu
	return InsertOneDoc("Profile", "DataOrtu", dataOrtu)
}

func GetDataOrtuFromNamaIbu(namaIbu string) (ort DataOrtu) {
	ortu := MongoConnect("Profile").Collection("dataOrtu")
	filter := bson.M{"namaIbu": namaIbu}
	err := ortu.FindOne(context.TODO(), filter).Decode(&ort)
	if err != nil {
		fmt.Printf("getDataOrtuFromNamaIbu: %v\n", err)
	}
	return ort
}

func GetAllDataOrtu() (data []Profile) {
	mahasiswa := MongoConnect("Profile").Collection("dataOrtu")
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
