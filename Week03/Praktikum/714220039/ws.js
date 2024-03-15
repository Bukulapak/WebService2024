function PostSignUp(namadepan, namabelakang, email, password) {
    var myHeaders = new Headers();
    myHeaders.append("Login", "chapter03");
    myHeaders.append("Content-Type", "application/json");

    var raw = JSON.stringify({
        "namadepan": namadepan,
        "namabelakang": namabelakang,
        "email": email,
        "password": password
    });

    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: raw,
        redirect: 'follow'
    };

    fetch("https://eo22qqddg27143x.m.pipedream.net", requestOptions)
        .then(response => response.text())
        .then(result => GetResponse(result))
        .catch(error => console.error('error', error));
}

function GetResponse(result) {
    document.getElementById("formsignup").innerHTML = result;
}

function PushButton() {
    var namadepan = document.getElementById("namadepan").value;
    var namabelakang = document.getElementById("namabelakang").value;
    var email = document.getElementById("email").value;
    var password = document.getElementById("password").value;
    PostSignUp(namadepan, namabelakang, email, password);
}
