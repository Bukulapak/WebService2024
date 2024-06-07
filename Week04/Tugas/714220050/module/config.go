package _714220050

import (
	"github.com/aiteung/atdb"
	"os"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "kemahasiswaan",
}

var MongoConn = atdb.MongoConnect(MongoInfo)