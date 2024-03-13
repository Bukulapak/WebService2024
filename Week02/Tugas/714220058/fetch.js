const requestOptions = {
  method: "GET",
  redirect: "follow"
};

fetch("https://www.theaudiodb.com/api/v1/json/2/search.php?s=coldplay", requestOptions)
  .then((response) => response.json()) // Ubah response menjadi JSON
  .then((result) => tampilkan(result)) 
  .catch((error) => console.error(error));

function tampilkan(result) {
  console.log(result);
  const artists = result.artists; 

  let tabelHTML = ""; 

  // Loop melalui setiap artis dan tambahkan baris tabel untuk setiap artis
  artists.forEach((artist) => {
    tabelHTML += "<tr>";
    tabelHTML += "<td class='pb-3 text-start min-w-[175px]'>" + artist.strArtist + "</td>";
    tabelHTML += "<td class='pb-3 text-end min-w-[100px]'>" + artist.strBiographyEN + "</td>";
    tabelHTML += "<td class='pb-3 text-end min-w-[100px]'>" + artist.strGenre + "</td>";
    tabelHTML += "<td class='pb-3 pr-12 text-end min-w-[175px]'>" + artist.intFormedYear + "</td>";
    tabelHTML += "<td class='pb-3 pr-12 text-end min-w-[100px]'>" + artist.strCountry + "</td>";
    // Menambahkan tag img untuk menampilkan gambar 
    tabelHTML += "<td class='pb-3 text-end min-w-[50px]'><img src='" + artist.strArtistThumb + "' alt='" + artist.strArtist + "' style='max-width:100px;max-height:100px;'></td>";
    tabelHTML += "</tr>";
  });
  document.getElementById("inidata").innerHTML = tabelHTML;
}
