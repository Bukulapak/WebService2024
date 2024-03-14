const myHeaders = new Headers();
myHeaders.append("Cookie", "connect.sid=s%3Az3QRPcoatnMvXjMzG4VezIvFJ_Usy0-_.e%2FlRGUrLBeQqFjYoSSm%2F1Ehli9qGXKUs6Xpy9xtHrqw");

const requestOptions = {
    method: "GET",
    redirect: "follow"
  };
  
  fetch("https://api.jikan.moe/v4/anime?q=One%20Piece&sfw", requestOptions)
    .then(response => response.json())
    .then(result => tampilkan(result))
    .catch(error => console.log('error', error));
  
  function tampilkan(result) {
    const data = result.data;
    const tbody = document.querySelector('#inidata tbody');
  
    tbody.innerHTML = '';
  
    data.forEach(item => {
      const row = document.createElement('tr');
  
      const namaCell = document.createElement('td');
      namaCell.textContent = item.type;
      row.appendChild(namaCell);

      const posterCell = document.createElement('td');
      const posterImage = document.createElement('img');
      posterImage.src = item.images.jpg.large_image_url; // Menampilkan poster film
      posterImage.alt = "Poster";
      posterImage.width = 100; // Atur lebar poster sesuai kebutuhan
      posterCell.appendChild(posterImage);
      row.appendChild(posterCell);
  
      const judulCell = document.createElement('td');
      judulCell.textContent = item.titles[0].title; // Menampilkan judul default
      row.appendChild(judulCell);
  
      const trailerCell = document.createElement('td');
      const trailerIframe = document.createElement('iframe');
      trailerIframe.width = "360";
      trailerIframe.height = "215";
      trailerIframe.src = item.trailer.embed_url; // Menampilkan trailer dalam tag iframe
      trailerIframe.allow = "accelerometer; encrypted-media; gyroscope; picture-in-picture";
      trailerIframe.allowFullscreen = true;
      trailerCell.appendChild(trailerIframe);
      row.appendChild(trailerCell);
  
      tbody.appendChild(row);
    });
  }
  