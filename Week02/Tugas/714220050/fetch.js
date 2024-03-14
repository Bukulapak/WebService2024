const requestOptions = {
    method: 'GET',
    redirect: 'follow'
  };
  
  fetch("https://api.disneyapi.dev/character", requestOptions)
    .then(response => response.json())
    .then(result => displayCharacters(result.data))
    .catch(error => console.log('error', error));
  
  function displayCharacters(characters) {
    const table = document.getElementById("charactersTable");
    const tbody = document.getElementById("charactersTableBody");
    tbody.innerHTML = "";
  
    characters.forEach(character => {
      const row = tbody.insertRow();
      row.insertCell().textContent = character.name;
      row.insertCell().textContent = character.films.join(', ');
      row.insertCell().textContent = character.tvShows.join(', ');
      row.insertCell().textContent = character.videoGames.join(', ');
  
      const sourceUrlCell = row.insertCell();
      const sourceUrlLink = document.createElement('a');
      sourceUrlLink.href = character.sourceUrl;
      sourceUrlLink.textContent = "More Info";
      sourceUrlCell.appendChild(sourceUrlLink);
    });
  }
  