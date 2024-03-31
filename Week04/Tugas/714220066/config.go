package ws

import (
	"github.com/kamagasaki/go-utils/mongo"
)

var TugasMDB = "DBTugas"

// var MongoString = os.Getenv("MONGOSTRCONNECT")
var MongoString = " "

var TugasDB = mongo.DBInfo{
	DBString: MongoString,
	DBName:   TugasMDB,
}

var PertemuanColl = "pertemuan"
var PresensiColl = "presensi"
var MhsColl = "mahasiswa"
