const requestOptions = {
  method: "GET",
  redirect: "follow"
};

fetch("https://api.artic.edu/api/v1/artworks", requestOptions)
  .then(response => response.json()) // Mengubah respon ke JSON
  .then(result => tampilkan(result)) // Memanggil fungsi tampilkan dengan data JSON
  .catch(error => console.error(error));

function tampilkan(result) {
  console.log(result);
  const iniJson = result.data; // Akses ke data hasil API

  let hasil = "";

  for (let i = 0; i < iniJson.length; i++) {
    const imageUrl = "https://www.artic.edu/iiif/2/" + iniJson[i].image_id + "/full/843,/0/default.jpg";
    const description = iniJson[i].description ? iniJson[i].description : "-"; // Kondisi jika deskripsi null

    hasil += `
      <div class="bg-gray-900 shadow-md rounded-md overflow-hidden w-80 font-['Poppins] cursor-default">
      <img src="${imageUrl}" alt="${iniJson[i].thumbnail.alt_text}" class="w-full h-64 object-cover transition duration-300 hover:h-auto">
      <div class="p-4">
          <h3 class="font-bold text-lg">${iniJson[i].title}</h3>
          <div class="h-28 overflow-y-auto">
            <p class="text-gray-700">${description}</p>
          </div>
          <p class="text-sm text-gray-600 mt-2">ID: ${iniJson[i].id}</p>
          <p class="text-sm text-gray-600 mb-2">${iniJson[i].artist_display}</p>
        </div>
      </div>
    `;
  }

  document.getElementById("inidata").innerHTML = hasil;
} 
