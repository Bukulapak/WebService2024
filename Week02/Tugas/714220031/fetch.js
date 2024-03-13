// Definisikan headers
var myHeaders = new Headers();
myHeaders.append("Cookie", "connect.sid=s%3AM6CfLJhCFlu92tNvS7cRegIIcR8rhhUG.AN2Ss3OKnMLlBJEwcDELKykDb293dBuH%2FhX1M3mZI2w");

// Konfigurasi permintaan
var requestOptions = {
    method: 'GET',
    headers: myHeaders,
    redirect: 'follow'
};

// Fungsi untuk menampilkan data
function tampilkan(result){
    console.log(result);
    var iniJson = JSON.parse(result);

    var length = iniJson.length;
    console.log(length);

    var hasil = "";

    for (var i = 0; i < length; i++){
        hasil += "<tr>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].text + "</td>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].user + "</td>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].type + "</td>";
        hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].updateat + "</td>";
        hasil += "</tr>";
    }

    // Masukkan hasil ke dalam tabel HTML
    document.getElementById("inidata").innerHTML = hasil;
}

// Lakukan fetch data dari API
fetch("https://cat-fact.herokuapp.com/facts", requestOptions)
    .then(response => response.text())
    .then(result => tampilkan(result))
    .catch(error => console.log('error', error));
