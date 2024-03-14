const requestOptions = {
  method: "GET",
  redirect: "follow"
};

fetch("https://www.moogleapi.com/api/v1/games", requestOptions)
  .then((response) => response.json())
  .then((result) => {
    console.log(result);
    tampilkan(result);
  })
  .catch((error) => console.error(error));

function tampilkan(tampil) {
  let hasil = ""; 
  console.log(tampil);

  for (let i = 0; i < tampil.length; i++) {
    hasil += "<tr>";
    hasil += "<td>" + tampil[i].title + "</td>";
    hasil += "<td><img src='" + tampil[i].picture + "' alt='" + tampil[i].title + "' width='100'></td>"; 
    hasil += "<td>" + tampil[i].platform + "</td>";
    hasil += "<td>" + tampil[i].releaseDate + "</td>";
    hasil += "</tr>";
  }
  document.getElementById("game").innerHTML = hasil;
}
