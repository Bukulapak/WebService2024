# Dokumentasi Golang Fiber menggunakan Swagger

## Intro
Swagger adalah sebuah alat yang digunakan untuk mendokumentasikan dan menguji API. Dengan menggunakan Swagger, kita dapat membuat dokumentasi API yang lengkap dan interaktif secara otomatis berdasarkan kode sumber aplikasi.

Golang Fiber adalah sebuah framework web yang ringan dan cepat untuk bahasa pemrograman Go. Fiber menyediakan fitur-fitur yang kuat untuk membangun aplikasi web, termasuk dukungan untuk routing, middleware, dan pengelolaan respons HTTP yang efisien.

Menggabungkan Swagger dengan Golang Fiber memungkinkan kita untuk secara otomatis menghasilkan dokumentasi API yang komprehensif dan mudah diakses. Dengan menggunakan anotasi atau tag khusus dalam kode sumber aplikasi, kita dapat menentukan detail rute API, metode yang didukung, parameter yang diterima, dan respons yang dihasilkan.

Beberapa langkah yang umumnya dilakukan untuk menggunakan Swagger dengan Golang Fiber adalah sebagai berikut:

1. Instalasi dependensi: Pertama, kita perlu menginstal paket-paket terkait Swagger dan Golang Fiber. Dalam kasus ini, kita dapat menggunakan package `swaggo` untuk Swagger menggunakan Golang Fiber.
2. Anotasi kode sumber: Pada titik-titik yang sesuai dalam kode sumber aplikasi Golang Fiber, kita dapat menambahkan anotasi atau tag Swagger untuk mendefinisikan rute API, parameter, respons, dan informasi lainnya. Ini biasanya dilakukan menggunakan komentar pada fungsi handler atau struktur data yang terlibat.
3. Menghasilkan dokumentasi: Setelah anotasi telah ditambahkan pada kode sumber (Golang), kita dapat menggunakan tools `go-swagger` untuk menghasilkan file Swagger JSON atau YAML yang berisi dokumentasi lengkap tentang API. Tools ini nantinya akan membaca anotasi dan menghasilkan file Swagger berdasarkan informasi tersebut.
4. Mengintegrasikan dengan Golang Fiber: Selanjutnya, kita perlu mengintegrasikan file Swagger yang dihasilkan dengan aplikasi Golang Fiber. Dalam hal ini, kita dapat menyediakan endpoint khusus pada aplikasi yang akan digunakan untuk menyajikan file Swagger kepada pengguna.
5. Uji API dan dokumentasi: Setelah aplikasi berjalan, kita dapat mengakses dokumentasi API yang dihasilkan melalui endpoint yang telah ditentukan. Dari sana, kita dapat menjelajahi rute API, mencoba permintaan, dan melihat respons yang diharapkan.

Dengan menggunakan Swagger dan Golang Fiber, kita dapat meningkatkan efisiensi dalam mengembangkan dan memelihara API. Dokumentasi yang dihasilkan secara otomatis memudahkan pengguna atau pengembang lain untuk memahami dan menggunakan API yang kita buat. Selain itu, Swagger juga memungkinkan kita untuk menguji API dengan mudah dan memberikan kesempatan untuk berinteraksi dengan endpoint-endpoint yang ada secara langsung dari dalam dokumentasi.

## Membuat Dokumentasi
1. Yang pertama dilakukan adalah silahkan install terlebih dahulu swagger pada project Boilerplate masing-masing. Buka project Boilerplate, kemudian klik [disini](https://github.com/gofiber/swagger) untuk instalasi pada terminal. Dan lakukan sampai tahap 4 (namun pada tahap `Canonical example` tidak perlu membuat file main.go , cukup menambahkan anotasi @title hingga @BasePath pada bagian atas function main)

    **Jangan lupa juga lakukan import seperti gambar di bawah pada line 14 sesuai direktori kalian masing-masing**
    ```sh
    // @title TES SWAGGER ULBI
    // @version 1.0
    // @description This is a sample swagger for Fiber

    // @contact.name API Support
    // @contact.url https://github.com/indrariksa
    // @contact.email indra@ulbi.ac.id

    // @host ws-indra2024-878f7e6fab92.herokuapp.com
    // @BasePath /
    // @schemes https http
    ```
<p align="center">
  <img src="https://github.com/Bukulapak/WebService2024/assets/26703717/a37343d8-b6e5-491f-a1c3-566ba73add74">
</p>

  > Anotasi Swagger di atas adalah metadata yang digunakan untuk mengonfigurasi dan mendokumentasikan API menggunakan Swagger pada framework Golang Fiber. Berikut adalah penjelasan dari setiap anotasi tersebut:
  > * @title: Menentukan judul atau nama dari API yang akan didokumentasikan. Pada contoh di atas, judul API adalah "TES SWAG".
  > * @version: Menentukan versi dari API yang akan didokumentasikan. Pada contoh di atas, versi API adalah "1.0".
  > * @description: Memberikan deskripsi atau penjelasan singkat tentang API yang akan didokumentasikan.
  > * @contact.name: Menyediakan informasi tentang nama kontak untuk dukungan API.
  > * @contact.url: Menyediakan URL yang mengarah ke sumber daya yang memberikan dukungan atau informasi tambahan tentang API. Pada contoh di atas, URL GitHub (https://github.com/indrariksa) digunakan.
  > * @contact.email: Menyediakan alamat email untuk kontak dukungan API.
  > * @host: Menentukan host atau server tempat API dijalankan. Pada contoh di atas, API dijalankan di host "ws-indra2024-878f7e6fab92.herokuapp.com".
  > * @BasePath: Menentukan basis path atau awalan URL dari semua rute API. Pada contoh di atas, basis path adalah "/" yang berarti tidak ada awalan tambahan.
  > * @schemes: Menentukan skema protokol yang didukung oleh API. Pada contoh di atas, API mendukung skema HTTPS dan HTTP.
  > Dengan menggunakan anotasi Swagger ini, kalian dapat mengkonfigurasi dan mendokumentasikan API Anda secara lebih terstruktur dan informatif menggunakan Swagger.

2. Selanjutnya adalah membuat routing swagger pada file `url.go`
```go
page.Get("/docs/*", swagger.HandlerDefault)
```
Di atas kita akan menampilkan swagger pada endpoint `docs` menggunakan method GET

Jangan lupa tambahkan import
```sh
import "github.com/gofiber/swagger" // swagger handler
```

3. Jika sudah, kita harus melakukan perintah swag init kembali untuk menggenerate `Swagger Specification`
```sh
swag init
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/52e6e09e-277b-4054-8e19-7f0a437dcc86)

Jika berhasil, kalian akan melihat 3 file yang terupdate di dalam folder docs. File-file ini adalah:
* docs.go => Diperlukan untuk menghasilkan SwaggerUI.
* swagger.json => Spesifikasi Swagger dalam format file json.
* swagger.yaml => Spesifikasi Swagger dalam format file yaml.

4. Sekarang lakukan push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint `/docs`
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/ca0ae85c-debc-4046-b27d-de0be035f2b9)

5. Dan kita berhasil menampilkan swagger pada project kita. Namun disini kita belum membuat Dokumentasi lengkapnya (seperti method Get,Post,Put,Delete)

### Dokumentasi GET (Get All dan Get By ID)
1. Buka Project Package Backend, kemudian kita cari struct pada folder `model/type.go`, kemudian kita copy semua code di atas dan paste pada Project Boilerplate di folder `controller` dan buat file bernama `struct.go` 

    **(nama package samakan dengan file coba.go yang ada pada file controller)**
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/96476edf-9a53-42c0-bb22-fd1fda54c8c2)

    ```go
    package controller

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

2. Pada project Boilerplate tambahkan anotasi berikut pada file `coba.go` di atas function `GetAllPresensi`
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/7a1d6b3c-203a-4806-a7e4-708855e3d71b)

    ```sh
    // GetPresensi godoc
    // @Summary Get All Data Presensi.
    // @Description Mengambil semua data presensi.
    // @Tags Presensi
    // @Accept json
    // @Produce json
    // @Success 200 {object} Presensi
    // @Router /presensi [get]
    ```
    > Anotasi di atas digunakan untuk mendokumentasikan endpoint `GetPresensi` yang mengambil semua data presensi.
    > - `// @Summary Get All Data Presensi.`: Memberikan ringkasan singkat tentang apa yang dilakukan oleh endpoint ini, yaitu mengambil semua data presensi.
    > - `// @Description Mengambil semua data presensi.`: Menyediakan deskripsi lebih rinci tentang fungsi endpoint ini.
    > - `// @Tags Presensi`: Menyertakan tag "Presensi" untuk mengelompokkan endpoint ini bersama dengan endpoint-endpoint terkait yang memiliki tag yang sama.
    > - `// @Accept json`: Menginformasikan bahwa endpoint ini dapat menerima permintaan dengan tipe konten json.
    > - `// @Produce json`: Menginformasikan bahwa endpoint ini akan menghasilkan respons dengan tipe konten json.
    > - `// @Success 200 {object} Presensi`: Menyatakan bahwa jika operasi berhasil (respons kode status 200), respons yang dikembalikan akan berbentuk objek dengan struktur yang ditentukan oleh tipe `Presensi`.
    > - `// @Router /presensi [get]`: Menentukan rute URL untuk endpoint ini, yaitu `/presensi` dengan metode HTTP GET.

3. Lakukan perintah `swag init` pada terminal. Seharusnya terdapat error seperti di bawah
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/3eed1cb9-e652-4847-bf21-bcbdcb27a397)

    agar swag mengenali tipe tersebut, kalian bisa memberikan anotasi tambahan yang memberi tahu Swagger untuk menganggapnya sebagai tipe tertentu, seperti `string`. **(note : struct pada Boilerplate hanya digunakan sebagai contoh response pada Swagger)**
    ```sh
    swaggertype:"string"
    ```
    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/a08b67b2-e513-49ba-8236-e6138dba160f)

    Dalam contoh di atas, anotasi `swaggertype:"string"` memberi tahu Swagger untuk menganggap primitive.DateTime sebagai string.

4. Lakukan perintah `swag init` kembali, seperti gambar di bawah tidak terdapat error
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/81088035-5918-443b-8a5b-7ecf2d006bff)

5. Push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint `/docs`
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/c58f29ff-3def-4606-94fb-1b90b2e49618)

    Bisa dilihat contoh response (example value) di bawah didapatkan dari `struct` yang kita buat di project Boilerplate
    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/68442ea9-07e1-4884-8080-7effe9280889)

    Untuk ujicoba menggunakan SwaggerUI bisa klik pada Tombol `Try it out` kemudian pilih `Execute`, bisa dilihat data berhasil ditampilkan dengan response code 200
    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/a2ed296b-7c50-4971-b67d-bb6e21a60b71)

6. Sampai disini kita berhasil membuat contoh dokumentasi dengan method GET (Get All)
7. Selanjutnya kita akan membuat untuk mendapatkan data berdasarkan parameter ID
8. Buka file coba.go kemudian tambah anotasi berikut di atas function GetPresensiID
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/d1e73136-af3d-4bbc-a09f-8592f1ceea22)

    ```sh
    // GetPresensiID godoc
    // @Summary Get By ID Data Presensi.
    // @Description Ambil per ID data presensi.
    // @Tags Presensi
    // @Accept json
    // @Produce json
    // @Param id path string true "Masukan ID"
    // @Success 200 {object} Presensi
    // @Failure 400
    // @Failure 404
    // @Failure 500
    // @Router /presensi/{id} [get]
    ```
    > Anotasi Swagger pada kode di atas adalah sebagai berikut:
    > - @Summary, @Description, @Tags, @Accept, @Produce sudah di jelaskan sebelumnya
    > - @Param: Mendefinisikan parameter yang dibutuhkan oleh operasi, yaitu id yang merupakan bagian dari path dan memiliki tipe data string. Parameter ini wajib diisi (required) dan akan digunakan untuk mengambil data presensi berdasarkan ID.
    > - @Success: Menyatakan bahwa jika operasi berhasil (respons kode status 200), respons yang dikembalikan akan berbentuk objek dengan struktur yang ditentukan oleh tipe `Presensi`.
    > - @Failure: Mendefinisikan respons gagal dari operasi ini dengan kode status yang berbeda-beda. Pada contoh di atas, terdapat beberapa kode status yang ditangani, yaitu 400 (Bad Request), 404 (Not Found), dan 500 (Internal Server Error).
    > - @Router: Menentukan rute URL untuk endpoint ini, yaitu `/presensi/{id}` dengan metode HTTP GET.

9. Lakukan perintah `swag init` pada terminal, push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint `/docs` (pada gambar di bawah terlihat sudah ada untuk Get By ID)
    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/0a05609b-7a78-492a-bb0f-04653fa2f377)

    Untuk ujicoba menggunakan SwaggerUI bisa klik pada Tombol `Try it out` masukkan ID kemudian pilih `Execute`, bisa dilihat data berhasil ditampilkan dengan response code 200
    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/c5536575-fbba-465e-8577-874a4132ead8)

    Jika kondisi gagal seperti berikut yang menghasilkan response code 400 dengan memasukkan ID misal `123`
    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/039bf692-b26b-447b-a70a-372b478d807e)

### Dokumentasi POST
1. Buka file `coba.go` kemudian tambah anotasi berikut di atas function `InsertDataPresensi`
    ```sh
    // InsertDataPresensi godoc
    // @Summary Insert data presensi.
    // @Description Input data presensi.
    // @Tags Presensi
    // @Accept json
    // @Produce json
    // @Param request body Presensi true "Payload Body [RAW]"
    // @Success 200 {object} Presensi
    // @Failure 400
    // @Failure 500
    // @Router /insert [post]
    ```

    > Anotasi Swagger pada kode di atas adalah sebagai berikut:
    > - `@Summary, @Description, @Tags, @Accept, @Produce sudah di jelaskan sebelumnya
    > - `@Param`: Mendefinisikan parameter yang dibutuhkan oleh operasi, yaitu `request` yang merupakan body payload untuk data presensi. Parameter ini memiliki tipe data `Presensi` dan wajib diisi (required).
    > - `@Success`: Mendefinisikan respons sukses, yaitu kode status 200 (OK) dan objek yang dihasilkan adalah tipe data `Presensi`.
    > - `@Failure`: Mendefinisikan respons gagal dengan kode status yang berbeda-beda. Pada contoh di atas, terdapat beberapa kode status yang ditangani, yaitu 400 (Bad Request) dan 500 (Internal Server Error).
    > - `@Router`: Mendefinisikan endpoint URL, yaitu "/ins" dengan metode POST.

2. Lakukan perintah `swag init` pada terminal, push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint `/docs` (pada gambar di bawah terlihat sudah ada untuk method POST)
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/a64543a7-73fc-4a34-a26c-bd8dca674a0d)

3. Namun bisa dilihat ketika kita memilih `Try it out`, contoh data RAW nya hanya menampilkan tipe data seperti int(0) dan string.

    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/3b112efb-fab2-4dba-ba1a-d30f3ec66ad9)
    
    Untuk memudahkannya kita dapat mengubah sedikit code pada `struct.go`(project boilerplate). Dengan menambah `example:"Contoh"`
    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/cf3c79b4-ae16-4ad9-8059-8257c9eec484)

    Atau bisa copy paste di bawah
    ```go
    type Karyawan struct {
        ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
        Nama         string             `bson:"nama,omitempty" json:"nama,omitempty" example:"Tes Swagger"`
        Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08123456789"`
        Jabatan      string             `bson:"jabatan,omitempty" json:"jabatan,omitempty" example:"Anonymous"`
        Jam_kerja    []JamKerja         `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
        Hari_kerja   []string           `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty" example:"Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu"`
    }

    type JamKerja struct {
        Durasi     int      `bson:"durasi,omitempty" json:"durasi,omitempty" example:"8"`
        Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty" example:"08:00"`
        Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty" example:"16:00"`
        Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty" example:"7"`
        Hari       []string `bson:"hari,omitempty" json:"hari,omitempty" example:"Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu"`
        Shift      int      `bson:"shift,omitempty" json:"shift,omitempty" example:"2"`
        Piket_tim  string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty" example:"Piket Z"`
    }

    type Presensi struct {
        ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
        Longitude    float64            `bson:"longitude,omitempty" json:"longitude,omitempty" example:"123.11"`
        Latitude     float64            `bson:"latitude,omitempty" json:"latitude,omitempty" example:"123.12"`
        Location     string             `bson:"location,omitempty" json:"location,omitempty" example:"Bandung"`
        Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08123456789"`
        Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty" swaggertype:"string" example:"2024-09-01T00:00:00Z" format:"date-time"`
        Checkin      string             `bson:"checkin,omitempty" json:"checkin,omitempty" example:"Masuk"`
        Biodata      Karyawan           `bson:"biodata,omitempty" json:"biodata,omitempty"`
    }
    ```

4. Lakukan perintah `swag init` pada terminal, push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint `/doc`. Pada method POST dan GET contoh response code pun akan berubah mengikuti struct yang ada di Boilerplate.

5. Klik pada Tombol `Try it out` maka contoh RAW data pun akan mengikuti struct `Presensi` isinya sama dengan example yang sudah kita buat di atas.
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/2fe99484-5f96-4a03-a0fe-24ee894a3adb)
6. Kemudian kalian coba execute dengan RAW data yang ada, seharusnya kalian mendapat response code 500 seperti di bawah
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/c02c40d5-9356-4590-8a0e-c409e0b1accd)
7. Error di atas berarti bahwa proses unmarshalling data JSON ke dalam ObjectID MongoDB gagal karena panjang string yang diberikan tidak sesuai dengan panjang yang diharapkan. Perlu diingat bahwa field ID sebenarnya tidak perlu diinputkan secara manual karena akan terisi otomatis oleh MongoDB, begitu juga dengan DateTime. Jadi, cara sederhana untuk mengatasi masalah ini adalah dengan menghapus field ID dari input. Namun, jika kita menghapus field ini dari struct, maka semua field ID di response semua endpoint swagger tidak akan tampil. Oleh karena itu, daripada menghapus field ID pada struct, lebih baik kita membuat struct baru khusus untuk request pada endpoint POST dan PUT di swagger.

    Untuk field ID hanya terdapat pada struct Karyawan dan Presensi, jadi buatlah struct baru pada project Boilerplate folder `controller/struct.go` untuk melakukan request seperti berikut
    ```go
    type ReqKaryawan struct {
        Nama         string     `bson:"nama,omitempty" json:"nama,omitempty" example:"Tes Swagger"`
        Phone_number string     `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08123456789"`
        Jabatan      string     `bson:"jabatan,omitempty" json:"jabatan,omitempty" example:"Anonymous"`
        Jam_kerja    []JamKerja `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
        Hari_kerja   []string   `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty" example:"Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu"`
    }

    type ReqPresensi struct {
        Longitude    float64     `bson:"longitude,omitempty" json:"longitude,omitempty" example:"123.11"`
        Latitude     float64     `bson:"latitude,omitempty" json:"latitude,omitempty" example:"123.12"`
        Location     string      `bson:"location,omitempty" json:"location,omitempty" example:"Bandung"`
        Phone_number string      `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08123456789"`
        Checkin      string      `bson:"checkin,omitempty" json:"checkin,omitempty" example:"Masuk"`
        Biodata      ReqKaryawan `bson:"biodata,omitempty" json:"biodata,omitempty"`
    }
    ```

    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/9ad6a8ca-8475-490c-a8cf-a18c08bb81c9)

8. Jangan lupa juga ubah anotasi @Param pada function `InsertDataPresensi` pada file `coba.go` diubah menjadi struct `ReqPresensi`
    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/75f48bdf-1bb6-4398-83dc-593a071a2b05)
9. Lakukan perintah `swag init` pada terminal, push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint `/doc`. Coba lakukan POST dengan klik pada Tombol `Try it out` maka contoh RAW data pun akan mengikuti struct `ReqPresensi`. Bisa dilihat field ID sudah tidak ada.
    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/d8bd77aa-7a34-46a5-ae8a-79160b30e1bc)
10. Kemudian kalian coba execute dengan RAW data yang ada, seharusnya kalian mendapat response code 200 seperti di bawah

    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/e669aa67-7e71-41d4-a3de-05bacbe44b8f)

### Dokumentasi PUT
1. Buka file `coba.go` kemudian tambah anotasi berikut di atas function `UpdateData`
```sh
// UpdateData godoc
// @Summary Update data presensi.
// @Description Ubah data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body ReqPresensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /update/{id} [put]
```
> Anotasi Swagger pada kode di atas hampir sama seperti POST hanya berbeda pada:
> - `@Param`: Mendefinisikan parameter yang dibutuhkan oleh operasi. Terdapat dua parameter dalam contoh ini. Pertama adalah `id` yang diambil dari path URL, dan harus diisi (required). Kedua adalah `request` yang merupakan body payload untuk data presensi dan juga wajib diisi (required).
> - `@Router`: Mendefinisikan endpoint URL , yaitu "/upd/{id}" dengan metode PUT.

2. Lakukan perintah `swag init` pada terminal, push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint `/docs` (pada gambar di bawah terlihat sudah ada untuk method PUT)

    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/0e6db202-c8e7-4a72-9666-3428cd94d202)

3. Kemudian kalian ujicoba perintah PUT ini, jika berhasil akan menampilkan response code 200, di bawah saya coba ubah pada field `nama`
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/c240acf3-5e1b-45cb-aeb1-e3df7707fd68)

4. Kemudian jika tidak ada perubahan apapun ketika melakukan update data maka akan menampilkan response code 500
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/67c2dd61-26ed-4e15-bf16-8e4a46c8bf44)

### Dokumentasi DELETE
1. Buka file `coba.go` kemudian tambah anotasi berikut di atas function `DeletePresensiByID`
```sh
// DeletePresensiByID godoc
// @Summary Delete data presensi.
// @Description Hapus data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
```
> Anotasi Swagger pada kode di atas adalah sebagai berikut:
> - `@Param`: Mendefinisikan parameter yang dibutuhkan oleh operasi. Terdapat satu parameter dalam contoh ini, yaitu `id` yang diambil dari path URL, dan harus diisi (required).
> - `@Success`: Mendefinisikan respons sukses, yaitu kode status 200 (OK).
> - `@Failure`: Mendefinisikan respons gagal dengan kode status yang berbeda-beda. Pada contoh di atas, terdapat beberapa kode status yang ditangani, yaitu 400 (Bad Request) dan 500 (Internal Server Error).
> - `@Router`: Mendefinisikan endpoint URL, yaitu "/delete/{id}" dengan metode DELETE.

2. Lakukan perintah `swag init` pada terminal, push ke repo github dan heroku, kemudian buka base URL heroku kalian dengan menambahan endpoint `/docs (pada gambar di bawah terlihat sudah ada untuk method DELETE)
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/0217c528-36be-41ec-a275-2bc4c8cc899e)

3. Kemudian kalian ujicoba perintah DELETE ini, jika berhasil akan menampilkan response code 200
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/0a990878-59c5-4038-92e4-c77e3b5e6348)

4. Kemudian jika id tidak terdapat pada database maka akan menampilkan response code 500
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/cd13084e-dd95-479f-8460-5af199312183)

5. Kemudian jika format id tidak sesuai maka akan menampilkan response code 400
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/489e6991-79db-4614-adc3-b20c62e4fd30)


---
*UNTUK MELIHAT DOKUMENTASI LENGKAP GOLANG SWAGGER https://github.com/swaggo/swag#getting-started*
