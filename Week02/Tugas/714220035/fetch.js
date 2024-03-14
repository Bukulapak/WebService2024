const requestOptions = {
  method: "GET",
  redirect: "follow"
};

fetch("https://hp-api.onrender.com/api/characters", requestOptions)
  .then((response) => response.text())
  .then((result) => tampilkan(result))
  .catch((error) => console.error(error));

  function tampilkan(result) {
    console.log(result);
    iniJson = JSON.parse(result);
  
    length = iniJson.length;
    console.log(length);
  
    let tableBody = document.querySelector("#inidata tbody");
    let hasil = "";
  
    for (let i = 0; i < length; i++) {
      hasil += "<tr>";
      hasil += "<td>" + iniJson[i].name + "</td>";
      hasil += "<td>" + iniJson[i].alternate_names + "</td>";
      hasil += "<td>" + iniJson[i].dateOfBirth + "</td>";
      hasil += "<td><img class='small-image' src='" + iniJson[i].image + "' alt='" + iniJson[i].image + "'></td>";
      hasil += "</tr>";
    }
  
    tableBody.innerHTML = hasil;
  }
  