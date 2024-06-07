# UPDATE DAN DELETE DATA

# Backend Package

Buka project package backend presensi pada VScode atau GoLand

### Update Data
Pada folder `module/praktikum.go` tambahkan code berikut
```go
func UpdatePresensi(db *mongo.Database, col string, id primitive.ObjectID, long float64, lat float64, lokasi string, phonenumber string, checkin string, biodata model.Karyawan) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"longitude":    long,
			"latitude":     lat,
			"location":     lokasi,
			"phone_number": phonenumber,
			"checkin":      checkin,
			"biodata":      biodata,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdatePresensi: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}
```
> Fungsi UpdatePresensi digunakan untuk memperbarui data presensi dalam database MongoDB. Berikut adalah penjelasan masing-masing bagian dari fungsi tersebut:
> 1. Parameter:
> * db *mongo.Database: Objek database MongoDB yang digunakan untuk mengakses koleksi data.
> * col string: Nama koleksi data presensi dalam database.
> * id primitive.ObjectID: ID objek presensi yang akan diperbarui.
> * long float64: Nilai longitude baru yang akan diupdate.
> * lat float64: Nilai latitude baru yang akan diupdate.
> * lokasi string: Lokasi baru yang akan diupdate.
> * phonenumber string: Nomor telepon baru yang akan diupdate.
> * checkin string: Nilai check-in baru yang akan diupdate.
> * biodata model.Karyawan: Objek biodata karyawan baru yang akan diupdate.
> 2. Membuat filter:
> * Filter dibuat dengan menggunakan ID objek presensi yang akan diperbarui.
> 3. Membuat update:
> * Update dilakukan menggunakan operator $set pada MongoDB untuk mengatur nilai baru dari field-field yang akan diperbarui.
> * Field-field yang diperbarui adalah longitude, latitude, location, phone_number, checkin, dan biodata.
> 4. Melakukan pembaruan data:
> * Fungsi UpdateOne digunakan untuk memperbarui data dalam database.
> * Hasil pembaruan data disimpan dalam variabel result.
> 5. Menangani hasil pembaruan:
> * Jika terjadi error saat memperbarui data, pesan error akan dicetak dan fungsi akan mengembalikan error.
> * Jika tidak ada data yang berubah (modified count = 0), fungsi akan mengembalikan error dengan pesan "No data has been changed with the specified ID".
> * Jika pembaruan data berhasil, fungsi akan mengembalikan nil (tanpa error).
> Fungsi UpdatePresensi ini dapat digunakan untuk memperbarui data presensi karyawan dengan memberikan nilai baru untuk field-field yang ingin diperbarui dalam database MongoDB.

* Note : untuk update data kita tidak perlu membuat file testing, nanti kita coba saja via postman

Rapihkan dependensi
```go
go mod tidy
```

### Delete Data
Pada folder `module/praktikum.go` tambahkan code berikut
```go
func DeletePresensiByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	karyawan := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := karyawan.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
```
> Fungsi `DeletePresensiByID` digunakan untuk menghapus data presensi berdasarkan ID. Berikut adalah penjelasannya:
> - Fungsi ini menerima tiga parameter: `_id` yang merupakan ID dari data yang akan dihapus, `db` yang merupakan objek koneksi database tipe `*mongo.Database`, dan `col` yang merupakan nama koleksi.
> - Variabel `karyawan` merepresentasikan koleksi tempat data akan dihapus.
> - Variabel `filter` dibuat menggunakan parameter `_id` untuk menentukan dokumen yang akan dihapus. Digunakan tipe `bson.M` untuk membuat filter berdasarkan ID.
> - Selanjutnya, fungsi `DeleteOne` digunakan untuk menghapus data yang sesuai dengan filter yang diberikan. Hasil penghapusan disimpan dalam variabel `result` dan error disimpan dalam variabel `err`.
> - Jika terjadi error dalam proses penghapusan, fungsi akan mengembalikan error yang menjelaskan pesan error. Jika penghapusan berhasil dan tidak ada data yang dihapus, maka fungsi akan mengembalikan error yang menyatakan bahwa data dengan ID tersebut tidak ditemukan.
> - Jika tidak ada error dalam proses penghapusan, fungsi akan mengembalikan nilai `nil` untuk menunjukkan bahwa penghapusan berhasil dilakukan.

Kemudian pada file praktikum_test.go isikan code berikut untuk melakukan testing hapus data, untuk id yang ada di bawah sesuaikan saja dengan id yang ada di mongoDB kalian.

```go
func TestDeletePresensiByID(t *testing.T) {
	id := "6412ce78686d9e9ba557cf8a" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeletePresensiByID(objectID, module.MongoConn, "presensi")
	if err != nil {
		t.Fatalf("error calling DeletePresensiByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = module.GetPresensiFromID(objectID, module.MongoConn, "presensi")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}
```

Rapihkan dependensi
```go
go mod tidy
```

Jalankan testing
```go
go test -run ^TestDeletePresensiByID$
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/1c36ec9f-a641-4807-94fc-b48ee698340e)

Cek pada MongoDB apakah data sudah terhapus.

Kemudian lakukan git tag agar package terupdate, tutor [disini](https://github.com/Bukulapak/WebService2024/tree/main/Week04#publish-package)

# Boilerplate

### Update Data
Buka project boilerplate pada VScode atau GoLand, pada terminal ketikkan perintah go get untuk update package backend (project latihan di atas yang sudah upgrade versi)
```go
go get {url package backend di pkg.go.dev}@versi terbaru

contoh : go get github.com/indrariksa/cobapakcage@v0.0.8
```

Kemudian pada file `coba.go` di folder controller kita tambahkan fungsi `UpdateData`
```go
func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var presensi inimodel.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = cek.UpdatePresensi(db, "presensi",
		objectID,
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}
```
Fungsi `UpdateData` digunakan sebagai handler untuk mengupdate data presensi dalam webservice menggunakan framework Fiber. Berikut adalah penjelasan masing-masing bagian dari fungsi tersebut:

1. Menghubungkan ke database:
   - Mengambil koneksi database dari konfigurasi `config.Ulbimongoconn` dan disimpan dalam variabel `db`.

2. Mendapatkan ID dari parameter URL:
   - Mengambil nilai ID dari parameter URL menggunakan `c.Params("id")` dan disimpan dalam variabel `id`.

3. Parsing ID menjadi ObjectID:
   - Mengubah string ID menjadi tipe `primitive.ObjectID` menggunakan `primitive.ObjectIDFromHex(id)`.
   - Jika terjadi kesalahan dalam parsing, akan mengembalikan respons HTTP dengan status `http.StatusInternalServerError` dan pesan error.

4. Parsing request body:
   - Mengurai (parse) body request menjadi objek `Presensi` menggunakan `c.BodyParser(&presensi)`.
   - Jika terjadi kesalahan dalam parsing, akan mengembalikan respons HTTP dengan status `http.StatusInternalServerError` dan pesan error.

5. Memanggil fungsi `UpdatePresensi`:
   - Memanggil fungsi `cek.UpdatePresensi` untuk mengupdate data presensi dengan parameter yang diperlukan, seperti ID, nilai longitude, latitude, lokasi, nomor telepon, check-in, dan biodata.
   - Jika terjadi kesalahan saat memperbarui data, akan mengembalikan respons HTTP dengan status `http.StatusInternalServerError` dan pesan error.

6. Mengembalikan respons berhasil:
   - Jika pembaruan data berhasil, akan mengembalikan respons HTTP dengan status `http.StatusOK` dan pesan "Data successfully updated".

Fungsi `UpdateData` ini berfungsi sebagai handler untuk menerima permintaan update data presensi dari client melalui webservice. Kemudian, data yang diterima akan diparsing dan disimpan dalam database MongoDB menggunakan fungsi `UpdatePresensi`. Setelah pembaruan data berhasil dilakukan, respons HTTP dengan status sukses akan dikirim kembali ke client. Jika terdapat kesalahan dalam proses parsing atau pembaruan data, respons HTTP dengan status kesalahan akan dikirim bersama dengan pesan error.

Pada file `url.go` di folder url kita tambah endpoint baru
```go
page.Put("/update/:id", controller.UpdateData)
```
UpdateData adalah fungsi handler yang digunakan untuk mengelola permintaan HTTP dengan metode PUT pada PATH atau endpoint `"/upd/:id"`.

Jangan lupa selalu menjalankan perintah
```go
go mod tidy
```

### Delete Data
Kembali pada file `coba.go` di folder controller kita tambahkan fungsi `DeletePresensiByID`
```go
func DeletePresensiByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = cek.DeletePresensiByID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}
```

Penjelasan:
- Fungsi `DeletePresensiByID` menggunakan parameter `c *fiber.Ctx`, yang merupakan konteks dari permintaan HTTP yang diterima.
- Pertama, kita mendapatkan nilai `id` dari parameter URL menggunakan `c.Params("id")`.
- Jika `id` kosong, kita mengembalikan respons dengan status HTTP 500 (Internal Server Error) dan pesan error.
- Selanjutnya, kita mengonversi `id` menjadi `primitive.ObjectID` menggunakan `primitive.ObjectIDFromHex(id)`.
- Jika terjadi kesalahan dalam konversi, kita mengembalikan respons dengan status HTTP 400 (Bad Request) dan pesan error.
- Setelah itu, kita memanggil fungsi `DeletePresensiByID` dengan menyediakan `objID`, koneksi MongoDB (`config.Ulbimongoconn`), dan nama koleksi `"presensi"`.
- Jika terjadi kesalahan dalam penghapusan data, kita mengembalikan respons dengan status HTTP 500 (Internal Server Error) dan pesan error.
- Jika penghapusan berhasil dilakukan, kita mengembalikan respons dengan status HTTP 200 (OK) dan pesan sukses.

Pada file `url.go` di folder url kita tambah endpoint baru
```go
page.Delete("/delete/:id", controller.DeletePresensiByID)
```
DeletePresensiByID adalah fungsi handler yang digunakan untuk mengelola permintaan HTTP dengan metode DELETE pada PATH atau endpoint "/delete/:id".

Jangan lupa selalu menjalankan perintah
```go
go mod tidy
```

Kemudian lakukan commit dan push pada repo github dan heroku.
```sh
git push heroku main
```

### Testing UPDATE dan DELETE

Sebelum melakukan update dan delete kita coba dulu get id presensi (1 data) yang akan kita gunakan untuk update dan delete pada praktikum kali ini. Cek pada Postman menggunakan method Get dan kemudian cari berdasarkan ID, bisa dilihat data berhasil ditampilkan. (pastikan endpoint `URLHEROKU/presensi/ID`). 

catatan : untuk ID gunakan yang ada pada MongoDB kalian.

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/47227b1f-eb4b-424f-bde7-874081e44d22)

- Update

Jika sudah, kita coba testing terlebih dahulu via postman sebelum kita implementasi pada frontend. 

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/94d321bf-0033-4664-bbe3-18569ef40af3)

Pada Postman gunakan method PUT, masukkan endpoint webservice (untuk ID disesuaikan saja dengan yang ada pada MongoDB), kemudian pilih Body->raw->JSON isikan dengan kode berikut
```json
{
    "longitude" : 91.31111,
    "latitude" : 125.56165,
    "location" : "ULBI",
    "phone_number" : "628123456789",
    "checkin" : "Masuk",
    "biodata" : {
        "nama": "Coba Update",
		"phone_number": "628123456789",
		"jabatan": "Tidak Diketahui",
		"jam_kerja": [{
                "durasi":     8,
                "jam_masuk":  "08:00",
                "jam_keluar": "16:00",
                "gmt":        7,
                "hari":       ["Senin", "Rabu", "Kamis"],
                "shift":      1,
                "piket_tim":  "Piket A"
            }],
		"hari_kerja":  ["Senin", "Kamis"]
	}
}
```

Klik send, maka response terlihat seperti berikut

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/6c756a9f-56b2-4b3f-a98f-54bc938a27e8)

Kemudian coba kalian cek pada Postman menggunakan method Get dan kemudian cari berdasarkan ID yang kita ubah saat melakukan update, bisa dilihat data berhasil berubah. (pastikan endpoint **URLHEROKU/presensi/ID** bisa digunakan karena nanti akan digunakan pada frontend)

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/6d110217-ff7a-4b43-89af-55891e118db5)

Pada tahap ini update data sudah aman, selanjutnya kita akan melakukan delete data.

- Delete

Pada Postman gunakan method DELETE, masukkan endpoint webservice (untuk ID disesuaikan saja dengan yang ada pada MongoDB atau gunakan ID yang sudah digunakan pada method update sebelumnya)

**PASTIKAN ENDPOINTNYA : `URLHEROKU/delete/ID` DAN METHODNYA ADALAH `DELETE`**

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/0a9e868b-9cf2-4820-8c95-23178a4aa119)

Klik tombol send, maka response terlihat seperti berikut

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/dc633b6b-20c4-4234-ad2f-f07c85e3b339)

Kemudian coba kalian cek pada Postman menggunakan method Get dan kemudian cari berdasarkan ID yang sudah kita hapus, seharusnya datanya sudah tidak ada seperti gambar di bawah. (pastikan endpoint **`URLHEROKU/presensi/ID`** berjalan untuk melanjutkan pada frontend)

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/159ae545-0123-4e27-8906-028e89b56e3e)

Pada tahap ini testing update dan delete data pada postman sudah aman, selanjutnya kita akan melakukan PUT dan DELETE data pada frontend

# FRONTEND

### Ubah pada tampilan GET
- Pada file `index.html` tambah header tabel Aksi pada line 298
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/db538554-41ba-4d6b-854c-d88b8e46b282)
- Pada file `js/temp/table.js` tambahkan 2 button baru untuk edit dan delete (berikut contoh gambar dan codenya)
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/8a1c2e8c-314b-4be0-a292-870cdfab3784)
```html
<th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
    <a type="button" href="edit.html?presensiId=#IDEDIT#"> Edit
    </a>
    |
    <button type="button" id="del_button" onclick="deleteData('#IDHAPUS#')"> Delete
    </button>
</th>
```
- Pada file `js/controller/get.js` tambahkan code untuk mengidentifikasi ID untuk button edit dan delete
```js
.replace("#IDEDIT#", value._id)
.replace("#IDHAPUS#", value._id)
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/e33b5bbb-f8ff-45f8-b156-f1bccdbbead0)
- Kemudian lakukan push pada repo frontend masing-masing
- Jika sudah, seharusnya terdapat 2 tombol Edit dan Delete pada halaman index.html
![image](https://github.com/indrariksa/WS/assets/26703717/b6b36003-866d-4304-abd3-3948194cba0a)

### Update Data
Selanjutnya kita masuk ke materi untuk mengubah data

- Download file `edit.html` [DISINI](https://github.com/Bukulapak/WebService2024/tree/main/Week11/Praktikum/edit.html), kemudian simpan pada frontend kalian masing-masing di dalam folder template
- Kemudian pada baris di bawah sebelum penutup tag body tambahkan 2 kode berikut pada line 414 dan 415
```html
<script type="module" src="../js/fetch_edit.js"></script>
<script type="module" src="../js/controller/put.js"></script>
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/6a3a8e1a-39f8-490e-8d85-6bb115ead4cf)
- Kemudian buat file `url_get_detail.js` pada `folder js/config` (**NOTE : UNTUK URL HEROKU SESUAIKAN DENGAN URL KALIAN MASING-MASING**)
```js
//Mendapatkan parameter dari URL
const urlParams = new URLSearchParams(window.location.search);
const presensiId = urlParams.get("presensiId");

export let urlFetch = "URLHEROKU/presensi/" + presensiId;
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/f1c77247-bdf8-44ee-8b9f-38d9f8e22adf)

> Fungsi kode yang Anda berikan digunakan untuk mendapatkan parameter `presensiId` dari URL dan membangun URL pengambilan (`urlFetch`) dengan menggunakan nilai tersebut. Berikut adalah penjelasan singkat tentang kode tersebut:
> 1. `const urlParams = new URLSearchParams(window.location.search);`: Kode ini membuat objek `URLSearchParams` dari `window.location.search`, yang mewakili string query parameter di URL saat ini.
> 2. `const presensiId = urlParams.get('presensiId');`: Kode ini menggunakan objek `urlParams` untuk mendapatkan nilai dari parameter query `'presensiId'`. Jadi jika URL-nya adalah `https://example.com/?presensiId=12345`, maka nilai `presensiId` akan menjadi `'12345'`.
> 3. `export let urlFetch = 'https://ws-indra2024-878f7e6fab92.herokuapp.com/presensi/' + presensiId;`: Kode ini membangun URL pengambilan (`urlFetch`) dengan menggabungkan string `'https://ws-indra2024-878f7e6fab92.herokuapp.com/presensi/'` dengan nilai `presensiId` yang telah diperoleh dari parameter query. Hasilnya akan menjadi URL lengkap yang akan digunakan dalam permintaan fetch.
> Dengan kode tersebut, kalian dapat menggunakan `urlFetch` untuk mengirim permintaan fetch ke URL yang dibangun dengan parameter `presensiId`.
- Selanjutnya kita akan menampilkan data pada form di bawah
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/d88957a7-411b-424a-a1b0-b180adf5e0b3)

Untuk dapat menampilkannya, kita perlu membuat file `edit.js` pada folder `js/controller` yaitu sebagai berikut
```js
export function isiData(results) {  
    const inputMapping = [
      { id: 'nama', path: 'biodata.nama' },
      { id: 'jabatan', path: 'biodata.jabatan' },
      { id: 'hari_kerja', path: 'biodata.hari_kerja' },
      { id: 'phone_number', path: 'phone_number' },
      { id: 'location', path: 'location' },
      { id: 'latitude', path: 'latitude' },
      { id: 'longitude', path: 'longitude' },
      { id: 'checkin', path: 'checkin' },
      { id: 'jam_masuk', path: 'biodata.jam_kerja', index: 0, property: 'jam_masuk' },
      { id: 'jam_keluar', path: 'biodata.jam_kerja', index: 0, property: 'jam_keluar'  },
      { id: 'durasi', path: 'biodata.jam_kerja', index: 0, property: 'durasi'  },
    ];
  
    inputMapping.forEach(({ id, path, index, property }) => {
      const inputElement = document.getElementById(id);
      const value = getNestedValue(results, path, index, property);
      inputElement.value = value;
    });
}
  
function getNestedValue(obj, path, index, property) {
    const value = path.split('.').reduce((value, key) => (value && value[key]) ? value[key] : '', obj);
    // console.log(`Value at path ${path}:`, value);
  
    if (Array.isArray(value) && value.length > index && value[index].hasOwnProperty(property)) {
      return value[index][property];
    }
  
    return value;
}
```
> Kode di atas adalah fungsi `isiData` yang digunakan untuk mengisi nilai-nilai elemen input berdasarkan hasil yang diterima. Berikut adalah penjelasan singkat dan jelas untuk kode tersebut:
> 1. `const inputMapping`: Ini adalah array yang berisi objek-objek yang mendefinisikan pemetaan antara `id` elemen input dan `path` (jalur) di dalam objek `results` untuk mendapatkan nilai yang sesuai.
> 2. `inputMapping.forEach(({ id, path, index, property }) => { ... })`: Ini melakukan iterasi melalui setiap objek dalam `inputMapping`. Dalam setiap iterasi, kita mendapatkan `id`, `path`, `index`, dan `property` dari objek saat ini.
> 3. `const inputElement = document.getElementById(id);`: Ini mencari elemen input berdasarkan `id` yang diperoleh dari objek saat ini.
> 4. `const value = getNestedValue(results, path, index, property);`: Ini memanggil fungsi `getNestedValue` dengan parameter `results` (hasil yang diterima), `path` (jalur), `index`, dan `property` untuk mendapatkan nilai yang sesuai.
> 5. `inputElement.value = value;`: Ini mengisi nilai elemen input dengan nilai yang diperoleh.
> 6. `function getNestedValue(obj, path, index, property) { ... }`: Ini adalah fungsi yang digunakan untuk mendapatkan nilai yang terdalam di dalam objek berdasarkan jalur yang diberikan. Fungsi ini menggunakan metode `reduce` dan `split` untuk mengakses nilai yang terdalam sesuai dengan jalur yang diberikan. Jika nilai tersebut adalah array dan memiliki indeks dan properti yang sesuai, maka nilai tersebut akan dikembalikan. Jika tidak, nilai tersebut akan dikembalikan secara langsung.
> Dengan menggunakan fungsi `isiData` dan `getNestedValue`, nilai-nilai dari objek `results` akan dipetakan ke elemen-elemen input yang sesuai berdasarkan pemetaan yang diberikan.
- Sabar guys belum beres, kita harus buat file `fetch_edit.js` pada folder `js` dengan struktur folder seperti berikut

    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/00876cfe-b4b1-4aea-8c51-77e49ef03c2c)

    kemudian isi `fetch_edit.js` sebagai berikut
    ```js
    import { get } from "https://bukulapak.github.io/api/process.js";
    import { isiData } from "./controller/edit.js";
    import { urlFetch } from "./config/url_get_detail.js";
    get(urlFetch, isiData);
    ```
- Jika sudah, kalian bisa push ke repo dan heroku masing-masing kemudian coba running, jangan lupa lakukan hard refresh(Ctrl+F5 atau Ctrl+Shift+R, dll). Maka hasilnya akan terisi data pada formnya ketika kalian mengklik tombol edit (kalau belum bisa coba baca lagi yang teliti barangkali ada yang terlewat)
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/c8cc50be-c973-4704-8774-cb6c09a2e7d3)

- Eitssss belum bisa update ya, kita baru menampilkan data pada form saja, so kita belum membuat function untuk mengubah data yang dikirim dari form ya
---

- Yuk lanjut, selanjutnya kita akan membuat function untuk melakukan edit data. Buat file `put.js` pada folder `js/controller` seperti berikut

    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/aabe7f98-2e83-457c-aabf-0b16294610ed)

Panjang ya kode nya? Copy paste dari file `post.js` aja cuma sedikit kok perubahannya, teliti ya jangan asal copy paste.
> Kode di atas adalah implementasi fungsi untuk mengirim data melalui metode PUT ke sebuah URL. Berikut adalah penjelasan untuk kode di atas:
> 1. `import { putData } from "https://bukulapak.github.io/api/process.js";`: Ini mengimpor fungsi `putData` dari modul `process.js` yang berada di URL eksternal "https://bukulapak.github.io/api/process.js". Fungsi ini digunakan untuk melakukan permintaan PUT.
> 2. `import { onClick, getValue } from "https://bukulapak.github.io/element/process.js";`: Ini mengimpor fungsi `onClick` dan `getValue` dari modul `process.js` yang berada di URL eksternal "https://bukulapak.github.io/element/process.js". Fungsi `onClick` digunakan untuk menangani peristiwa klik pada elemen HTML, sedangkan `getValue` digunakan untuk mendapatkan nilai dari elemen HTML.
> 3. `import { urlPUT, AmbilResponse} from "../config/url_put.js";`: Ini mengimpor variabel `urlPUT` dan `AmbilResponse` dari modul `url_put.js` yang berada di direktori "../config/url_put.js". Variabel ini berisi URL target untuk permintaan PUT dan fungsi untuk mengambil respons.
> 4. `function pushData() { ... }`: Ini adalah fungsi `pushData` yang akan dipanggil saat tombol diklik. Fungsi ini berfungsi untuk mengumpulkan nilai-nilai dari elemen HTML dan membentuk objek data yang akan dikirim melalui permintaan PUT.
> 5. `var hari_kerja = getValue("hari_kerja");`: Ini mendapatkan nilai dari elemen HTML dengan ID "hari_kerja" menggunakan fungsi `getValue` dan menyimpannya dalam variabel `hari_kerja`.
> 6. `let data = { ... }`: Ini merupakan pembentukan objek `data` yang akan dikirim melalui permintaan PUT. Nilai-nilai dari elemen HTML dan variabel `hari_kerja` dimasukkan ke dalam objek ini.
> 7. `putData(urlPUT, data, AmbilResponse);`: Ini memanggil fungsi `putData` dengan parameter `urlPUT` (URL target), `data` (objek data yang akan dikirim), dan `AmbilResponse` (fungsi yang akan menangani respons).
> 8. `onClick("button", pushData);`: Ini mengaitkan fungsi `pushData` dengan peristiwa klik pada elemen HTML dengan tag "button".
> Dengan demikian, saat tombol dengan tag "button" diklik, fungsi `pushData` akan dijalankan, mengumpulkan nilai-nilai dari elemen-elemen HTML, membentuk objek data, dan mengirimnya melalui permintaan PUT ke URL target.

- Selanjutnya buat file `url_put.js` pada folder `js/config` seperti berikut (**NOTE : UNTUK URL HEROKU SESUAIKAN DENGAN URL KALIAN MASING-MASING**)
    ```js
    const urlParams = new URLSearchParams(window.location.search);
    const presensiId = urlParams.get('presensiId');

    export let urlPUT = "https://ws-indra2024-878f7e6fab92.herokuapp.com/update/" + presensiId;

    export function AmbilResponse(result) {
        console.log(result); //menampilkan response API pada console
        alert(result.message); //menampilkan response API pada alert
        window.location.href = "index.html"; //reload halaman setelah klik ok pada alert
    }
    ```
    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/77cd9014-1862-405f-b66a-2e6ac9e866f3)
- Jika sudah, kalian bisa push ke repo dan tunggu deploy github pages berhasil kemudian coba buka browser, jangan lupa lakukan hard refresh(Ctrl+F5 atau Ctrl+Shift+R, dll). 
- Coba lakukan update data seperti gambar di bawah yang menampilkan response bahwa data berhasil dilakukan update
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/05aa82fa-d6c1-4137-ab95-89a9c03e5aad)

---

**Masih semangat kan? kita lanjut ke delete data**
### Delete Data

- Kita buat file `delete.js` di dalam folder `js/controller` yaitu sebagai berikut
```js
function deleteData(IDHAPUS) {
    var presensiId = IDHAPUS;
    var target_url = "URLHEROKU/delete/" + presensiId;

    var requestOptions = {
        method: 'DELETE',
        redirect: 'follow'
    };

    fetch(target_url, requestOptions)
        .then(response => response.json())
        .then(result => {
            alert(result.message);
            location.reload();
        })
        .catch(error => console.log('Error:', error));
}
```
**BACA!!!!! URLHEROKU di atas diisi dengan url kalian masing-masing**
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/eefae665-8ee6-4f6d-b7c1-26a23576f7ce)

> Kode fungsi `deleteData()` yang akan dipanggil saat tombol delete diklik. Berikut adalah penjelasan dari kode tersebut:
> - Variabel `presensiId` di dapat dari #IDHAPUS# yang berasal dari get.js yang memiliki value id (_id).
> - Membentuk URL target dengan menggabungkan nilai `presensiId` ke dalam URL dasar.
> - Membuat objek `requestOptions` yang berisi konfigurasi untuk metode `fetch`, dengan menggunakan metode HTTP DELETE.
> - Mengirim permintaan DELETE ke URL target menggunakan metode `fetch`.
> - Menangani respon dengan metode `then()`:
	> - Mengonversi respon menjadi objek JSON menggunakan metode `response.json()`.
	> - Menampilkan pesan hasil respon menggunakan fungsi `alert()`.
	> Me-refresh halaman menggunakan `location.reload()` untuk memperbarui tampilan setelah penghapusan data.
> - Menangani kesalahan menggunakan metode `catch()` untuk menampilkan pesan error di konsol.
> Dengan cara ini, saat tombol delete diklik, fungsi `deleteData()` akan mengambil nilai `presensiId`, mengirim permintaan DELETE ke URL target, dan menampilkan pesan hasil respon sebelum memperbarui halaman.
> Disini kita tidak lagi menggunakan import modul fetch seperti pada get all untuk menampilkan seluruh data serta get detail untuk edit

- Buka kembali file `index.html`, kita lakukan import file `delete.js` yang sudah dibuat
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/b70325dc-d885-4113-9a29-23f7993317c6)
    > * Membuat elemen tombol delete dengan id del_button dan menetapkan fungsi deleteData() sebagai event handler menggunakan atribut onclick.
    > * Ketika tombol delete diklik, fungsi deleteData() akan dipanggil untuk menghapus data.


- Jika sudah, push ke repo masing-masing dan lakukan hard refresh(Ctrl+F5 atau Ctrl+Shift+R, dll), kemudian bisa dicoba klik tombol delete
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/d40be129-093d-4a4a-a8a1-2ae714158999)

- Kemudian lihat seharusnya menampilkan response berhasil kemudian klik tombol OK
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/4beaac4c-ce5a-45d3-8999-cc927a8519f3)

- Data sudah hilang
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/5005b801-798b-4edf-b496-3c2285083e1f)

- Jika kalian ingin menambahkan konfirmasi hapus data setelah menekan tombol Delete maka tambahkan kode berikut di folder `js/controller/delete.js` tambahkan kode berikut di atas function `deleteData`
    ```js
    function confirmDelete(IDHAPUS) {
        if (confirm("Apakah ingin menghapus data ID " + IDHAPUS + "?")) {
            deleteData(IDHAPUS);
        }
    }
    ```
    Kemudian buka folder `js/temp/table.js` ubah sedikit kode button Delete pada onclick dengan memanggil fungsi confirmDelete yang sudah dibuat menjadi seperti berikut
    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/21987647-7c29-4612-9bc6-81817ea21760)

- Sekarang coba klik tombol delete maka ada konfirmasi untuk penghapusan data
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/cde98b77-cc25-40fd-8504-f7551147a67a)


## Kumpulkan 
Buat file README.md dan masukkan :
- skrinsut hasil frontend update dan delete
- url github pages frontend presensi 

Dikumpulkan di dalam folder Week11/Praktikum/{NPM}
