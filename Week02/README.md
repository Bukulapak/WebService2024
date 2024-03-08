# TOOLS

1. POSTMAN
   * Install postman.
   ```sh
   https://www.postman.com/
   ```
   * lakukan request ke public api dengan method GET dengan postman contoh : https://alexwohlbruck.github.io/cat-facts/docs/
   ![image](https://user-images.githubusercontent.com/26703717/220868366-271e81d8-707e-4611-9aa7-74f8b31cd008.png)
   
2. Template HTML
   * Untuk membuat antarmuka menggunakan komponen tailwind contoh :https://github.com/Bukulapak/WebService2024/tree/main/Week02/Praktikum/index.html
   
3. Plugin LiveServer (VScode)
   ![image](https://user-images.githubusercontent.com/11188109/218396548-483f109a-c88c-4bc6-96d0-5d784a447556.png)

# Menghubungkan situs dengan Public API

1. Buat folder NPM masing-masing (contoh :123456) di dalam folder Praktikum
2. Simpan file Template HTML yaitu index.html ke dalam folder NPM kalian
3. buat file index.html panggil di bagian paling bawah dengan script, sebelum tag penutup body
   ```html
   <script src="./fetch.js"></script>
   ```
4. Buat file fetch.js di dalam folder NPM kalian
5. Pada aplikasi Postman, pilih menu </> atau code pilih javascript fetch ,kemudian paste pada file fetch.js yang baru dibuat
   ![image](https://user-images.githubusercontent.com/26703717/220873867-a5685abb-3e1b-43ce-a1dc-57a5b3b38065.png)
6. Kemudian tambahkan kode berikut pada script fetch.js :
    ```js
    var myHeaders = new Headers();
    myHeaders.append("Cookie", "connect.sid=s%3AM6CfLJhCFlu92tNvS7cRegIIcR8rhhUG.AN2Ss3OKnMLlBJEwcDELKykDb293dBuH%2FhX1M3mZI2w");
    
    var requestOptions = {
      method: 'GET',
      headers: myHeaders,
      redirect: 'follow'
    };
    
    fetch("https://cat-fact.herokuapp.com/facts", requestOptions)
      .then(response => response.text())
      .then(result => tampilkan(result))
      .catch(error => console.log('error', error));
    
    function tampilkan(result) {
      console.log(result);
      iniJson = JSON.parse(result);
    }   
    }
    ```
7. Buka dengan live server, inspect lihat di console dan hasilnya ditampilkan berupa teks.

    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/cde521d7-f30b-484c-bbda-fd2897375efa)
    
    untuk akses variabel global hasil kita pakai console log
    
    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/d78693ed-017f-4a71-ad8b-8bd13a6bd57b)
    
    kemudian tambahkan ramuan looping untuk mengeluarkan semuanya
    
    ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/756f45cf-77ef-4d9e-8280-d7c336151592)
    
# Setting Repo untuk Github Pages
  * Lakukan pada repo masing-masing mahasiswa
  * Setting pages pilih deployment from branch pilih main. Maka form kita bisa diakses dari menjadi github pages. Alamat github pages biasanya sub domain dari github.io. Contoh : {username github}.github.io --> https://bukulapak.github.io/WebService2024/Week02/Praktikum
  ![image](https://github.com/Bukulapak/WebService2024/assets/26703717/51759110-b433-44c4-a7e1-76fe60c2616e)
