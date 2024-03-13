var agentIndex = 0;

function generateAgents() {
  var myHeaders = new Headers();
  myHeaders.append("Cookie", "connect.sid=s%3AM6CfLJhCFlu92tNvS7cRegIIcR8rhhUG.AN2Ss3OKnMLlBJEwcDELKykDb293dBuH%2FhX1M3mZI2w");

  var requestOptions = {
    method: 'GET',
    headers: myHeaders,
    redirect: 'follow'
  };

  fetch("https://valorant-api.com/v1/agents", requestOptions)
    .then(response => response.json())
    .then(data => displayRandomAgents(data))
    .catch(error => console.log('error', error));
}

function displayRandomAgents(data) {
  console.log(data);
  
  var agents = getRandomAgents(data.data, 5);
  var agentsTable = document.getElementById("AgentsTable");
  agentsTable.innerHTML = "";

  agents.forEach(agent => {
    var show = "<tr>";
    show += "<td scope='col' class ='px-6 py-4 font-medium text-white-900'>" + agent.displayName + "</td>";
    show += "<td scope='col' class ='px-6 py-4 font-medium text-white-900'>" + agent.role.displayName + "</td>";
    show += "<td scope='col' class ='px-6 py-4 font-medium text-white-900'>" + agent.developerName + "</td>";
    show += "<td scope='col' class='px-6 py-4 font-medium text-white-900'><img src='" + agent.fullPortrait + "' alt='" + agent.displayName + "'></td>";
    show += "</tr>";  
    agentsTable.innerHTML += show;
  });
}


function getRandomAgents(agentData, n) {
  var shuffledAgents = agentData.sort(() => 0.5 - Math.random());
  return shuffledAgents.slice(0, n);
}