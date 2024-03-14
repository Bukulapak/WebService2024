// Mendefinisikan fungsi untuk mengambil data dari API dan mengisi tabel
function getDataAndPopulateTable() {
    var requestOptions = {
      method: 'GET',
      redirect: 'follow'
    };
  
    fetch("https://strangerthings-quotes.vercel.app/api/quotes/5", requestOptions)
      .then(response => response.json())
      .then(result => populateTable(result))
      .catch(error => console.error(error));
  }
  
  // Fungsi untuk mengisi tabel dengan data yang diterima dari API
  function populateTable(result) {
    var tableData = ""; // String untuk menyimpan baris-baris tabel
  
    // Loop melalui hasil JSON
    result.forEach(item => {
      // Membuat baris baru untuk setiap item
      tableData += "<tr>";
      // Memasukkan quote dan author ke dalam kolom tabel
      tableData += "<td>" + item.quote + "</td>";
      tableData += "<td>" + item.author + "</td>";
      tableData += "</tr>";
    });
  
    // Memasukkan data ke dalam tbody tabel
    document.getElementById("inidata").innerHTML = tableData;
  }
  
  // Memanggil fungsi untuk mengambil data dan mengisi tabel saat halaman dimuat
  getDataAndPopulateTable();
  