function Signup(firstname, lastname, username, email, password) {
  var myHeaders = new Headers();
  myHeaders.append("Signup", "Genshin Website");
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
    firstname: firstname,
    lastname: lastname,
    username: username,
    email: email,
    password: password,
  });

  var requestOptions = {
    method: "POST",
    headers: myHeaders,
    body: raw,
    redirect: "follow",
  };

  fetch("https://eo5u30meas128sd.m.pipedream.net", requestOptions)
    .then((response) => response.text())
    .then((result) => GetResponse(result))
    .catch((error) => GetError(error));
}

function GetResponse(result) {
  document.getElementById("forms").innerHTML = result + "<br>You will be redirected in 3 sec...";
  setTimeout(function () {
    window.location.reload();
  }, 3000);
}

function GetError(error) {
  document.getElementById("forms").innerHTML = "<b>Can't sign up</b></br>" + error + "<br><br>You will be redirected in 5 sec...";
  setTimeout(function () {
    window.location.reload();
  }, 5000);
}

function signupNow() {
  document.getElementById("btnsignup").innerHTML = `<button
        class="flex items-center justify-center w-full text-lg rounded-full bg-indigo-500 hover:bg-indigo-600 disabled"
        ">
        <svg class="mr-3 h-5 w-5 animate-spin text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
        <span class="font-medium">Wait for a while...</span></button>`;
  firstname = document.getElementById("firstname").value;
  lastname = document.getElementById("lastname").value;
  username = document.getElementById("username").value;
  email = document.getElementById("email").value;
  password = document.getElementById("password").value;
  Signup(firstname, lastname, username, email, password);
}
