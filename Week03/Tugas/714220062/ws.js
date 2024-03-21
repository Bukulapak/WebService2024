function PostSignUp(firstname, lastname, email, password) {
    const myHeaders = new Headers();
    myHeaders.append("SignUp", "Overwatch");
    myHeaders.append("Content-Type", "application/json");

    const raw = JSON.stringify({
        "firstname": firstname,
        "lastname": lastname,
        "email": email,
        "password": password
    });

    const requestOptions = {
        method: "POST",
        headers: myHeaders,
        body: raw,
        redirect: "follow"
    };

    fetch("https://eosgo0pogw8uh5k.m.pipedream.net", requestOptions)
        .then((response) => response.text())
        .then((result) => GetResponse(result))
        .catch((error) => console.error(error));
}

function GetResponse(result) {
    document.getElementById("formsignup").innerHTML = result;
}

function PushButton() {
    firstname = document.getElementById("firstname").value;
    lastname = document.getElementById("lastname").value;
    username = document.getElementById("username").value;
    email = document.getElementById("email").value;
    password = document.getElementById("password").value;
    PostSignUp(firstname, lastname, username, email, password); 
}
