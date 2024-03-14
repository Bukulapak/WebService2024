const requestOptions = {
    method: "GET",
    redirect: "follow"
  };
  
  fetch("https://api.brawlapi.com/v1/brawlers", requestOptions)
    .then(response => response.json()) // Parse response as JSON
    .then(result => {
      const tableBody = document.querySelector("#BrawlTable tbody");
  
      // Loop through each brawler and add a new row to the table
      result.list.forEach(brawler => {
        const newRow = tableBody.insertRow();
        newRow.innerHTML = `
          <td>${brawler.name}</td>
          <td><img src="${brawler.imageUrl}" alt="${brawler.name}" width="100"></td>
          <td>${brawler.class.name}</td>
          <td style="color: ${brawler.rarity.color}">${brawler.rarity.name}</td>
          <td>${brawler.description}</td>
        `;
      });
    })
    .catch(error => console.error(error));  