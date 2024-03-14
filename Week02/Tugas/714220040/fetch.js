var myHeaders = new Headers();
myHeaders.append("Cookie", "connect.sid=s%3AM6CfLJhCFlu92tNvS7cRegIIcR8rhhUG.AN2Ss3OKnMLlBJEwcDELKykDb293dBuB%2FhX1M3mZI2w");

var requestOptions = {
  method: 'GET',
  headers: myHeaders,
  redirect: 'follow'
};

fetch("https://api.trace.moe/search?anilistID=1&url=https://images.plurk.com/32B15UXxymfSMwKGTObY5e.jpg", requestOptions)
  .then(response => response.json()) 
  .then(result => tampilkan(result)) 
  .catch(error => console.log('error', error));

function tampilkan(result) {
  var hasil = ""; 
  console.log(result);

  for (var i = 0; i < result.length; i++) { // Menggunakan result langsung
    hasil += "<tr>";
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+ result[i].anilist +"</td>"; 
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+ result[i].filename +"</td>"; 
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>"+ result[i].episode +"</td>"; 
    hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'><img src='" + result[i].image + "' alt='Emoticon Image'></td>"; // Menambahkan tag <img> untuk menampilkan gambar
    hasil += "</tr>";
  }

  document.getElementById("emoticon").innerHTML = hasil; // Menggunakan ID "emoticon"
}
