package _1174025

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var MongoString string = os.Getenv("MONGOSTRING")

// Fungsi untuk menghubungkan ke MongoDB
func MongoConnect(dbname string) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(MongoString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database(dbname)
	collection := db.Collection("presensi")

	return collection, nil
}

func InsertKaryawan(nama string, phone_number string, jabatan string, jamKerja []JamKerja, hariKerja []string) (interface{}, error) {
	collection, err := MongoConnect("tesdb2024")
	if err != nil {
		return primitive.NilObjectID, err
	}

	karyawan := Karyawan{
		ID:           primitive.NewObjectID(),
		Nama:         nama,
		Phone_number: phone_number,
		Jabatan:      jabatan,
		Jam_kerja:    jamKerja,
		Hari_kerja:   hariKerja,
	}

	insertResult, err := collection.InsertOne(context.Background(), karyawan)
	if err != nil {
		return nil, err
	}

	return insertResult.InsertedID, nil
}

// Fungsi untuk menyisipkan data presensi ke MongoDB
func InsertPresensi(long float64, lat float64, lokasi string, phonenumber string, checkin string, biodata Karyawan) (interface{}, error) {
	collection, err := MongoConnect("tesdb2024")
	if err != nil {
		return nil, err
	}

	presensi := Presensi{
		ID:           primitive.NewObjectID(),
		Longitude:    long,
		Latitude:     lat,
		Location:     lokasi,
		Phone_number: phonenumber,
		Datetime:     primitive.NewDateTimeFromTime(time.Now().UTC()),
		Checkin:      checkin,
		Biodata:      biodata,
	}

	insertResult, err := collection.InsertOne(context.Background(), presensi)
	if err != nil {
		return nil, err
	}

	return insertResult.InsertedID, nil
}

// Fungsi untuk mengambil data karyawan dari MongoDB berdasarkan nomor telepon
func GetKaryawanFromPhoneNumber(phone_number string) (*Presensi, error) {
	collection, err := MongoConnect("tesdb2024")
	if err != nil {
		return nil, err
	}

	filter := bson.M{"biodata.phone_number": phone_number}
	var result Presensi
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Fungsi untuk menyisipkan data penggajian ke MongoDB
func InsertPenggajian(karyawanID primitive.ObjectID, bulan string, tahun int, gaji_pokok, tunjangan, potongan, total_gaji float64, tanggalBayar primitive.DateTime) (interface{}, error) {
	collection, err := MongoConnect("tesdb2024")
	if err != nil {
		return nil, err
	}

	penggajian := Penggajian{
		ID:            primitive.NewObjectID(),
		KaryawanID:    karyawanID,
		Bulan:         bulan,
		Tahun:         tahun,
		Gaji_pokok:    gaji_pokok,
		Tunjangan:     tunjangan,
		Potongan:      potongan,
		Total_gaji:    total_gaji,
		Tanggal_bayar: tanggalBayar,
	}

	insertResult, err := collection.InsertOne(context.Background(), penggajian)
	if err != nil {
		return nil, err
	}

	return insertResult.InsertedID, nil
}

// Fungsi untuk menyisipkan data honor ke MongoDB
func InsertHonor(karyawanID primitive.ObjectID, pekerjaan string, tanggalMulai, tanggalSelesai primitive.DateTime, honorPerjam, totalHonor float64, tanggalBayar primitive.DateTime) (interface{}, error) {
	collection, err := MongoConnect("tesdb2024")
	if err != nil {
		return nil, err
	}

	honor := Honor{
		ID:              primitive.NewObjectID(),
		KaryawanID:      karyawanID,
		Pekerjaan:       pekerjaan,
		Tanggal_mulai:   tanggalMulai,
		Tanggal_selesai: tanggalSelesai,
		Honor_perjam:    honorPerjam,
		Total_honor:     totalHonor,
		Tanggal_bayar:   tanggalBayar,
	}

	insertResult, err := collection.InsertOne(context.Background(), honor)
	if err != nil {
		return nil, err
	}

	return insertResult.InsertedID, nil
}
