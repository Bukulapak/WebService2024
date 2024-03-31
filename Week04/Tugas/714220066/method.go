package ws

import (
	"errors"
	"github.com/kamagasaki/go-utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertPertemuan(requestData Pertemuan) error {
	db := mongo.MongoConnect(TugasDB)
	insertedID := mongo.InsertOneDoc(db, PertemuanColl, requestData)
	if insertedID == nil {
		return errors.New("Yikes, can't insert data!")
	}
	return nil
}

func GetPertemuan(filter bson.M) ([]Pertemuan, error) {
	db := mongo.MongoConnect(TugasDB)
	var pertemuan []Pertemuan
	err := mongo.GetAllDocByFilter[Pertemuan](db, PertemuanColl, filter, &pertemuan)
	if err != nil {
		return nil, err
	}
	return pertemuan, nil
}

func InsertMhs(requestData Mahasiswa) error {
	db := mongo.MongoConnect(TugasDB)
	insertedID := mongo.InsertOneDoc(db, MhsColl, requestData)
	if insertedID == nil {
		return errors.New("Yikes, can't insert data!")
	}
	return nil
}

func GetMhs(filter bson.M) ([]Mahasiswa, error) {
	db := mongo.MongoConnect(TugasDB)
	var mahasiswa []Mahasiswa
	err := mongo.GetAllDocByFilter[Mahasiswa](db, MhsColl, filter, &mahasiswa)
	if err != nil {
		return nil, err
	}
	return mahasiswa, nil
}

func InsertPresensi(requestData Presensi) error {
	db := mongo.MongoConnect(TugasDB)
	insertedID := mongo.InsertOneDoc(db, PresensiColl, requestData)
	if insertedID == nil {
		return errors.New("Yikes, can't insert data!")
	}
	return nil
}

func GetPresensi(filter bson.M) ([]Presensi, error) {
	db := mongo.MongoConnect(TugasDB)
	var presensi []Presensi
	err := mongo.GetAllDocByFilter[Presensi](db, PresensiColl, filter, &presensi)
	if err != nil {
		return nil, err
	}
	return presensi, nil
}
