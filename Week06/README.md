# Deployment Aplikasi Frontend

Kita akan membuat frontend yang sudah kita buat berkomunikasi dengan backend yang sudah dibuat sebelumnya. Contoh backend punya saya :
https://ws-indra2024-878f7e6fab92.herokuapp.com/presensi

Adapun langkah yang akan kita lakukan :
1. Membuat tabel di dalam kontainer dari tampilan yang sudah ada
2. Menambahkan import js dari repo [Bukulapak](https://github.com/orgs/Bukulapak/repositories) di dalam file html yang kita buat
3. Membuat dan mengisi kerangka file js

## Get Package Presensi Karyawan
Import terlebih dahulu isi collection presensi karyawan di MongoDB pada database tes_db->collection presensi [download disini](https://drive.google.com/drive/folders/13DYUisWRmeoUnAy1X3cX2rVs13zDSWUt?usp=sharing)

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/2c04794c-a996-4640-a713-9f4314876453)

Buka project Boilerplate pada VScode atau GoLand, kemudian pada terminal ketik kode berikut untuk import update terbaru dari package di bawah
```go
go get -u github.com/indrariksa/cobapakcage
```

Jika sudah pastikan import package di bawah seperti gambar pada `controller/coba.go` sudah dimasukkan
```go
"errors"
cek "github.com/indrariksa/cobapakcage/module"
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/104b51e2-1e39-46ef-91ab-6781390bde62)

Kemudian tambah fungsi berikut
* Fungsi `get by id`
```go
func GetPresensiID(c *fiber.Ctx) error {
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
	ps, err := cek.GetPresensiFromID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}
```
Pada file `url.go` di folder `url` kita tambah path baru
```go
page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
```

Rapikan depedensi
```go
go mod tidy
```

Coba run pada local terlebih dahulu
```go
go run main.go
```

Klik link yang ada pada terminal

![image](https://user-images.githubusercontent.com/26703717/228802288-5d23ebff-6b11-4141-a7a3-f53b19e7bc1b.png)

Coba endpoint `/presensi/id` (http://127.0.0.1:8080/presensi/645df264d1b7263ad2710ec9)
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/650ee890-3a02-46a4-86cd-148527908352)

Jika sudah berjalan lancar semua, lakukan commit dan push pada repo github dan heroku.

## FRONTEND (Kontainer, Tabel, tag, Content)

Pertama kita buat akan buat repo baru untuk frontend kemudian melakukan deploy di github pages. 
Buat repo dengan klik tanda "+" di pojok kanan atas, kemudian pilih new repository

![image](https://user-images.githubusercontent.com/26703717/228426421-80029fad-8cf1-4017-ae30-c2d3ce3885ae.png)

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/b9f91d15-5e76-40ca-8869-1b659d10f75f)


Kemudian clik pada button `code`, pilih `copy` https atau ssh nya

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/10843e89-95af-4db0-be27-787f49cf532f)

Buka gitbash di directory yang kita inginkan, kemudian ketik
```
git clone {hasil copas dari https}
Contoh : git clone git@github.com:indrariksa/fe_ulbi2024.git
```

Setelah itu gunakan template frontend https://github.com/Bukulapak/WebService2024/tree/main/Week06/Praktikum/TEMPLATE kemudian simpan pada directory repo kalian (contoh copy paste hasil download template ke dalam folder fe_ulbi2024 yang merupakan repo yang baru diclone). Kemudian buka pada VScode.

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/0b97ad8e-9bee-4877-9bcc-a959e5f43722)

Hal yang pertama kali kita lakukan adalah membuka file `index.html` dan menambahkan tag di dalam section head untuk memanggil file `js/fetch.js`.
```html
<script type="module" src="../js/fetch.js"></script>
```
![image](https://user-images.githubusercontent.com/26703717/228425589-dd68591d-a724-4167-a259-0f78b80a5c64.png)

Kemudian lakukan setting github pages (contoh disini : https://github.com/Bukulapak/WebService2024/blob/main/Week02/README.md#setting-repo-untuk-github-pages)

### Pengambilan Data dari Backend

Kita buat file `fetch.js` di dalam folder `js` yang berisi.
```js
import { get } from "https://bukulapak.github.io/api/process.js"; 
let urlAPI = "https://ws-indra2024-878f7e6fab92.herokuapp.com/presensi";
get(urlAPI,isiTablePresensi);
function isiTablePresensi(results){
    console.log(results);
}
```
Untuk variabel **urlAPI** isikan dengan link heroku masing-masing mahasiswa yang berisi data presensi latihan. Kemudian kita commit dan push, kemudian tunggu hingga centang hijau pertanda github pages sudah terdeploy dengan baik

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/f0c028ef-5455-45e1-b50f-6e08a01ca291)

kemudian kita akses url github pages nya. Kita lakukan inspect dan masuk ke tab console terdapat error CORS tampak sebagai berikut.

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/3266e591-1338-4983-939b-fd2e2ab0a411)

Artinya kita perlu mendaftarkan url (contoh url punya saya adalah 'https://indrariksa.github.io') ke dalam config `cors.go` di repo boilerplate (ws-indra2024).

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/4eaa5419-6e03-4507-8273-8a573f8166d3)

Simpan, commit dan push ke heroku kemudian kita ujicoba lagi frontend kita. Liat console kembali sudah terlihat hasil dari backend

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/7217871d-0a46-4bd8-93a4-5d9ce9339194)

Karena hasil dari backend berupa array dari json object. maka kita ubah kode program tambahkan looping foreach pada `fetch.js`
```js
import { get } from "https://bukulapak.github.io/api/process.js";
import { setInner } from "https://bukulapak.github.io/element/process.js";
let urlAPI = "https://ws-indra2024-878f7e6fab92.herokuapp.com/presensi";
get(urlAPI,isiTablePresensi);
function isiTablePresensi(results){
    console.log(results);
    results.forEach(isiRow);
}
function isiRow(value){
    console.log(value)
}
```
Hasilnya kita dapatkan object yang keluar dari console.log fungsi isiRow. Lakukan commit dan push, kemudian lakukan hard refresh menggunakan SHIFT + F5 atau CTRL + F5 setiap ada perubahan pada file js

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/a6b2e4a0-1d24-47b3-a57e-6cefbd0aae98)

### Memasukkan data ke dalam tabel

Sekarang kita masukkan ke dalam tabel. Sebelumnya edit `index.html` tambahkan id pada tabelnya
```html
<table class="w-full" id="iniTabel">
```

Kita ambil script TR Tabel HTML yang kemudian kita taruh di file table.js yang di deklarasikan ke dalam variabel isiTabel.
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/3d81c1a3-dbd6-43bc-bb0c-484ffc99eb3c)
Cut semua code yang diblok seperti gambar di bawah kemudian paste pada table.js
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/9c8acf7e-b725-488f-a1ba-d9c877f5f380)

isi file `table.js` (di dalam folder `js`)
```js
export let isiTabel = 
`
<tr class="h-18 border-b border-coolGray-100">
    <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="flex items-center -m-2">
            <div class="w-auto p-2">
                <input class="w-4 h-4 bg-white rounded" type="checkbox">
            </div>
            <div class="w-auto p-2">
                <div class="flex items-center justify-center w-10 h-10 text-base font-medium text-#WARNALOGO#-600 bg-#WARNALOGO#-200 rounded-md">ULBI</div>
            </div>
            <div class="w-auto p-2">
                <p class="text-xs font-semibold text-coolGray-800">#NAMA#</p>
                <p class="text-xs font-medium text-coolGray-500">#NOHP#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#JABATAN#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-center">#STATUS#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-500 text-left">#LOKASI#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#HARIKERJA#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#JAMKERJA#</th>
    <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="w-auto p-2">
        <p class="text-xs font-semibold text-coolGray-800">#JAMMASUK#</p>
        <p class="text-xs font-medium text-coolGray-500">#JAMKELUAR#</p>
        </div>
    </th>
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
</tr>
`
```
Dimana #NAMA#, #NOHP#, #JABATAN#, #STATUS# dst adalah variabel yang akan kita replace dengan data dari API. File `table.js` ditaruh di dalam folder `js`, sehingga strukturnya tampak sebagai berikut.

![image](https://user-images.githubusercontent.com/26703717/228436793-242e36e9-c7c8-4c61-8a47-195a0aa96051.png)

kita ubah `fetch.js` lagi menjadi
```js
import { get } from "https://bukulapak.github.io/api/process.js";
import { addInner } from "https://bukulapak.github.io/element/process.js";
import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabel } from "./table.js";
let urlAPI = "https://ws-indra2024-878f7e6fab92.herokuapp.com/presensi";
get(urlAPI, isiTablePresensi);
function isiTablePresensi(results) {
    results.forEach(isiRow);
}
function isiRow(value) {
    let content = 
    isiTabel.replace("#NAMA#", value.biodata.nama)
            .replace("#NOHP#", value.biodata.phone_number)
            .replace("#JABATAN#", value.biodata.jabatan)
            .replace("#LOKASI#", value.location)
            .replace("#STATUS#", value.checkin)
            .replace("#HARIKERJA#", value.biodata.hari_kerja)
            .replace("#WARNA#", getRandomColor())
            .replace(/#WARNALOGO#/g, getRandomColorName());
        addInner("iniTabel", content);
}
```
jadilah tabel presensi karyawan sudah terisi

![image](https://user-images.githubusercontent.com/26703717/228453151-18138ad9-8a92-47e7-8bf7-780f0921e9c9.png)

### Membuat boilerplate framework frontend

Sekarang kita rapihkan kode program js kita, sehingga bisa terbaca orang lain dan sesuai dengan sudut pandang framework pada umumnya. Kita ubah dengan struktur folder config,controller,template dengan isi file masing-masing seperti terlihat pada gambar di bawah ini.

![image](https://user-images.githubusercontent.com/26703717/228769390-f7aa93f9-8f67-4139-9594-d29ada559f9b.png)

file `fetch.js` berisi
```js
import { get } from "https://bukulapak.github.io/api/process.js";
import { isiTablePresensi } from "./controller/get.js";
import { urlAPI } from "./config/url.js";
get(urlAPI, isiTablePresensi);
```
file `config/url.js` berisi :
```js
export let urlAPI = "https://ws-indra2024-878f7e6fab92.herokuapp.com/presensi";
```
untuk urlAPI di atas isikan dengan url Heroku kalian

file `controller/get.js` berisi :
```js
import { addInner } from "https://bukulapak.github.io/element/process.js";
import { getRandomColor, getRandomColorName } from "https://bukulapak.github.io/image/process.js";
import { isiTabel } from "../temp/table.js";
export function isiTablePresensi(results) {
    results.forEach(isiRow);
}
function isiRow(value) {
    let content =
        isiTabel.replace("#NAMA#", value.biodata.nama)
            .replace("#NOHP#", value.biodata.phone_number)
            .replace("#JABATAN#", value.biodata.jabatan)
            .replace("#LOKASI#", value.location)
            .replace("#STATUS#", value.checkin)
            .replace("#HARIKERJA#", value.biodata.hari_kerja)
            .replace("#WARNA#", getRandomColor())
            .replace(/#WARNALOGO#/g, getRandomColorName());
    addInner("iniTabel", content);
}
```
file `temp/table.js` berisi :
```js
export let isiTabel = 
`
<tr class="h-18 border-b border-coolGray-100">
    <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="flex items-center -m-2">
            <div class="w-auto p-2">
                <input class="w-4 h-4 bg-white rounded" type="checkbox">
            </div>
            <div class="w-auto p-2">
                <div class="flex items-center justify-center w-10 h-10 text-base font-medium text-#WARNALOGO#-600 bg-#WARNALOGO#-200 rounded-md">ULBI</div>
            </div>
            <div class="w-auto p-2">
                <p class="text-xs font-semibold text-coolGray-800">#NAMA#</p>
                <p class="text-xs font-medium text-coolGray-500">#NOHP#</p>
            </div>
        </div>
    </th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-left">#JABATAN#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-800 text-center">#STATUS#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-coolGray-500 text-left">#LOKASI#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#HARIKERJA#</th>
    <th class="whitespace-nowrap px-4 bg-white text-sm font-medium text-#col#-500 text-left">#JAMKERJA#</th>
    <th class="whitespace-nowrap px-4 bg-white text-left">
        <div class="w-auto p-2">
        <p class="text-xs font-semibold text-coolGray-800">#JAMMASUK#</p>
        <p class="text-xs font-medium text-coolGray-500">#JAMKELUAR#</p>
        </div>
    </th>
    <th class="whitespace-nowrap pr-4 bg-white text-sm font-medium text-coolGray-800">
        <svg class="ml-auto" width="16" height="16" viewbox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M8 6.66669C7.73629 6.66669 7.47851 6.74489 7.25924 6.89139C7.03998 7.0379 6.86908 7.24614 6.76816 7.48978C6.66724 7.73341 6.64084 8.0015 6.69229 8.26014C6.74373 8.51878 6.87072 8.75636 7.05719 8.94283C7.24366 9.1293 7.48124 9.25629 7.73988 9.30773C7.99852 9.35918 8.26661 9.33278 8.51025 9.23186C8.75388 9.13094 8.96212 8.96005 9.10863 8.74078C9.25514 8.52152 9.33333 8.26373 9.33333 8.00002C9.33333 7.6464 9.19286 7.30726 8.94281 7.05721C8.69276 6.80716 8.35362 6.66669 8 6.66669ZM3.33333 6.66669C3.06963 6.66669 2.81184 6.74489 2.59257 6.89139C2.37331 7.0379 2.20241 7.24614 2.10149 7.48978C2.00058 7.73341 1.97417 8.0015 2.02562 8.26014C2.07707 8.51878 2.20405 8.75636 2.39052 8.94283C2.57699 9.1293 2.81457 9.25629 3.07321 9.30773C3.33185 9.35918 3.59994 9.33278 3.84358 9.23186C4.08721 9.13094 4.29545 8.96005 4.44196 8.74078C4.58847 8.52152 4.66667 8.26373 4.66667 8.00002C4.66667 7.6464 4.52619 7.30726 4.27614 7.05721C4.02609 6.80716 3.68696 6.66669 3.33333 6.66669ZM12.6667 6.66669C12.403 6.66669 12.1452 6.74489 11.9259 6.89139C11.7066 7.0379 11.5357 7.24614 11.4348 7.48978C11.3339 7.73341 11.3075 8.0015 11.359 8.26014C11.4104 8.51878 11.5374 8.75636 11.7239 8.94283C11.9103 9.1293 12.1479 9.25629 12.4065 9.30773C12.6652 9.35918 12.9333 9.33278 13.1769 9.23186C13.4205 9.13094 13.6288 8.96005 13.7753 8.74078C13.9218 8.52152 14 8.26373 14 8.00002C14 7.6464 13.8595 7.30726 13.6095 7.05721C13.3594 6.80716 13.0203 6.66669 12.6667 6.66669Z" fill="#WARNA#"></path>
        </svg>
    </th>
</tr>
`
```

Contoh frontend yang sudah jadi dengan alamat github pages : https://indrariksa.github.io/fe_ulbi2024/template

Note : Lakukan HARD REFRESH menggunakan CTRL + F5 atau SHIFT+F5 setiap ada perubahan pada file js

## Kumpulkan 
Buat file README.md dan masukkan skrinsut hasil dan url heroku di folder Week06/Praktikum/{NPM}
