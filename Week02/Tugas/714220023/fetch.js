const apiUrl = "https://api.gameofthronesquotes.xyz/v1/characters";
const container = document.getElementById("data-container");

// Fungsi untuk menampilkan data dalam tabel
function displayData(characters) {
  characters.forEach(character => {
    const row = document.createElement("tr");

    // Name
    const nameCell = document.createElement("td");
    nameCell.textContent = character.name;
    row.appendChild(nameCell);

    // House
    const houseCell = document.createElement("td");
    houseCell.textContent = character.house.name;
    row.appendChild(houseCell);

    // Quotes
    const quotesCell = document.createElement("td");
    const quotesList = document.createElement("ul");

    character.quotes.forEach(quote => {
      const quoteItem = document.createElement("li");
      quoteItem.textContent = quote;
      quotesList.appendChild(quoteItem);
    });

    quotesCell.appendChild(quotesList);
    row.appendChild(quotesCell);

    container.appendChild(row);
  });
}

// Fetch data dari API dan konversi respons ke JSON
fetch(apiUrl)
  .then(response => {
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }
    return response.json(); // Mengubah respons menjadi objek JSON
  })
  .then(data => displayData(data))  // Menampilkan data dalam tabel
  .catch(error => console.error(error));
