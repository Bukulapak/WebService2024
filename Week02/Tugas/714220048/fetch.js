const requestOptions = {
    method: "GET",
    redirect: "follow"
  };
  
  fetch("https://api.thecatapi.com/v1/images/search?limit=10", requestOptions)
    .then((response) => response.json())
    .then((result) => tampilkan(result))
    .catch((error) => console.error(error));
  
  let hasil = "";
  
  function tampilkan(result) {
    console.log(result);
    const iniJson = result;
    const length = iniJson.length;
  
    for (let i = 0; i < length; i++) {
      const imageUrl = iniJson[i].url;
      hasil += "<tr>";
      hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].id + "</td>";
      hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'><img src='" + imageUrl + "' style='max-width: 200px; max-height: 200px;'></td>";
      hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].width + "</td>"; // Perbaikan typo
      hasil += "<td scope='col' class='px-6 py-4 font-medium text-gray-900'>" + iniJson[i].height + "</td>";
      hasil += "</tr>";
    }
    document.getElementById("inidata").innerHTML = hasil;
  }
  