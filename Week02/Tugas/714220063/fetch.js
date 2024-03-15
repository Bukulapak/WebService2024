const url = "https://api.imgflip.com/get_memes";

fetch(url)
  .then(response => response.json())
  .then(data => {
    const memes = data.data.memes;
    const tableBody = document.getElementById("memesBody");

    memes.forEach(meme => {
      const row = document.createElement("tr");
      const idCell = document.createElement("td");
      const nameCell = document.createElement("td");
      const imageCell = document.createElement("td");

      idCell.textContent = meme.id;
      nameCell.textContent = meme.name;

      const image = document.createElement("img");
      image.src = meme.url;
      image.alt = meme.name;
      image.style.maxWidth = "100px"; 

      imageCell.appendChild(image);

      row.appendChild(idCell);
      row.appendChild(nameCell);
      row.appendChild(imageCell);

      tableBody.appendChild(row);
    });
  })
  .catch(error => console.error(error));
