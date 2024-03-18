const requestOptions = {
    method: "GET",
    redirect: "follow"
  };
  
  fetch("https://api.nasa.gov/neo/rest/v1/feed?start_date=2015-09-07&end_date=2015-09-08&api_key=DEMO_KEY", requestOptions)
    .then((response) => response.json())
    .then((result) => {
      // Mengambil data NEO dari result
      const neoData = Object.values(result.near_earth_objects).flat();
      tampilkan(neoData);
    })
    .catch((error) => console.error(error));
  
  function tampilkan(tampil) {
    let hasil = ""; 
    console.log(tampil);
  
    for (let i = 0; i < tampil.length; i++) {
      hasil += "<tr>";
      hasil += "<td>" + tampil[i].id + "</td>";
      hasil += "<td>" + tampil[i].neo_reference_id + "</td>";
      hasil += "<td>" + tampil[i].name + "</td>";
      hasil += "<td>" + tampil[i].absolute_magnitude_h + "</td>"; // Perhatikan typo di sini
      hasil += "</tr>";
    }
    document.getElementById("inidata").innerHTML = hasil;
  }
  