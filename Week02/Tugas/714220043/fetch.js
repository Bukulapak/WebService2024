const requestOptions = {
  method: 'GET',
  redirect: 'follow',
};

fetch('https://api.attackontitanapi.com/titans', requestOptions)
  .then((response) => response.json()) // Parse response as JSON
  .then((data) => {
    const tableBody = document.querySelector('#AOT tbody');

    data.results.forEach((titan) => {
      const row = tableBody.insertRow();
      row.innerHTML = `
          <td>${titan.name}</td>
          <td><img src="${titan.img}" alt="${titan.name}" width="100"></td>
          <td>${titan.height}</td>
          <td>${titan.abilities.join(', ')}</td>
          <td>${titan.allegiance}</td>
        `;
    });
  })
  .catch((error) => console.error(error));
