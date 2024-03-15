# HTTP Header and Body Capture

## Setting Endpoint di Pipedream

Pipedream merupakan endpoint testing API, untuk membaca header dan body(raw message) dari http client yang kita request. Http client request bisa dari postman maupun kode javascript kita. Pada langkah ini kita akan melakukan:
1. Membuat Akun pipedream
2. Testing Endpoint menggunakan Postman
3. Membangun HTML Form dengan Tailwind CSS Component
4. Membuat file js untuk mengirimkan data form ke endpoint pipedream
5. Melihat hasilnya di dashboard pipedream
6. Production mode menggunakan github pages

Disini kita akan melakukan request dengan menggunakan javascript fetch.

### Membuat Akun

* Kunjungi https://pipedream.com/ , terus klik Sign Up Free! 

  ![image](https://user-images.githubusercontent.com/11188109/220200037-5f556ae6-3bd6-4aa3-9869-00c1d119bb51.png)
  
* Klik sign up with github

  ![image](https://user-images.githubusercontent.com/26703717/222384917-1c39dada-86f2-401e-a394-bdb561eaeac3.png) 
* Authorize

  ![image](https://user-images.githubusercontent.com/26703717/222384917-1c39dada-86f2-401e-a394-bdb561eaeac3.png) 
  
* Pilih Nama Workspace yang tersedia, ditandai dengan tanda centang hijau, kemudian klik Continue

  ![image](https://user-images.githubusercontent.com/26703717/222385437-4206ac91-7c76-46b5-9ed1-807f59f47136.png)
* Opsional membagikan link workspace seperti ini, https://pipedream.com/@chapter03wsulbi/invite?token=d271f288ef995089871c5bdfbec91fd4
* Lanjutkan dengan menekan tombol Skip.
  
  ![image](https://user-images.githubusercontent.com/26703717/222385986-dbf17b1e-2d40-45d9-b6f9-6ad5ce111a0d.png)

* Buat Project Baru
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/7931de30-ebb4-4a38-b81c-4d8ee4634a2f)
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/46ec8f9a-cd7d-4148-bb63-dbd1fc25207e)

* Buat Workflow Baru
![image](https://github.com/Bukulapak/WebService2024/assets/26703717/4c588571-bec3-4036-a076-d7bacc87d6f8)

* Klik New HTTP / Webhook Request
  ![image](https://user-images.githubusercontent.com/26703717/222385768-722b1c42-a0a7-4d67-8508-e5645308e928.png)
* Pilih Event Data : Raw REquest, HTTP Response : 200 OK . Kemudian klik Save and Continue
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/3912d977-8fdc-40f7-a8de-2d233d5409c7)
* Akan keluar unique URL untuk endpoint : https://eocvksohetdfo7p.m.pipedream.net (Ini adalah contoh punya SAYA). Kemudian kita coba dengan postman dahulu
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/fe94bcf0-f078-4507-b6cc-0c7e6f0fa258)

### Melakukan Testing Endpoint

Disini kita akan membuka Postman untuk melakukan testing endpoint dahulu, dengan contoh :
* Method POST Headers kita isi dengan Key : Login , Value : Bebas. 
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/7a424526-6eca-4c39-a66f-61f013e433e1)
  
* Pada bagian body isi dengan data json dengan cara klik Body-->raw-->JSON. Kemudian klik Send
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/5ba5e49c-7074-468e-b9d4-f5ee9b5961a6)

* Setelah klik Send maka akan menampilkan response seperti berikut
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/d26abc39-c42f-4723-86a5-54c7ca9b0138)

* Dashboard Pipedream akan muncul 1 New Event, kita buka event tersebut.
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/bb61f426-7cfc-48ef-a28e-ef2e8a30d0e0)

* Disana akan terlihat pada bagian headers ada Login yang kita masukkan dan pada bagian body ada json yang kita masukkan dari postman. Artinya endpoint dan http request bekerja dengan baik untuk menangkap header dan body yang dikirimkan.
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/c94a591f-fe9b-4b7e-a452-72ef99a65ba5)

### Membuat Form Pendaftaran

* Kita gunakan template berikut. Link : https://github.com/Bukulapak/WebService2024/tree/main/Week03/Praktikum/Template/index.html
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/89cb1eb1-ba82-4c79-9405-961d467e781f)
* Kita copy kodenya untuk kita pakai ditaruh di repo kita pada Folder Week03-->Praktikum-->{NPMKALIAN}-->index.html dan coba kita live server dari vscode.
  ![image](https://user-images.githubusercontent.com/26703717/222399786-36156e16-adc7-4327-8bdd-92f3f2e9de2c.png)
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/a8fbcb9e-a32a-4f9d-b666-b23b159327be)
* Cari element dari input dan buttonnya kemudian kita beri id
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/f6720026-a621-4e4d-b968-9814f7c95972)
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/c9816f17-604b-4375-9788-10054f9f34e4)
* Buat file js kita masukkan dahulu javascript yang didapatkan dari postman ke js tersebut, kemudian panggil pada bagian bawah sebelum tag penutup body

  ![image](https://user-images.githubusercontent.com/26703717/222392397-900e8868-f866-46e4-b939-abaad852c444.png)
  ```html
  <script src="ws.js"></script>
  ```


### Membuat Fungsi Javascript

Pada bagian ini kita akan membuat fungsi-fungsi di javascript untuk mengirimkan data dari html menuju endpoint pipedream yang kita buat. Kenapa harus dibuat fungsi? agar kode program javascript tetap rapih dan mudah terbaca.
1. Membuat fungsi PostSignUp() yang berfungsi untuk melakukan Post Form Data Sign Up. Buat file ws.js di Folder Week03-->Praktikum-->{NPMKALIAN}. 
   ```javascript
    function PostSignUp(namadepan,namabelakang,email,password){
    var myHeaders = new Headers();
    myHeaders.append("Login", "chapter03");
    myHeaders.append("Content-Type", "application/json");
   
    var raw = JSON.stringify({
      "namadepan": namadepan,
      "namabelakang": namabelakang,
      "email": email,
      "password": password
    });
   
    var requestOptions = {
      method: 'POST',
      headers: myHeaders,
      body: raw,
      redirect: 'follow'
    };
   
    fetch("ISILINKPIPEDREAMKALIAN", requestOptions)
      .then(response => response.text())
      .then(result => console.log(result))
      .catch(error => console.log('error', error));
    }
   ```
2. Pada line 20 isikan link pipe dream kalian

   ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/a4809473-09ce-4704-988d-5b9f47e57228)
   ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/411c8e88-afbf-4c86-81b8-694ad3e903d2)

3. Membuat fungsi PushButton() untuk melakukan aksi setelah menekan tombol, pada bagian html button tambahkan atribut onclick.

   ![image](https://user-images.githubusercontent.com/11188109/220208507-0d5cb2cc-4979-410c-a3e0-a804caa732c4.png)
   ```javascript
    function PushButton(){
          namadepan=document.getElementById("namadepan").value;
          namabelakang=document.getElementById("namabelakang").value;
          email=document.getElementById("email").value;
          password=document.getElementById("password").value;
          PostSignUp(namadepan,namabelakang,email,password);
    }
   ```
4. Kita test dengan klik kanan Open with live server, kita isi form nya sambil inspect console lalu klik button yang kita buat. Terlihat dari console data success dikirim. 
   ![image](https://user-images.githubusercontent.com/26703717/222400423-db80721c-07df-412f-bee6-cd5928b76e98.png)
   ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/5ade8525-9813-40e0-aec2-ba01f87e595e)
5. Kemudian kita lihat pada dashboard pipedream data sudah diterima dengan baik oleh endpoint baik itu header maupun body.
   ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/838f4b6b-6b99-48ee-a496-1fe7b937dd17)

### Tambahan Estetika UX

Pasti merasa aneh bukan setelah menekan tombol, tapi tampilan tidak berubah sama sekali seolah tidak terjadi apa-apa. Disini kita akan coba untuk mengubah tampilan jika tombol di klik, maka form akan disembunyikan dan menampilkan data yang diterima dari endpoint pipedream. Oke kita cukup menambahkan satu fungsi lagi dan memodifikasi fungsi PostSignUp() khususnya dibagian result then fetch nya. Langkahnya sebagai berikut :

* Pertama kita cari dulu element yang akan kita hidden, bisa menggunakan inspect elemetns untuk identifikasinya, kemudian kita kasih id
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/988d4508-6697-4f67-995d-5bf672a3a602)
  ![image](https://user-images.githubusercontent.com/11188109/220210342-a81493b4-453b-4cd1-a526-3ca1092f6ebc.png)
* Kita coba script style display pada bagian console dengan menggunakan id yang sudah kita buat, style display mana yang bisa menghilangkan formsignup apakah block atau none. Terlihat di gambar none bisa mengilangkan element formsignup, kita **mungkin** akan pakai script ini di fungsi kita yang akan datang. Dari sini kita paham bagaimana menyembunyikan element dengan js.
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/172ad78e-bafb-444b-af5f-e0c24a19c7be)
  ```javascript
    document.getElementById("formsignup").style.display = 'none';
    document.getElementById("formsignup").style.display = 'block';
  ```
* Tambahkan fungsi GetResponse() di file js kita yang sudah dibuat sebelumnya, dan melakukan modifikasi dari fungsi PostSignUp pada bagian then result.
  ```javascript
    function GetResponse(result){
          document.getElementById("formsignup").innerHTML = result;
    }
  ```
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/93eadef9-f3a6-4422-ad9b-aaeebc104e8e)
* Kita ujicoba dengan mengisi form dan klik tombol submit, maka form kita sudah berhasil.
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/725687af-a66a-42e8-a500-015266486c99)
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/419c0e6c-e335-42c7-bc43-cee0548fe7bf)
* Cek pada web pipedream, seharusnya terdapat event baru seperti gambar di bawah
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/881525bb-f28b-42ad-b22f-82ff893de1cb)
* Push pada repo github masing-masing Week03/Praktikum/{NPM masing-masing}, kemudian lakukan pull request, dan coba jalankan hasil yang sudah dibuat pada Github Pages. Alamat github pages biasanya sub domain dari github.io. Contoh : https://indrariksa.github.io/WS/Week3/Site/NPM

## Kerjakan (Tugas Take Home)

* Buatlah form sign up yang melakukan POST ke pipedream.com ketika klik button.
* Form inputan yang dibuat minimal berisi 5 form.
* Buat Folder NPM masing-masing didalam Week03/Tugas/{NPM masing-masing} yang berisi 2 file minimal js dan html, dengan nama index.html, ws.js. Boleh menambahkan file css.
* file ws.js minimal berisi 3 fungsi, tidak boleh ada kode js diluar dari fungsi, semua harus masuk ke dalam fungsi js.
* Form dibangun dengan menggunakan CSS tailwind, desain html tidak boleh sama, harus berbeda satu sama lain.
* Pull Request dengan nama 3-Kelas-NPM-NAMA di folder Week03/Tugas/{NPM masing-masing}, dengan deskripsi disertakan di bawah ini.
* Sertakan screenshoot dari live server aplikasi dan pipedream.com
* Sertakan link Github Pages Sudah jalan di repo masing-masing yang dimasukkan pada file README.md
* Sertakan screenshoot juga sertifikat dari sertifikasi berikut (Minimal 1 sertifikat): 
  * https://www.mygreatlearning.com/academy/learn-for-free/courses/go-programming-language 
  * https://www.mindluster.com/certificate/3394
  * https://www.codecademy.com/learn/learn-go
  * https://www.coursera.org/specializations/google-golang

