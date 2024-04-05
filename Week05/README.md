## Dashboard Heroku dan Heroku CLI (catatan : JIKA CREDIT HEROKU MASIH $0 jangan lanjut dulu ke tahap ini)
Sebelum mulai kita subscribe eco dyno pada heroku, di menu Billing 
![image](https://user-images.githubusercontent.com/26703717/225517047-8c697835-602d-4ef5-ac8d-2ea39729d6ef.png)

Setelah login, masuk ke laman https://dashboard.heroku.com/apps. Maka akan muncul list aplikasi yang sudah kita buat. Klik tombol New di pojok kanan atas dan pilih Create new app untuk melakukan deployment aplikasi baru kita.
![image](https://user-images.githubusercontent.com/26703717/225284159-bdd4a4d0-32e7-4887-ad16-a2a9b0f203c7.png)

Masukkan nama aplikasi kita, dan pilih lokasi server kita apakah amerika atau eropa (saya pilih amerika). Kemudian klik Create app.
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/08132915-6b2b-405f-b6c1-5fe09c1640a8)

Lakukan instalasi Heroku CLI, untuk menghubungkan komputer kita dengan server heroku. Link Instalasi https://devcenter.heroku.com/articles/heroku-cli
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/a8b9fae5-bbba-44f4-b44e-14cd4ff56bf8)

Pilih 64-bit installer jika laptop kalian 64 bit
![image](https://user-images.githubusercontent.com/26703717/225332966-01c6bdfd-9ed7-4ed5-bb0c-24d5eb80f2c7.png)

Lanjutkan sampai selesai langkah instalasinya

![image](https://user-images.githubusercontent.com/15622730/224493999-d208a079-df02-4bcf-b6f6-618c52414d54.png)

Buka git bash dan ketik heroku login, maka akan muncul button login heroku di browser, klik saja login. 
![image](https://user-images.githubusercontent.com/26703717/225334281-53356654-fff8-47b0-b7fc-7c9d7a0b6f3b.png)

Kemudian di gitbash akan ada tulisan Logged in as .....
![image](https://user-images.githubusercontent.com/26703717/225334590-fdbef3bf-969b-4449-9079-754413e31cfa.png)

## Deployment Boiler Plate

Disini kita akan mencoba testing deployment ke Heroku. Aplikasi web yang akan dilakukan deployment adalah Boiler Plate yang berada di repo https://github.com/Bukulapak/boilerplate2024

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/19da842e-911c-4f1f-8b73-5e3947bfb8a0)
Kita lakukan fork ke repositori kita, kita beri nama sesuai dengan nama aplikasi di heroku (contoh di atas : ws-indra2024).
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/cfb74e60-0259-40de-9cab-37989dcf5fe5)

Kemudian, lakukan clone repo ke komputer kita.
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/460b97ba-ec82-4a98-87ef-8da1bf042ead)
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/2f0305b4-485e-4bf1-bf03-db82048c8e94)

Setelah di clone menggunakan git bash kemudian masuk ke direktori repo di PC kita. Lakukan add remote heroku sesuai nama aplikasi yang sudah kita buat di heroku dengan perintah
```sh
cd ws-indra2024/

heroku git:remote -a ws-indra2024
```
dimana ws-indra2024 adalah nama aplikasi kita di heroku.

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/cb7b56f3-90fb-4258-8a53-9b4a45ee3192)

Kita cek dulu apakah remote kita sudah sesuai (mengarah ke repositori github kita dan heroku dengan nama app punya kalian masing-masing *(punya saya ws-indra2024) )
```sh
git remote -v
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/1d4641a7-412f-4aee-b78b-9300791dd982)

Sekarang buka code editor (VScode atau GoLand)
Selanjutnya kita harus melakukan go mod init, karena terlihat belum ada file gomod dan go.sum di folder repo. Kita lakukan dulu init kemudian go mod tidy dan berhasil.
```sh
go mod init github.com/{username github}/{nama repo}

go mod tidy
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/7df1a799-1cdc-4f0e-8b02-55e07cf12df1)

Muncul error berikutnya. Tenang saja jangan panik. Karena error kita jadi belajar. terlihat tiba tiba terdeteksi package iteung, ini berarti kode program di dalam file masih import nama package yang lama. Tinggal kita ganti semua importnya dengan nama package yang di deklarasikan pada saat kita go mod init tadi.
Kita buka VS Code, kita buka terminal pada bagianPROBLEMS akan mengeluarkan error yang serupa. Selesaikan satu persatu di masing-masing file yang muncul problems.

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/298b5a46-1f73-4af3-a04b-56b71e02660f)

Masalah selesai dengan mengganti import menuju direktori awal kita melakukan go mod init di ketiga file yaitu main.go dan url.go.

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/9f064827-0aca-4cdc-bdfa-d9329104c838)

Jika terdapat error pada controller.go line 22 ubah RunModule menjadi RunModuleLegacy
![image](https://user-images.githubusercontent.com/26703717/225715134-58dc071b-f169-41f7-ae90-d2785c0c4901.png)

Sebelum melakukan push kita cek dulu apakah remote kita sudah sesuai (mengarah ke repositori github kita dan heroku dengan nama app ws-indra2024)
```sh
git remote -v
```
![image](https://user-images.githubusercontent.com/26703717/225715293-cf61b1f6-e9ad-41ae-9ebc-1bdcafb431d4.png)
Pada gambar di atas terlihat bahwa remote sudah sesuai mengarah ke repo kita dan ws-indra2024 pada heroku

Sekarang kita tinggal add commit dan push ke github dan heroku kembali.
Tapi, sebelum itu kita cek dulu apakah file yang akan di push sudah sesuai
```sh
git status
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/6112cf11-d1df-4b69-a5bb-2249f30e9fd4)

Jika sudah kita coba git add
```sh
git add .
git status
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/249106b9-aeb3-411f-beb9-704f5953c598)

Jika sudah hijau semua selanjutnya adalah commit
```sh
git commit -m "KOMENTAR"
```

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/794af20c-2680-4bb6-990b-482e543302b7)

Kemudian lakukan push pada repository kita dan ke heroku
```sh
git push

git push heroku main
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/02d710b8-f63e-4716-8c57-fde2f3d5723e)

Terlihat damai dan tentram akhirnya kita berhasil melakukan deployment ke heroku
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/c53c1802-f1a5-44a8-a9ea-fcdaf719d3ef)

Oke kita coba buka url aplikasi kita, dicontohkan di heroku CLI disini https://ws-indra2024.herokuapp.com/ kita coba buka apakah sesuai dengan yang diharapkan.
![image](https://user-images.githubusercontent.com/26703717/225516583-0b051e7e-2737-497e-bc09-05f6f28016e7.png)

Dan kita menemukan error selanjutnya. kita coba jalankan heroku logs --tail pada git bash atau terminal
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/0ecde335-4e4a-4079-8aff-6b55e0e609de)

Disitu tertera error terkait database yang tidak ditemukan.
Sebelum itu, pada tab Resource pastikan sudah aktif untuk Eco Dynos di aplikasi kita pada dashboard heroku. Oleh karena itu kita kunjungi kembali dashboard heroku kita.

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/9213074c-7e77-4840-8d69-9f952c3f498c)

Pastikan sudah tertera tulisan "Included in Eco subscription" seperti gambar di atas. Jika belum aktif pada bagian Menu Resources, dan pada bagian main bin/ws-indra2024 kita klik tombol tanda pensil, kita geser tombol menjadi aktif dan klik tombol confirm

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/36b59c3d-48fd-41f5-92f7-71081691e5a3)

Selanjutnya pada tab Settings, pilih Reveal Config Vars.
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/d4811e64-3268-4258-a5b1-c1d2467f47cb)

Isikan Key "MONGOSTRING" dan value didapatkan dari MongoCompass seperti cara di bawah dengan melakukan Copy connection string
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/fa058781-c551-4325-b770-bcbc66d54842)

Kemudian paste pada value di Heroku, kemudia tekan tombol Add.
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/1b3febe8-d306-40a4-ab3c-597aae9416b1)

Dan akhirnya web sudah bisa diakses.

![image](https://github.com/Bukulapak/WebService2024/assets/26703717/77aaf4d3-bec6-4589-a640-cef8033d30d8)

## Iteung Boiler Plate

### Memanggil package golang

Dari package yang sudah dibuat pada chapter sebelumnya kita akan coba panggil dari aplikasi boiler plate iteung yang sudah di deploy di heroku. Sebagai contoh kita akan mencoba memanggil package musik di https://github.com/aiteung/musik hal yang pertama kali adalah kita membuka file url.go di folder url kita tambahkan baris di dalam func Web sebuah baris kode yang memanggil controller.Homepage yaitu fungsi yang akan kita buat :


pada terminal ketikkan perintah go get package yang akan digunakan
```sh
go get github.com/aiteung/musik
```

kemudian kita buat file coba.go di folder controller kita tambahkan fungsi Homepage
```go
func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/c70fb7a5-04f6-466a-8bb2-9f84f782ab05)

pada file url.go
```go
page.Get("/checkip", controller.Homepage) //ujicoba panggil package musik
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/4b80ffb0-010b-4168-9b37-03087963d57a)

Jangan lupa selalu menjalankan perintah 
```sh
go mod tidy
```
**Kemudian lakukan commit dan push pada repo github dan heroku.**

Jika sudah, kita buka alamat aplikasi dengan menambahkan perintah di bawah dan terlihat ip address dari host heroku
```sh
"{url heroku kalian}/checkip"
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/37c7565c-dc52-48ff-a1eb-4c9d7fd4c56f)

### Deployment Package

Pada pertemuan sebelumnya. Package yang sudah dibuat bisa kita panggil di controller. Cukup dengan 3 tahap, yaitu :
1. Go get package yang akan digunakan melalui terminal di vscode
   ```sh
   go get github.com/indrariksa/cobapakcage
   ```
   ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/be11c3c9-2da4-479a-ad2b-0e4436507f32)
2. Buat fungsi di file coba.go
   ```go
    func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetAllPresensi()
	return c.JSON(ps)
    }
   ```
   dan juga menambah import di atasnya
   ```go
   cek "github.com/indrariksa/cobapakcage/module"
   ```
   ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/30b7daf3-3b5a-49fe-83f7-2a0f4634ce0b)
   
1. Definisikan alamat URL untuk akses package beserta nama fungsi di controller yang akan kita buat pada file url.go
   ```go
   page.Get("/presensi", controller.GetPresensi)
   ```
Jangan lupa selalu menjalankan perintah 
```sh
go mod tidy
```
Jika kalian belum mau melakukan push ke repo dan ingin cek running di LOCAL terlebih dahulu kalian dapat menjalankan perintah berikut
```sh
go run main.go
```
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/9f2faf18-8673-4ed7-ba58-77df133eb952)

Kemudian copy url di atas kemudian tambahkan "/presensi" dan akan tampil seperti gambar di bawah
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/ee46d734-ba14-44ce-9115-109e778fd3c6)

**Langkah di atas sudah berhasil maka selanjutnya lakukan commit push ke repo dan heroku**. Kemudian kita akses melaui url yang kita deklarasikan di file url.go. Untuk method GET bisa menggunakan browser saja, untuk POST menggunakan POSTMAN

![image](https://user-images.githubusercontent.com/26703717/225527396-66eba03d-7dd8-43ef-8320-53be5500b6ef.png)

## Kumpulkan 
Buat file README.md dan masukkan link repo boilerplate masing-masing mahasiswa dan url heroku di folder Week05/Praktikum/{NPM}
