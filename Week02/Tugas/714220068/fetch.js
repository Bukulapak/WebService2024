var agentIndex = 0;

function generateAgents() {
  var myHeaders = new Headers();
  myHeaders.append("Cookie", "connect.sid=s%3AM6CfLJhCFlu92tNvS7cRegIIcR8rhhUG.AN2Ss3OKnMLlBJEwcDELKykDb293dBuH%2FhX1M3mZI2w");

  var requestOptions = {
    method: 'GET',
    headers: myHeaders,
    redirect: 'follow'
  };

  fetch("https://www.thecocktaildb.com/api/json/v1/1/random.php", requestOptions)
    .then(response => response.json())
    .then(data => displayRandomCocktails(data))
    .catch(error => console.log('error', error));
}

function displayRandomCocktails(data) {
  console.log(data);
  
  var cocktails = getRandomCocktails(data.drinks, 5);
  var cocktailsTable = document.getElementById("CocktailsTable");
  cocktailsTable.innerHTML = "";

  cocktails.forEach(cocktail => {
    var show = "<tr>";
    show += "<td scope='col' class ='px-6 py-4 font-medium text-white-900'>" + cocktail.strDrink + "</td>";
    show += "<td scope='col' class ='px-6 py-4 font-medium text-white-900'>" + cocktail.strCategory + "</td>";
    show += "<td scope='col' class ='px-6 py-4 font-medium text-white-900'>" + cocktail.strAlcoholic + "</td>";
    show += "<td scope='col' class ='px-6 py-4 font-medium text-white-900'>" + cocktail.strGlass + "</td>";
    show += "<td scope='col' class='px-6 py-4 font-medium text-white-900'><img src='" + cocktail.strDrinkThumb + "' alt='" + cocktail.strDrink + "'></td>";
    show += "</tr>";  
    cocktailsTableBody.innerHTML += show;
  });
}


function getRandomCocktails(cocktailData, n) {
  var shuffledCocktails = cocktailData.sort(() => 0.5 - Math.random());
  return shuffledCocktails.slice(0, n);
}