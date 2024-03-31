package _714220023

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
	"os"
	//"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
		return nil
	}
	return client.Database(dbname)
}

func InsertOneDoc(dbname, collection string, doc interface{}) interface{} {
	insertResult, err := MongoConnect(dbname).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
		return nil
	}
	return insertResult.InsertedID
}

func InsertPelanggan(nama string, phoneNumber string, alamat string, email []string) interface{} {
	var pelanggan Pelanggan
	pelanggan.ID = primitive.NewObjectID()
	pelanggan.Nama = nama
	pelanggan.Phone_number = phoneNumber
	pelanggan.Alamat = alamat
	pelanggan.Email = email
	return InsertOneDoc("kantin", "kantin_pelanggan", pelanggan)
}

func GetPelangganByID(pelangganID primitive.ObjectID) (pelanggan Pelanggan) {
	collection := MongoConnect("kantin").Collection("kantin_pelanggan")
	filter := bson.M{"_id": pelangganID}
	err := collection.FindOne(context.TODO(), filter).Decode(&pelanggan)
	if err != nil {
		fmt.Printf("GetPelangganByID: %v\n", err)
	}
	return pelanggan
}

func GetAllPelanggan() (pelanggans []Pelanggan) {
	collection := MongoConnect("kantin").Collection("kantin_pelanggan")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("GetAllPelanggan: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.Background()) {
		var pelanggan Pelanggan
		if err := cursor.Decode(&pelanggan); err != nil {
			fmt.Printf("GetAllPelanggan: %v\n", err)
			continue
		}
		pelanggans = append(pelanggans, pelanggan)
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("GetAllPelanggan: %v\n", err)
	}
	return pelanggans
}

func InsertProduk(namaProduk string, deskripsi string, harga int) interface{} {
	var produk Produk
	produk.ID = primitive.NewObjectID()
	produk.Nama_Produk = namaProduk
	produk.Deskripsi = deskripsi
	produk.Harga = harga
	return InsertOneDoc("kantin", "Menu_produk", produk)
}

func GetProdukByID(produkID primitive.ObjectID) (produk Produk) {
	collection := MongoConnect("kantin").Collection("Menu_produk")
	filter := bson.M{"_id": produkID}
	err := collection.FindOne(context.TODO(), filter).Decode(&produk)
	if err != nil {
		fmt.Printf("GetProdukByID: %v\n", err)
	}
	return produk
}

func GetAllProduk() (produks []Produk) {
	collection := MongoConnect("kantin").Collection("Menu_produk")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("GetAllProduk: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.Background()) {
		var produk Produk
		if err := cursor.Decode(&produk); err != nil {
			fmt.Printf("GetAllProduk: %v\n", err)
			continue
		}
		produks = append(produks, produk)
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("GetAllProduk: %v\n", err)
	}
	return produks
}

func InsertTransaksi(metodePembayaran string, tanggalWaktu string) interface{} {
	var transaksi Transaksi
	transaksi.ID = primitive.NewObjectID()
	transaksi.Metode_Pembayaran = metodePembayaran
	transaksi.Tanggal_Waktu = tanggalWaktu
	return InsertOneDoc("kantin", "kantin_transaksi", transaksi)
}

func GetTransaksiByID(transaksiID primitive.ObjectID) (transaksi Transaksi) {
	collection := MongoConnect("kantin").Collection("kantin_transaksi")
	filter := bson.M{"_id": transaksiID}
	err := collection.FindOne(context.TODO(), filter).Decode(&transaksi)
	if err != nil {
		fmt.Printf("GetTransaksiByID: %v\n", err)
	}
	return transaksi
}

func GetAllTransaksi() (transaksis []Transaksi) {
	collection := MongoConnect("kantin").Collection("kantin_transaksi")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("GetAllTransaksi: %v\n", err)
		return nil
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.Background()) {
		var transaksi Transaksi
		if err := cursor.Decode(&transaksi); err != nil {
			fmt.Printf("GetAllTransaksi: %v\n", err)
			continue
		}
		transaksis = append(transaksis, transaksi)
	}
	if err := cursor.Err(); err != nil {
		fmt.Printf("GetAllTransaksi: %v\n", err)
	}
	return transaksis
}
