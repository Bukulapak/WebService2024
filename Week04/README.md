## Latihan kumpul disini

# Backend Package

Pada bagian ini dijelaskan bagaimana membangun otorisasi pada backend, beberapa langkah yang dilakukan :

- Membuat database backend
- Membuat function golang
- Membuat package golang dan merilisnya

Pastikan github pages sudah jalan di repo masing-masing.

## Basis data

Dalam memilih basis data skala enterprise harus mengacu pada hasil survey [Gartner](https://www.gartner.com/reviews/market/cloud-database-management-systems).Buatlah masing-masing basis data dari studi kasus yang dipilih :
- [MongoDB](https://www.mongodb.com/)
- MySQL : https://www.freemysqlhosting.net/, https://www.db4free.net/, https://remotemysql.com/


- Jika belum punya akun silahkan buat akun menggunakan akun Github (sign in)
  ![image](https://user-images.githubusercontent.com/26703717/224049906-a344f8b7-94eb-474c-bc2e-dae0b970eaf8.png)
  ![image](https://user-images.githubusercontent.com/15622730/223630677-23c059d9-0236-4cc3-9d35-14e5724003ee.png)
  ![image](https://user-images.githubusercontent.com/15622730/223630763-a1fc6f61-61e4-4f2d-be1a-72212ce3420d.png)

  Pada saat deploy
  Pilih Free Plan, Provider pilih Google Cloud, Region Pilih Jakarta, Nama Cluster sesuaikan saja, kemudian Create
  ![image](https://user-images.githubusercontent.com/26703717/224051590-07cbbd19-fa37-4171-9cf5-1cba2bd19f86.png)
  Pada security quickstart silahkan pilih username dan password (Jangan lupa untuk save password di notepad kalian), kemudian create user
  ![image](https://user-images.githubusercontent.com/26703717/224052500-44ccccaa-56df-4d5d-9dd1-e55c140cf2cc.png)

  Untuk Tahap 2, klik Add Current My IP Address
  ![image](https://user-images.githubusercontent.com/26703717/224052768-d8e5fb7b-cde2-4b3e-bfff-5f013fed3e78.png)
  Jika sudah klik Finish and Close
  ![image](https://user-images.githubusercontent.com/26703717/224053067-fd1af33f-fb91-49f2-b167-2e453c86cd2f.png)
  ![image](https://user-images.githubusercontent.com/15622730/223631673-e509a334-363c-4d51-a287-4a8531d023f9.png)

  Jika belum punya MongoDB Compass, silahkan download terlebih dahulu
  ![image](https://user-images.githubusercontent.com/26703717/224069777-635bc2e4-fc18-475f-93fc-164c7f5e7abe.png)

  Buka Mongo DB Compas, Kemudian Add New Connection, copy paste connection string pada gambar di atas, ganti Passwordnya dengan yang tadi sudah disimpan dan hapus kurung sikunya
  ![image](https://user-images.githubusercontent.com/26703717/224053890-a3c043b4-2e71-43aa-a5ee-446ef7629a56.png)
  Hasilnya seperti ini
  ![image](https://user-images.githubusercontent.com/26703717/224055259-6cd2878c-b427-4ee5-8335-a88e39aea9eb.png)
  Struktur tabel/collection di MongoDB terlihat seperti JSON
  ![image](https://user-images.githubusercontent.com/26703717/224055791-0bb78b63-e56b-416b-bdef-0dd48685f346.png)
  Pada mongodb.com untuk dapat diakses pada semua IP Address, klik ALLOW ACCESS FROM ANYWHERE atau dengan menuliskan 0.0.0.0/0
  ![image](https://user-images.githubusercontent.com/11188109/223330948-06b8493d-a9b7-4f19-86ba-06c35d6aaa96.png)


## Pengembangan Backend

Pada sesi ini kita akan mencoba mengembangkan package golang. Langkah untuk membuat backend di golang :

- Instalasi golang (https://go.dev/dl/)
- Definisikan dahulu struct
- Buatlah package dan fungsi menggunakan struct tersebut
- Fungsi mengakses langsung database

### Struct di golang

Komunikasi di golang menggunakan json dipermudah dengan adanya struct type. Struct type ini akan mendefinisikan bagaimana bentuk json yang berkomunikasi dari frontend menuju backend. Sebelumnya inisiasi dulu package yang akan kita buat, masuk ke folder kerja kita dan lakukan inisialisasi package

```sh
go mod init github.com/{username github kalian}/{nama repo kalian}/Week04/Praktikum/{NPM masing-masing}
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/f79a7a80-116a-46c1-a3e1-0d3ac62b87ec)

Setelah melakukan go mod init maka akan terbentuk file go.mod yang ada di dalam folder Week04/Praktikum/{NPM(12345)}.
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/b1bf6459-349e-4fca-86f1-8469848b10e3)

Buat file type.go pada folder tersebut. Kemudian isilah dengan struct yang akan kita buat. Perlu digaris bawahi dan wajib ada atribut :

- bson : mendefinisikan nama field pada collection di database mongoDB
- json : mendefinisikan nama atribut pada pertukaran json melalui API
- omitempty : diperbolehkan tidak diisi / nullable
- [] menunjukan data array didalam array

```go
package namapackage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Karyawan struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jabatan      string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	Jam_kerja    []JamKerja         `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
	Hari_kerja   []string           `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty"`
}

type JamKerja struct {
	Durasi     int      `bson:"durasi,omitempty" json:"durasi,omitempty"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty"`
	Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty"`
	Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty"`
	Shift      int      `bson:"shift,omitempty" json:"shift,omitempty"`
	Piket_tim  string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Longitude    float64            `bson:"longitude,omitempty" json:"longitude,omitempty"`
	Latitude     float64            `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin      string             `bson:"checkin,omitempty" json:"checkin,omitempty"`
	Biodata      Karyawan           `bson:"biodata,omitempty" json:"biodata,omitempty"`
}

```
Harap namapackage di atas disesuiakan dengan nama package folder NPM (misal 12345).

Seharusnya code di atas akan terdapat error, untuk mengatasinya kita lakukan kompilasi dependensi dengan perintah

```sh
go mod tidy
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/0c4a2f15-7c52-446a-bfba-d4b79574594b)

Seharusnya file type.go sekarang sudah tidak ada error

### Fungsi di golang

Edit Environtments Variables pada Windows
Buka MongoDB Compass, kemudian pilih copy connection string
![image](https://user-images.githubusercontent.com/26703717/224064475-0f124e7b-a817-4058-85c8-ef34e3bff53e.png)
Paste pada environments variables, kemudian simpan
![image](https://user-images.githubusercontent.com/26703717/224064873-4cc79cf3-c675-42a8-8f36-1bb9daf331dc.png)
Jika sudah, close text editor (misal VSCode), kemudian buka lagi VSCodenya

Disini dipelajari bagaimana membuat fungsi dan menggunakan environment variabel di golang. Pertama simpan dahulu MONGOSTRING koneksi mongodb di environtment variabel windows atau sistem operasi.

Buat file praktikum.go pada folder NPM kalian yang berisi

```go
package namapackage

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

func InsertPresensi(long float64,lat float64, lokasi string, phonenumber string, checkin string, biodata Karyawan) (InsertedID interface{}) {
	var presensi Presensi
	presensi.Latitude = long
	presensi.Longitude = lat
	presensi.Location = lokasi
	presensi.Phone_number = phonenumber
	presensi.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	presensi.Checkin = checkin
	presensi.Biodata = biodata
	return InsertOneDoc("tesdb2024", "presensi", presensi)
}

func GetKaryawanFromPhoneNumber(phone_number string) (staf Presensi) {
	karyawan := MongoConnect("tesdb2024").Collection("presensi")
	filter := bson.M{"phone_number": phone_number}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getKaryawanFromPhoneNumber: %v\n", err)
	}
	return staf
}

func GetAllPresensi() (data []Presensi) {
	karyawan := MongoConnect("tesdb2024").Collection("presensi")
	filter := bson.M{}
	cursor, err := karyawan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

```

rapihkan dependensi

```sh
go mod tidy
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/41de2fc7-1dc0-48ae-8b90-d9574e9606c8)

### Testing Fungsi

Buat file dengan format praktikum_test.go
```go
package namapackage

import (
	"fmt"
	"testing"
)

func TestInsertPresensi(t *testing.T) {
	long := 98.345345
	lat := 123.561651
	lokasi := "Rumah"
	phonenumber := "68122221814"
	checkin := "masuk"
	biodata := Karyawan{
		Nama : "Drake",
		Phone_number : "628456456211",
		Jabatan : "Rakyat Biasa",
	}
	hasil:=InsertPresensi(long ,lat , lokasi , phonenumber , checkin , biodata )
	fmt.Println(hasil)
}

func TestGetKaryawanFromPhoneNumber(t *testing.T) {
	phonenumber := "68122221814"
	biodata:=GetKaryawanFromPhoneNumber(phonenumber)
	fmt.Println(biodata)
}

func TestGetAll(t *testing.T) {
	data := GetAllPresensi()
	fmt.Println(data)
}


```

Jalankan testing

```sh
go test
```
Berhasil create dan view
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/265a7151-0e10-4e27-9984-99169abd0efd)

Untuk melihat hasil test pada MongoDB Compass
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/3ada91f3-0820-412b-94c9-24eb184a0ea1)

Untuk menjalankan run test per fungsi bisa menggunakan perintah

```sh
go test -run ^NAMAFUNCTION$
```

Contoh penggunaannya kita panggil fungsi getKaryawanFromPhoneNumber untuk melihat hasilnya pada terminal
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/b216c486-29b1-45d0-a9e2-024fb4589dc7)

### Publish Package

Setting dahulu environtment variabel di windows :
GOPROXY=proxy.golang.org
![image](https://user-images.githubusercontent.com/15622730/223696410-11ffa42a-225e-4d50-a30e-97d1894407d9.png)

**Untuk melakukan listing package pastikan kalian sudah melakukan push ke repo** kalian masing-masing, setelah itu lakukanlah listing package menggunakan perintah berikut
```sh
go list -m github.com/{username github}/{namarepo}@{versi tag terbaru}
```
penggunaan listing package di atas adalah sebagai berikut

```sh
git tag
git tag v0.0.1
git push origin --tags
go list -m github.com/Bukulapak/WebService2024@v0.0.1
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/566a67a3-67f6-4996-972a-aac8f01b5a2d)

**Catatan** : untuk membuat tag pastikan versinya dibuat secara urut (sequential), jika sudah ada tag v0.0.1 maka tag selanjutnya adalah v0.0.2, dst.

Jika sudah seharusnya package kalian ada di [pkg.go.dev](https://pkg.go.dev/). Butuh waktu mungkin beberapa hari untuk rilis