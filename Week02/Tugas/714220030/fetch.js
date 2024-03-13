const requestOptions = {
    method: "GET",
    redirect: "follow"
  };
  
  fetch("https://db.ygoprodeck.com/api/v7/cardinfo.php?staple=yes", requestOptions)
    .then((response) => response.text())
    .then((result) => console.log(result))
    .catch((error) => console.error(error));

     
     function fetchData() {
        fetch('https://db.ygoprodeck.com/api/v7/cardinfo.php?staple=yes')
            .then(response => response.json())
            .then(data => {
                const cardTableBody = document.getElementById('cardTableBody');
                data.data.forEach(card => {
                    const row = document.createElement('tr');
                    row.innerHTML = `
                        <td>${card.id}</td>
                        <td>${card.name}</td>
                        <td>${card.type}</td>
                        <td>${card.attribute}</td>
                    `;
                    cardTableBody.appendChild(row);
                });
            })
            .catch(error => console.error('Error fetching data:', error));
    }

   
    window.onload = fetchData;