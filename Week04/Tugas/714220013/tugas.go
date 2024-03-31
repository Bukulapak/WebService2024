package _714220013

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"

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

func InsertJadwal(hari string, waktuMulai string, waktuSelesai string, mataKuliah MataKuliah, ruangan int, dosen Dosen) (InsertedID interface{}) {
	var jadwal Jadwal
	jadwal.Hari = hari
	jadwal.Waktu_Mulai = waktuMulai
	jadwal.Waktu_Selesai = waktuSelesai
	jadwal.Mata_Kuliah = mataKuliah
	jadwal.Ruangan = ruangan
	jadwal.Dosen = dosen

	return InsertOneDoc("week4", "jadwal", jadwal)
}

func GetJadwalByMataKuliah(matkul string) (jdw Jadwal, err error) {
	jadwal := MongoConnect("week4").Collection("jadwal")
	filter := bson.M{"mata_kuliah.nama_matkul": matkul}
	err = jadwal.FindOne(context.TODO(), filter).Decode(&jdw)
	if err != nil {
		fmt.Printf("GetJadwalByMataKuliah: %v\n", err)
	}
	return jdw, err
}

func GetDosenByWaktuMulai(wmulai string) (dsn Dosen, err error) {
	dosen := MongoConnect("week4").Collection("jadwal")
	filter := bson.M{"waktu_mulai": wmulai}
	err = dosen.FindOne(context.TODO(), filter).Decode(&dsn)
	if err != nil {
		fmt.Printf("GetJadwalByWaktuMulai: %v\n", err)
	}
	return dsn, err
}

func GetMataKuliahByNamaDosen(dosen string) (mk MataKuliah, err error) {
	jadwal := MongoConnect("week4").Collection("jadwal")
	filter := bson.M{"dosen.nama_dosen": dosen}
	err = jadwal.FindOne(context.TODO(), filter).Decode(&mk)
	if err != nil {
		fmt.Printf("GetMataKuliahByNamaDosen: %v\n", err)
	}
	return mk, err
}

func GetAllJadwal() (data []Jadwal) {
	jadwal := MongoConnect("week4").Collection("jadwal")
	filter := bson.M{}
	cursor, err := jadwal.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
